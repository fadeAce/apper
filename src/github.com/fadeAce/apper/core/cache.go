package core

import (
	"sync"
	"github.com/fadeAce/apper/const"
)

var cacheCenter = func() cache {
	return cache{data: make(map[string]*cacheUnit)}
}()

type cache struct {
	sync.RWMutex
	// txnID - cacheUnit
	data map[string]*cacheUnit
}

type unit struct {
	state   int
	pipeSeq int
	typ     string
	key     string
	val     interface{}
}

type cacheUnit struct {
	// cache chan
	ch chan interface{}
	// fragmentSum
	sum int
	// fragmentSeq - unit
	unit map[int]unit
	// txn state isReady
	ready bool
}

func Caching(fragSeq int, pipSeq int, key, txnID, typ string, value interface{}) {
	// txID - fragmentID - data
	// caching for consumer to persis as a unit
	cacheCenter.Lock()
	defer cacheCenter.Unlock()
	mid := cacheCenter.data[txnID]
	mid.unit[fragSeq] = unit{
		state:   _const.CACHING_STATE_NORMAL,
		pipeSeq: pipSeq,
		typ:     typ,
		key:     key,
		val:     value,
	}
	cacheCenter.data[txnID] = mid
}

func CachingFailure(fragSeq int, pipSeq int, key, txnID, typ string) {
	cacheCenter.Lock()
	defer cacheCenter.Unlock()
	mid := cacheCenter.data[txnID]
	mid.unit[fragSeq] = unit{
		state:   _const.CACHING_STATE_ERROR,
		pipeSeq: pipSeq,
		typ:     typ,
		key:     key,
		val:     nil,
	}
	cacheCenter.data[txnID] = mid
}
