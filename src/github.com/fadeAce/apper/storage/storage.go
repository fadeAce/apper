package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/fadeAce/apper/logger"
	"context"
	"time"
	"github.com/fadeAce/apper/const"
	"strings"
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
		return nil, err
	}
	log.Info("complete storage setting")
	return &d, nil
}

func (d *Database) CoutSeq(ctx context.Context) (seq int) {
	seq = d.scrapper.selectTxnSeq(ctx)
	return
}

func (d *Database) InsertFragment(
	ctx context.Context, txID, typ, key string, state int,
	val interface{},
) (err error) {
	statestr := "finish"
	if state == _const.CACHING_STATE_ERROR {
		statestr = "error"
	}
	intime := time.Now().Unix()
	if typ == _const.TYPE_HTML {
		restr := val.([]string)
		res := strings.Join(restr, ",")
		err = d.scrapper.insertFragment(ctx, txID, typ, key, res, []byte{}, intime, statestr, "")
		return
	}
	err = d.scrapper.insertFragment(ctx, txID, typ, key, "", val.([]byte), intime, statestr, "")
	return
}
