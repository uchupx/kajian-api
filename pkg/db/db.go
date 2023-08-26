package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/uchupx/kajian-api/pkg/logger"
)

type DB struct {
	*sqlx.DB
	isDebug bool
}

type txKey struct{}

type Stmt struct {
	*sqlx.Stmt
	RawQuery string
}

type logMeta struct {
	Query    string `json:"query"`
	ExecTime string `json:"exec_time"`
}

func (g *DB) FTransaction(ctx context.Context, fc func(c context.Context, tx *sqlx.Tx) error, opts *sql.TxOptions) error {
	tx, err := g.BeginTxx(ctx, opts)
	if err != nil {
		return err
	}

	c := g.injextTx(ctx, *tx)

	defer func() {
		if r := recover(); r != nil {
			// fileLine := errors.TracePanic()

			// meta := errors.ErrorMeta{
			// 	Message: "FAILED TRANSACTION",
			// 	Line:    fileLine,
			// 	IsPanic: true,
			// }

			// logger.Logger.WithFields(logrus.Fields{
			// 	"meta": meta, // Log response body
			// }).Error(r)

			tx.Rollback()
		}
	}()

	err = fc(c, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (g *DB) injextTx(ctx context.Context, tx sqlx.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

func (g *DB) extractTx(ctx context.Context) *sqlx.Tx {
	if val, ok := ctx.Value(txKey{}).(sqlx.Tx); ok {
		return &val
	}

	return nil
}

func (g *DB) SetDebug(isDebug bool) {
	g.isDebug = isDebug
}

func (g *DB) FPreparexContext(ctx context.Context, query string) (*Stmt, error) {
	var err error
	var stmt *sqlx.Stmt

	if g.isDebug {
		logger.Logger.Infof(query)
	}

	tx := g.extractTx(ctx)
	if tx != nil {
		stmt, err = tx.PreparexContext(ctx, query)
	} else {
		stmt, err = g.PreparexContext(ctx, query)
	}

	return &Stmt{Stmt: stmt, RawQuery: query}, err
}

func (g *Stmt) FExecContext(ctx context.Context, args ...interface{}) (res sql.Result, err error) {
	t := time.Now()
	res, err = g.ExecContext(ctx, args...)
	if err != nil {
		return res, err
	}

	latency := time.Since(t)
	slowQuery(t, g.RawQuery, latency)

	return res, nil
}

func (g *Stmt) FQueryRowxContext(ctx context.Context, args ...interface{}) (res *sqlx.Row) {
	t := time.Now()

	res = g.QueryRowxContext(ctx, args...)

	latency := time.Since(t)
	slowQuery(t, g.RawQuery, latency)

	return res
}

func (g *Stmt) FQueryxContext(ctx context.Context, args ...interface{}) (res *sqlx.Rows, err error) {
	t := time.Now()
	res, err = g.QueryxContext(ctx, args...)
	if err != nil {
		return res, err
	}

	slowQuery(t, g.RawQuery, time.Since(t))
	return res, nil
}

func slowQuery(t time.Time, query string, latency time.Duration) {
	if latency.Milliseconds() > 200 {
		logger.Logger.WithFields(logrus.Fields{
			"meta": logMeta{
				Query:    query,
				ExecTime: fmt.Sprintf("%d ms", latency.Milliseconds()),
			},
		}).Warn("slow query")
	}
}
