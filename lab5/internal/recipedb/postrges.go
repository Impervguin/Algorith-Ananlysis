package recipedb

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx"
)

const (
	maxConn        = 10
	acquireTimeout = time.Minute
)

type PgsStorage struct {
	conn *pgx.ConnPool
}

func DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:5432/%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
}

func NewPgsStorage(ctx context.Context) (*PgsStorage, error) {
	conf, err := pgx.ParseURI(DSN())

	if err != nil {
		return nil, err
	}
	poolConf := &pgx.ConnPoolConfig{
		ConnConfig:     conf,
		MaxConnections: maxConn,
		AcquireTimeout: acquireTimeout,
	}
	conn, err := pgx.NewConnPool(*poolConf)
	if err != nil {
		return nil, err
	}

	return &PgsStorage{conn: conn}, nil
}
