package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"gitlab.pandaminer.com/scar/apper/logger"
)

var log = logger.Log

const (
	ScrapperDBDriver = "postgres"
)

type Database struct {
	db       *sql.DB
	scrapper scrapperStatements
}

// NewDatabase creates a new presence database
func NewDatabase(dataSourceName string) (*Database, error) {
	var d Database
	var err error
	if d.db, err = sql.Open(ScrapperDBDriver, dataSourceName); err != nil {
		return nil, err
	}
	if err = d.scrapper.prepare(d.db); err != nil {
		fmt.Println(err)
		return nil, err
	}
	log.Info("complete storage setting")
	return &d, nil
}
