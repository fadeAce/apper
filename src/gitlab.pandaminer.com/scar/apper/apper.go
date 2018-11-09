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

func loadNewApper() (err error, apper *typ.Apperserver) {
	apper = &typ.Apperserver{}
	return
}

func Start(conf *typ.ApperConf) error {
	// prepare configs for new daemon
	err, apper := loadNewApper()
	if err != nil {
		return err
	}
	storage.NewDatabase(conf.Database)

	// start listener 					-- step.1
	shutCh := make(chan interface{})
	err = client.Listen(shutCh, apper)

	// start notifier                   -- step.2
	notifier := client.NewNotifier()

	// start Executors cushion and pool -- step.3
	core.StartPool(_const.DEFAULT_SUM_PIPE, notifier)

	return err
}
