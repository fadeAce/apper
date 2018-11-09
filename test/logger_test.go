package test

import (
	"github.com/sirupsen/logrus"
	"os"
	"github.com/lestrrat/go-file-rotatelogs"
	"time"
	"github.com/rifflock/lfshook"
	"testing"
)

var log *logrus.Logger

func init() {
	initLog()
}

func initLog() {
	log = logrus.New()
	log.Formatter = &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "01-02 15:04:05.000000",
	}
	log.WithFields(logrus.Fields{"package": "aaa"})
}

func setDebug() {
	log.SetLevel(logrus.DebugLevel)

	os.Mkdir("logs", os.ModePerm)
	debugLogPath := "logs/debug.log"
	warnLogPath := "logs/warn.log"
	debugLogWriter, err := rotatelogs.New(
		debugLogPath+".%Y%m%d",
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		log.Printf("failed to create rotatelogs debugLogWriter : %s", err)
		return
	}
	warnLogWriter, err := rotatelogs.New(
		warnLogPath+".%Y%m%d",
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		log.Printf("failed to create rotatelogs warnLogWriter : %s", err)
		return
	}
	log.Hooks.Add(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.DebugLevel: debugLogWriter,
			logrus.InfoLevel:  debugLogWriter,
			logrus.WarnLevel:  warnLogWriter,
			logrus.ErrorLevel: warnLogWriter,
			logrus.FatalLevel: warnLogWriter,
			logrus.PanicLevel: warnLogWriter,
		},
		&logrus.JSONFormatter{},
	))
}

func Test_logger(t *testing.T) {
	initLog()
	setDebug()
	log.Debug("aaaa")
	log.Warn("bbb")
}
