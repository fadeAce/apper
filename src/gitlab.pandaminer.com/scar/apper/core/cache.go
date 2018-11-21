package core

import (
	"sync"
)

var cacheCenter = func() cache {
	return cache{data: make(map[string]cacheUnit)}
}()

type cache struct {
	sync.RWMutex
	// txnID - cacheUnit
	data map[string]cacheUnit
}

type unit struct {
	state   int
	pipeSeq int
	typ     int
	key     string
	val     interface{}
}

type cacheUnit struct {
	// fragmentSum
	sum int
	// fragmentSeq - unit
	unit map[int]unit
	// txn state isReady
	ready bool
}

func Caching(key, txnID, typ string, value interface{}) {
	// txID - fragmentID - data
	// caching for consumer to persis as a unit
	cacheCenter.Lock()
	defer cacheCenter.Unlock()
	cacheCenter.data[txnID] = cacheUnit{

	}
	// persis to postgres db
}

func CachingFailure(s string, s2 string, s3 string) {

}
