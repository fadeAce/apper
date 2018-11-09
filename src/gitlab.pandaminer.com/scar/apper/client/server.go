package client

import (
	"github.com/nats-io/go-nats"
	"gitlab.pandaminer.com/scar/apper/const"
	"gitlab.pandaminer.com/scar/apper/types"
	"gitlab.pandaminer.com/scar/apper/logger"
	"gitlab.pandaminer.com/scar/apper/core"
)

var log = logger.Log

var sum = _const.DEFAULT_SUM_VALUE

var conn *nats.Conn


// Listen takes advantage of broker to make a connection between client-sdk and
// apper server.
// sdk - broker - server - distribute - cushion - exec - pool - handler - broker - sdk
func Listen(ch chan interface{}, apper *types.Apperserver) (err error) {
	log.Debug("begin NATS connection")
	nc, err := nats.Connect(nats.DefaultURL)
	conn = nc
	if err != nil {
		return err
	}
	// use Distribute method to finish callback
	nc.Subscribe("cmd", core.Distribute)
	// wait for signal to kill
	<-ch
	return
}

