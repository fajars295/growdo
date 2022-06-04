package componen

import (
	"context"
	"database/sql"
	"log"
	"time"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"

type querylogslow struct {
	dbTx  *sql.Tx
	db    *sql.DB
	ctx   context.Context
	query string
}

func NewQueryLogSlow(db *sql.DB, dbTx *sql.Tx, ctx context.Context, q string) *querylogslow {
	return &querylogslow{
		db:    db,
		dbTx:  dbTx,
		ctx:   ctx,
		query: q,
	}
}

func (r *querylogslow) Create() (int64, error) {

	firstQueryStart := time.Now()
	var id int64
	errors := r.dbTx.QueryRow(r.query).Scan(&id)
	firstQueryEnd := time.Now()
	if errors != nil {
		log.Printf("%s %s => %s  %s \n", Red, r.query, firstQueryEnd.Sub(firstQueryStart).String(), Reset)
		return 0, errors
	}
	log.Printf("%s => %s \n", r.query, firstQueryEnd.Sub(firstQueryStart).String())
	return id, nil
}

func (r *querylogslow) Get() (*sql.Rows, error) {

	firstQueryStart := time.Now()
	rows, err := r.db.QueryContext(r.ctx, r.query)
	firstQueryEnd := time.Now()

	if err != nil {
		log.Printf("%s query #%s => %s  %s \n", Red, r.query, firstQueryEnd.Sub(firstQueryStart).String(), Reset)
		return nil, err
	}
	log.Printf("%s %s => %s %s\n", Green, r.query, firstQueryEnd.Sub(firstQueryStart).String(), Reset)
	return rows, nil
}
