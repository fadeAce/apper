package apper

import (
	"gitlab.pandaminer.com/scar/apper/client"
	"gitlab.pandaminer.com/scar/apper/storage"
	typ "gitlab.pandaminer.com/scar/apper/types"
	"gitlab.pandaminer.com/scar/apper/logger"
	"gitlab.pandaminer.com/scar/apper/core"
	"gitlab.pandaminer.com/scar/apper/const"
)

var log = logger.Log

var Apper *typ.Apperserver

func loadNewApper() (err error, apper *typ.Apperserver) {
	apper = &typ.Apperserver{}
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
	// start listener todo: WIP					 -- step.1
	shutCh := make(chan interface{})
	err = client.Listen(shutCh, apper)

	// start notifier todo: WIP                  -- step.2
	notifier := typ.NewNotifier()

	// start Executors 			cushion and pool -- step.3
	var ths int
	if Apper.Cfg.ThreadPoolSize > 0 {
		ths = Apper.Cfg.ThreadPoolSize
	} else {
		ths = _const.DEFAULT_SUM_PIPE
	}
	core.StartPool(ths)

	// start pool consuming service todo: WIP    -- step.4
	StartService(notifier)

	return err
}

func StartService(notifier *typ.Notifier) {

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
		// block till it's done
		txID := task.Done()
		log.Info("task ", txID, " is done !")
	}

	//
	//// collect data to cache & persist
	//
	////
	//daemonCtx := context.Background()
	//ctx, cancel := context.WithCancel(daemonCtx)
	//go func(ctx context.Context) {
	//	// manufacture processing routine
	//	ctx.Done()
	//}(ctx)
	//fmt.Print(cancel)
}
