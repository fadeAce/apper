package core

import (
	"gitlab.pandaminer.com/scar/apper/logger"
	"sync"
	"gitlab.pandaminer.com/scar/apper/types"
)

var log = logger.Log

// pending centre for tasks
var panel = func() *pendCentre {
	pc := pendCentre{}
	return &pc
}()

var dataAssemble = func() *dataCentre {
	dc := &dataCentre{}
	return dc
}()

// type task represent a task that match a configuration
type task struct {
}

// core concept of cushion area for tasks
type pendCentre struct {
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
func (*pendCentre) pending(task *task) {

}

// pop up a task
func (*pendCentre) pop() (*task) {
	return nil
}

// generate task from configuration
func generate(conf types.ConfJ) *task {

	return nil
}
