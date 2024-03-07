package database

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	conn *pgxpool.Pool
)

type Pgx interface {
	Exec(context context.Context, sql string, arguments ...interface{}) (commandTag pgconn.CommandTag, err error)
	Query(context context.Context, sql string, args ...interface{}) (rows pgx.Rows, err error)
	QueryRow(context context.Context, sql string, args ...interface{}) pgx.Row
	Begin(context context.Context) (pgx.Tx, error)
	Close()
}

func Connect() Pgx {
	if conn != nil {
		return conn
	}

	err := error(nil)
	conn, err = pgxpool.New(context.Background(), "user=admin password=oot123 dbname=rinha host=localhost port=5432")

	if err != nil {
		panic(err)
	}

	conn.Config().MaxConns = 50
	conn.Config().MinConns = 40
	conn.Config().MaxConnIdleTime = time.Minute * 3

	if err != nil {
		panic(err)
	}

	log.Print("$$$$Connected to database")
	return conn
}
