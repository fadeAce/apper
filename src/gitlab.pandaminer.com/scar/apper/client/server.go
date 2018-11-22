package client

import (
	"github.com/nats-io/go-nats"
	"gitlab.pandaminer.com/scar/apper/const"
	"gitlab.pandaminer.com/scar/apper/types"
	"gitlab.pandaminer.com/scar/apper/logger"
	"encoding/json"
	"gitlab.pandaminer.com/scar/apper/core"
)

var log = logger.Log

var sum = _const.DEFAULT_SUM_VALUE

var conn *nats.Conn

var closeCh chan interface{}

// Listen takes advantage of broker to make a connection between client-sdk and
// apper server.
// sdk - broker - server - distribute - cushion - exec - pool - handler - broker - sdk
func Listen(ch chan interface{}, apper *types.Apperserver) (err error) {
	closeCh = ch
	log.Debug("begin NATS connection")
	nc, err := nats.Connect(nats.DefaultURL)
	conn = nc
	if err != nil {
		return err
	}
	// use Distribute method to finish callback
	nc.Subscribe("cmd", func(msg *nats.Msg) {
		cmd := types.Command{}
		json.Unmarshal(msg.Data, &cmd)
		switch cmd.Cmd {
		case _const.CMD_START:
			// add task to cushion
			sitemap := cmd.Configs
			// todo: sitemap - test
			task := core.Generate(apper.Ctx, sitemap, apper.Database, apper.Cfg.Timeout)
			// todo: Pending
			core.Panel.Init(apper.Cfg)
			core.Panel.Pending(task)
		case _const.CMD_STOP:
		case _const.CMD_LS:
		}
	})
	return
}
