package core

import (
	"gitlab.pandaminer.com/scar/apper/logger"
	"sync"
	"gitlab.pandaminer.com/scar/apper/types"
	"fmt"
	"gitlab.pandaminer.com/scar/apper/const"
	"gitlab.pandaminer.com/scar/apper"
)

var log = logger.Log

// pending centre for tasks
var Panel = func() *pendCentre {
	pc := pendCentre{}
	return &pc
}()

var dataAssemble = func() *dataCentre {
	dc := &dataCentre{}
	return dc
}()

// type task represent a task that match a configuration
type task struct {
	fragments fragments
	txID      string
	pips      map[int]*pipe
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

// type dataCentre represent the result assembled of all task
type dataCentre struct {
	sync.RWMutex
	dataAssemble map[*pipe]*DataUnit
}

type DataUnit struct {
	Flag bool
}

type executor interface {
	run() map[*pipe]*DataUnit
}

// pending tasks onto the panel
func (p *pendCentre) Pending(tsk *task) {
	if p.queue == nil {
		p.queue = make(chan *task, apper.Apper.Cfg.CushionSize)
	}
	p.queue <- tsk
}

// pop up a task
func (*pendCentre) pop() (*task) {
	return nil
}

// generate task from configuration
func Generate(conf types.ConfJ) *task {
	// step.1 generate txnID
	// seq as a sequence num in DB
	var seq int
	seq = apper.Apper.Seq()
	txnID := _const.TASK_TXN_PREFFIX + fmt.Sprintf("%d", seq)
	fragments := fragments{data: make(map[int]fragment)}
	counter := _const.DEFAULT_SUM_VALUE
	for site, config := range conf.Sites {
		for _, single := range config.Single {
			f := fragment{}
			f.single = types.Single{single.Type, single.Rule, single.Key}
			f.motherSite = site
			f.taken = false
			fragments.data[counter] = f
			counter ++
		}
	}
	fragments.sum = counter
	return &task{fragments, txnID, nil}
}

func PopTask() *task {
	return nil
}

func (t *task) TransactionID() string {
	return t.txID
}

func (t *task) RunPip(pip *pipe) {

}
