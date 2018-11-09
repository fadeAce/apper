package storage

import (
	"context"
	"database/sql"
)

const sourceSchema = `
CREATE SEQUENCE IF NOT EXISTS source_stream_id;

CREATE TABLE IF NOT EXISTS source_raw (
	id 					BIGINT 	PRIMARY KEY DEFAULT nextval('source_stream_id'),
	transaction_id 		TEXT 	NOT NULL,
	typ 				TEXT 	NOT NULL,
	content 			TEXT 	NOT NULL,
	intime 				BIGINT 	NOT NULL
);

CREATE INDEX IF NOT EXISTS transaction_index ON source_raw(transaction_id);
`

const insertSourceSQL = "" +
	"INSERT INTO source_raw(transaction_id, typ, content, intime) VALUES ($1, $2, $3, $4)"

const selectSourceByUIDSQL = "" +
	"SELECT content FROM source_raw WHERE transaction_id = $1 and typ = $2"

const selectAllSQL = "" +
	"SELECT content, typ FROM source_raw where transaction_id = $1"

type scrapperStatements struct {
	insertRawStmt         *sql.Stmt
	selectSourceByUIDStmt *sql.Stmt
	selectAllSourceStmt   *sql.Stmt
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

	return
}

func (s *scrapperStatements) insertPresence(
	ctx context.Context, tid, typ, cont string,
	intime int64,
) error {
	if _, err := s.insertRawStmt.ExecContext(ctx, tid, typ, cont, intime); err != nil {
		return nil
	}
	return nil
}

//
//func (s *scrapperStatements) selectAll() ([]interface{}, error) {
//	rows, err := s.selectAllSourceStmt.Query()
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//	var result []interface{}
//	//for rows.Next() {
//	//	var presence authtypes.Presence
//	//	if err = rows.Scan(&presence.UserID, &presence.State, &presence.StatusMsg, &presence.Mtime); err != nil {
//	//		return nil, err
//	//	}
//	//	result = append(result, presence)
//	//}
//	return result, err
//}
