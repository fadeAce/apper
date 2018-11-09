package main

import (
	sh "./shell"
	"flag"
	"gitlab.pandaminer.com/scar/apper"
	"gitlab.pandaminer.com/scar/apper/logger"
	"gitlab.pandaminer.com/scar/apper/types"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

var log = logger.Log

func main() {
	/*
		for apper server daemon
	*/
	ipc := flag.Bool("i", false, "whether to run apper in a interface mode")
	cfg := flag.String("f", "./apper.yaml", "config file for apper to run , when it's not in a interface mode ")
	flag.Parse()
	conf := types.ApperConf{}
	// in ipc mode
	if *ipc && *cfg == "./apper.yaml" {
		// manage with apper CLI
		log.Info("CLI activated at ", time.Now())
		shell := sh.InitShell()
		shell.Run()
	} else {
		log.Info("start apper server")
		// manage with apper daemon
		if *cfg != "" {
			d, err := ioutil.ReadFile(*cfg)
			if err != nil {
				log.Info(err)
			}
			err = yaml.Unmarshal([]byte(d), &conf)
			if err != nil {
				log.Info(err)
			}
			apper.Start(&conf)
		}
	}
}
