package core

import (
	"fmt"
	"time"
	"context"
	"gitlab.pandaminer.com/scar/apper/logger"
	"gitlab.pandaminer.com/scar/apper/types"
	"gitlab.pandaminer.com/scar/apper/const"
	"gitlab.pandaminer.com/scar/apper/storage"
	"gitlab.pandaminer.com/scar/apper/handler"
	"sync"
)

var log = logger.Log

// pending centre for tasks
var Panel = func() *pendCentre {
	pc := pendCentre{}
	return &pc
}()

// type task represent a task that match a configuration
type task struct {
	sync.RWMutex
	fragments fragments
	txID      string
	schedule  map[int]bool
	timeout   time.Duration
}

// fragments represents segments for task
type fragments struct {
	// sum
	sum int
	// fragments map index - fragment
	data map[int]fragment
}

type fragment struct {
	// true for taken, default false
	taken      bool
	single     types.Single
	motherSite string
}

// core concept of cushion area for tasks
type pendCentre struct {
	queue chan *task
}

func (p *pendCentre) Init(conf *types.ApperConf) {
	p.queue = make(chan *task, conf.CushionSize)
}

// pending tasks onto the panel
func (p *pendCentre) Pending(tsk *task) {
	p.queue <- tsk
}

// pop up a task
func (*pendCentre) pop() (*task) {
	return nil
}

// generate task from configuration
func Generate(
	ctx context.Context, conf types.ConfJ, database *storage.Database, timeout int,
) *task {
	// step.1 generate txnID
	// seq as a sequence num in DB
	var seq int
	seq = database.CoutSeq(ctx)
	txnID := _const.TASK_TXN_PREFFIX + fmt.Sprintf("%d", seq)
	fragments := fragments{data: make(map[int]fragment)}
	counter := _const.DEFAULT_SUM_VALUE
	for site, config := range conf.Sites {
		for _, single := range config.Single {
			f := fragment{}
			f.single = types.Single{
				Type: single.Type,
				Rule: single.Rule,
				Key:  single.Key,
			}
			f.motherSite = site
			f.taken = false
			fragments.data[counter] = f
			counter ++
		}
	}
	fragments.sum = counter
	return &task{fragments: fragments, txID: txnID, schedule: make(map[int]bool), timeout: time.Duration(timeout) * time.Second}
}

func PopTask() *task {
	// this would initialize memory in cache layer
	// it's better to do it in PopTask cause it's dynamic
	task := <-Panel.queue
	cacheCenter.Lock()
	defer cacheCenter.Unlock()
	cacheCenter.data[task.txID] = &cacheUnit{
		ch:    make(chan interface{}),
		sum:   task.fragments.sum,
		unit:  make(map[int]unit),
		ready: false,
	}
	return task
}

func (t *task) TransactionID() string {
	return t.txID
}

// important core logic unit that contains fault tolerance
/*
			   pip ---+ task +------- finish
				+		  |				+
				|		  |				|
				|		  +				|
				|		fragment ---+ done
				|						|
				|						|
				|						+
				+----------------- in progress
*/
func (t *task) RunPip(pip *pipe) {
	pip.timeout = t.timeout
	// forbid to run unmatched pipe
	if pip.txnID != t.txID {
		return
	}
	// quit tag marking if the task has been done
	var quit bool
	// fetch a fragment that are free to be load
	for i := 0; i < t.fragments.sum; i++ {
		if t.fragments.data[i].taken {
			continue
		}
		pip.fragmentSeq = i
		mid := t.fragments.data[i]
		mid.taken = true
		t.fragments.data[i] = mid
		quit = false
		break
	}
	if quit {
		pip.state = _const.PIP_DONE
		// it's safety to check done signal here
		t.check()
		return
	}
	// once been taken run it in using colly , and save the result to cache layer
	go func() {
		// synchronised fetching data
		tar, _ := t.fetchFragment(pip.fragmentSeq)
		// inside run it's a synchronised processing
		pip.run(tar)
		// it's not over yet this pip would fetch a sequenced task then
		t.RunPip(pip)
	}()
}

// return the copy of single and if it's been taken by other pipe
func (t *task) fetchFragment(i int) (frag fragment, taken bool) {
	return t.fragments.data[i], t.fragments.data[i].taken
}

func (t *task) check() {
	t.RLock()
	defer t.RUnlock()
	PipPool.Lock()
	txID := t.txID
	var tag bool
	for seq, _ := range t.schedule {
		if PipPool.pips[seq].state == _const.PIP_DONE {
			tag = true
			break
		}
	}
	PipPool.Unlock()
	if !tag {
		cacheCenter.Lock()
		close(cacheCenter.data[txID].ch)
		cacheCenter.Unlock()
	}
}

func (p *pipe) run(single fragment) {
	log.Debug("pip ", p.fragmentSeq, " begin to run ", p.txnID)
	// in using colly as core lib in fetching elements
	// this for showing configs
	url := single.motherSite
	path := single.single.Rule
	typ := single.single.Type
	key := single.single.Key
	// result data
	var jsonRes []byte
	var htmlRes []string
	var err error
	// using colly handler in handler package
	switch typ {
	case _const.TYPE_HTML:
		htmlRes, err = handler.MatchHTML(url, path, p.fragmentSeq, p.timeout)
	case _const.TYPE_JSON:
		jsonRes, err = handler.MatchJSON(url, path, p.fragmentSeq, p.timeout)
	}
	fragSeq := p.fragmentSeq
	pipSeq := p.pipSeq
	if err != nil {
		CachingFailure(fragSeq, pipSeq, key, p.txnID, typ)
	}
	Caching(fragSeq, pipSeq, key, p.txnID, typ, combine(jsonRes, htmlRes))
}

func combine(bytes []byte, strings []string) interface{} {
	if len(bytes) == _const.DEFAULT_SUM_VALUE {
		return strings
	}
	return bytes
}
