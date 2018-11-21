package core

import "gitlab.pandaminer.com/scar/apper/const"

// this package unit is for executors to exec pipes
// then assemble all data and make a transaction to store them

// important interface for starting s pool
func StartPool(sum int) {
	// create pipes step.1
	for i := 0; i < sum; i++ {
		p := &pipe{pipSeq: i, state: _const.PIP_IDLE}
		PipPool.addPip(p)
	}
}

// when task is attempt to store a assemble
// it will fetch them from dataAssemble by pipes it's currently using
// then release all pipes after store finished
func (*task) store() {

}
