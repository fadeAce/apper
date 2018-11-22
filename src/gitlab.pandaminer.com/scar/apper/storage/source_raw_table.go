package storage

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
)

const sourceSchema = `
CREATE SEQUENCE IF NOT EXISTS source_stream_id;

CREATE SEQUENCE IF NOT EXISTS transaction_stream_id;

CREATE TABLE IF NOT EXISTS source_raw (
	id 					BIGINT 	PRIMARY KEY DEFAULT nextval('source_stream_id'),
	transaction_id 		TEXT 	NOT NULL,
	typ 				TEXT 	NOT NULL,
	key_str 			TEXT 	NOT NULL,
	val_str 			TEXT 	NOT NULL,
	val_binary 			bytea 	NOT NULL,
	intime 				BIGINT 	NOT NULL,
	state    			TEXT 	NOT NULL,
	des 				TEXT    NOT NULL
);

CREATE INDEX IF NOT EXISTS transaction_index ON source_raw(transaction_id);
`

const insertSourceSQL = "" +
	"INSERT INTO source_raw(transaction_id, typ, key_str, val_str, val_binary, intime, state, des) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"

const selectSourceByUIDSQL = "" +
	"SELECT val_str FROM source_raw WHERE transaction_id = $1 and typ = $2"

const selectAllSQL = "" +
	"SELECT val_str, typ FROM source_raw where transaction_id = $1"

const selectTxnSeqSQL = "" +
	"SELECT nextval('transaction_stream_id')"

type scrapperStatements struct {
	insertRawStmt         *sql.Stmt
	selectSourceByUIDStmt *sql.Stmt
	selectAllSourceStmt   *sql.Stmt
	selectTxnSeqStmt      *sql.Stmt
}

func (s *scrapperStatements) getSchema() string {
	return sourceSchema
}

func (s *scrapperStatements) prepare(db *sql.DB) (err error) {
	_, err = db.Exec(sourceSchema)
	if err != nil {
		return
	}
	if s.insertRawStmt, err = db.Prepare(insertSourceSQL); err != nil {
		return
	}
	if s.selectSourceByUIDStmt, err = db.Prepare(selectSourceByUIDSQL); err != nil {
		return
	}
	if s.selectAllSourceStmt, err = db.Prepare(selectAllSQL); err != nil {
		return
	}
	if s.selectTxnSeqStmt, err = db.Prepare(selectTxnSeqSQL); err != nil {
		return
	}

	return
}

func (s *scrapperStatements) insertFragment(
	ctx context.Context,
	txID, typ, key, cont string,
	bi []byte,
	intime int64,
	state string,
	des string,
) error {
	bt := pq.ByteaArray{}
	err := bt.Scan(bi)
	if err != nil {
		return err
	}
	if _, err := s.insertRawStmt.ExecContext(ctx, txID, typ, key, cont, bt, intime, state, des); err != nil {
		return nil
	}
	return nil
}
func (s *scrapperStatements) selectTxnSeq(ctx context.Context) int {
	var res int
	s.selectTxnSeqStmt.QueryRowContext(ctx).Scan(&res)
	return res
}
