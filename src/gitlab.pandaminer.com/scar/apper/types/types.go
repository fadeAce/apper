package types

import (
	"gitlab.pandaminer.com/scar/apper/storage"
	"sync"
	"context"
	"github.com/nats-io/go-nats"
)

type CmdJSON struct {
	Cmd   string            `json:"cmd"`
	Param map[string]string `json:"param"`
}

type Apperserver struct {
	Quit func()
	sync.RWMutex
	Cfg      *ApperConf
	Ctx      context.Context
	Database *storage.Database
	nc       *nats.Conn
}

func (apperserver *Apperserver) Seq() int {
	apperserver.Lock()
	res := apperserver.Database.CoutSeq(apperserver.Ctx)
	apperserver.Unlock()
	return res
}

func (apperserver *Apperserver) Close() {
	apperserver.Lock()
	apperserver.nc.Close()
	apperserver.Unlock()
}
