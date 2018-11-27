package core

import (
	"sync"
	"github.com/fadeAce/apper/const"
	"time"
)

type pool struct {
	sync.RWMutex
	pips map[int]*pipe
}

type pipe struct {
	sync.RWMutex
	// if it's ready for taken
	state int
	txnID string
	// this for what fragment it caught with
	fragmentSeq int
	pipSeq      int
	timeout     time.Duration
}

var PipPool = func() *pool {
	p := &pool{
		pips: make(map[int]*pipe),
	}
	return p
}()

// algorithm for matching the num of pipes to a given task
// what returned are threads bond to selected task
// if no pipe is spared this invoke would be blocked
func (t *task) MatchPIP() (result map[int]*pipe) {
	PipPool.Lock()
	// there's always at least one routine to run task in serialized method
	// match pipes that suit for a log function which means
	// step.1 search for those idle
	tar := t.fragments.sum
	sum, pips := count()
	for {
		res := match(sum, tar)
		if res == _const.DEFAULT_SUM_VALUE {
			time.Sleep(5 * time.Second)
			continue
		}
		till := _const.DEFAULT_SUM_VALUE
		for index, val := range pips {
			mid := val
			// binding txnID to those pipes
			mid.txnID = t.txID
			result[index] = mid
			t.schedule[index] = false
			till++
			if till == sum {
				break
			}
		}
		break
	}
	PipPool.Unlock()
	return
}
func match(sum int, tar int) int {
	container := sum * 3 / 5
	if _const.DEFAULT_SUM_VALUE == container &&
		sum > _const.DEFAULT_SUM_VALUE {
		return sum
	}
	if container >= tar {
		return tar
	}
	return container
}

func count() (sum int, res map[int]*pipe) {
	sum = _const.DEFAULT_SUM_VALUE
	res = make(map[int]*pipe)
	PipPool.Lock()
	for seq, pip := range PipPool.pips {
		if pip.state == _const.PIP_IDLE {
			res[seq] = pip
			sum++
		}
	}
	PipPool.Unlock()
	return sum, res
}

func (t *task) release() {
	PipPool.Lock()
	defer PipPool.Unlock()
	for k, _ := range t.schedule {
		pip := PipPool.pips[k]
		pip.release()
	}
}

// task done involved a process that inject result from
// fragments to cache layer.
func (t *task) Done() string {
	txnID := t.txID
	cacheCenter.RLock()
	cacheLayer := cacheCenter.data[txnID]
	ch := cacheLayer.ch
	cacheCenter.RUnlock()
	// block here till works done
	<-ch
	// release the routines bond to counter task
	t.release()
	return t.txID
}

func (pip *pipe) release() {
	pip.state = _const.PIP_IDLE
	pip.txnID = ""
	pip.fragmentSeq = _const.DEFAULT_SUM_VALUE
	pip.timeout = _const.DEFAULT_TIMEOUT * time.Second
}

func (p *pool) addPip(pip *pipe) {
	p.Lock()
	defer p.Unlock()
	p.pips[pip.pipSeq] = pip
}
