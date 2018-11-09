package logger

import (
	logger "github.com/sirupsen/logrus"
	"time"
	"path"
	"github.com/pkg/errors"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"os"
)

var Log = func() *logger.Logger{
	log := logger.New()
	log.SetLevel(logger.DebugLevel)
	os.Mkdir("log", os.ModePerm)
	configLocalFilesystemLogger("./log", "apper", 7*24*60*60*time.Second, 24*60*60*time.Second, log)
	return log
}()

// config logrus log to local filesystem, with file rotation
func configLocalFilesystemLogger(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration, log *logger.Logger) {
	baseLogPaht := path.Join(logPath, logFileName)
	formalP := baseLogPaht + "_formal"
	formalW, err := rotatelogs.New(
		formalP+".%Y%m%d",
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)
	infoP := baseLogPaht + "_info"
	infoW, err := rotatelogs.New(
		infoP+".%Y%m%d",
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)
	errP := baseLogPaht + "_err"
	errW, err := rotatelogs.New(
		errP+".%Y%m%d",
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)
	if err != nil {
		logger.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logger.DebugLevel: formalW,
		logger.InfoLevel:  infoW,
		logger.WarnLevel:  infoW,
		logger.ErrorLevel: errW,
		logger.PanicLevel: errW,
	}, nil)
	log.Hooks.Add(lfHook)
}
