package types

import (
	"github.com/sirupsen/logrus"
	"github.com/fadeAce/apper/storage"
)

func (*Apperserver) Start() {

}

func (*Apperserver) Terminate() {

}

func (a *Apperserver) CreateApperDB() *storage.Database {
	db, err := storage.NewDatabase(string(a.Cfg.Database))
	if err != nil {
		logrus.WithError(err).Panicf("failed to connect to apper db")
	}
	return db
}
