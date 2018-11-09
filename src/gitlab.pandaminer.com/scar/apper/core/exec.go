package core

import (
	"gitlab.pandaminer.com/scar/apper/types"
	"gitlab.pandaminer.com/scar/apper/const"
	"golang.org/x/net/context"
	"fmt"
)

// this package unit is for executors to exec pipes
// then assemble all data and make a transaction to store them

// important interface for starting s pool
func StartPool(sum int, notifier *types.Notifier) {
	// create pipes step.1
	for i := 0; i < _const.DEFAULT_SUM_PIPE; i++ {
		p := &pipe{}
		PipPool.addPip(p)
	}

	// create cushion consumer step.2
	daemonCtx := context.Background()
	ctx, cancel := context.WithCancel(daemonCtx)
	go func(ctx context.Context) {

		ctx.Done()
	}(ctx)
	fmt.Print(cancel)
}

func (*task) run() map[*pipe]*DataUnit {
	return nil
}

// when task is attempt to store a assemble
// it will fetch them from dataAssemble by pipes it's currently using
// then release all pipes after store finished
func (*task) store() {

}
