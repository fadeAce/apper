package apper

import (
	"github.com/fadeAce/apper/client"
	"github.com/fadeAce/apper/storage"
	typ "github.com/fadeAce/apper/types"
	"github.com/fadeAce/apper/logger"
	"github.com/fadeAce/apper/core"
	"github.com/fadeAce/apper/const"
	"context"
	"encoding/json"
)

var log = logger.Log

var Apper *typ.Apperserver

func loadNewApper() (err error, apper *typ.Apperserver) {
	apper = &typ.Apperserver{}
	ctx := context.Background()
	apper.Ctx, apper.Quit = context.WithCancel(ctx)
	return
}

func Start(conf *typ.ApperConf) error {
	// prepare configs for new daemon
	err, apper := loadNewApper()
	apper.Cfg = conf
	if err != nil {
		return err
	}
	database, err := storage.NewDatabase(apper.Cfg.Database)
	if err != nil {
		return err
	}
	apper.Database = database
	Apper = apper
	// start listener 							 -- step.1
	shutCh := make(chan interface{})
	err = client.Listen(shutCh, apper)
	if err != nil {
		return err
	}
	// start notifier 			                 -- step.2
	notifier := typ.NewNotifier()

	// start Executors cushion and pool 		 -- step.3
	var ths int
	if Apper.Cfg.ThreadPoolSize > 0 {
		ths = Apper.Cfg.ThreadPoolSize
	} else {
		ths = _const.DEFAULT_SUM_PIPE
	}
	core.StartPool(ths)

	// start pool consuming service 			 -- step.4
	StartService(apper.Ctx, notifier, conf, database)

	return err
}

func StartService(
	ctx context.Context, notifier *typ.Notifier, conf *typ.ApperConf, database *storage.Database,
) {
	// loop
	for {
		if false {
			return
		}
		// pip only been created in core package itself with a registered task.
		// no other pips are exposed with public method.
		// so does task it's protected by Poptask() no other approaches except
		// StartPool itself.
		task := core.PopTask()
		pips := task.MatchPIP()
		for idx, pip := range pips {
			task.RunPip(pip)
			log.Debug("pip ", idx, " is running for task ", task.TransactionID())
		}
		go func() {
			// block till it's done, this is async block daemon process for task
			txID := task.Done()
			// done marks the finish of task , and beginning of storage.999
			failure, err := task.Store(ctx, database)
			if err != nil {
				return
			}
			// notify this work is over through the unified subject.
			fail := typ.Respond{
				State:  false,
				Falure: failure,
			}
			data, _ := json.Marshal(fail)
			notifier.Notify(_const.TASK_TXN_PREFFIX+txID, data)
			log.Info("task ", txID, " is done !")
		}()
	}
}
