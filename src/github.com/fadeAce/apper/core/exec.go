package core

import (
	"github.com/fadeAce/apper/const"
	"github.com/fadeAce/apper/storage"
	"context"
)

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
func (t *task) Store(ctx context.Context, database *storage.Database) (m map[string]string, err error) {
	txID := t.txID
	cacheCenter.Lock()
	data := cacheCenter.data[txID]
	cacheCenter.Unlock()
	count := _const.DEFAULT_SUM_VALUE
	for fragSeq, val := range data.unit {
		log.Info(txID, " ", fragSeq, " is persisted to database")
		// persist single value to database
		if val.state == _const.CACHING_STATE_ERROR {
			m[val.key] = "failure key in fetching data"
		}
		err = database.InsertFragment(
			ctx, txID, val.typ, val.key, val.state, val.val)
		count++
	}
	tag := data.sum == count
	if !tag {
		log.Warn("task ", txID, " is not well persisted")
	}
	return
}
