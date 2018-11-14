package types

import (
	"gitlab.pandaminer.com/scar/apper/storage"
	"sync"
	"golang.org/x/net/context"
)

type CmdJSON struct {
	Cmd   string            `json:"cmd"`
	Param map[string]string `json:"param"`
}

type Apperserver struct {
	sync.RWMutex
	Cfg      *ApperConf
	Ctx      context.Context
	Database *storage.Database
}

func (apperserver *Apperserver) Seq() int {
	apperserver.Lock()
	res := apperserver.Database.CoutSeq(apperserver.Ctx)
	apperserver.Unlock()
	return res
}
