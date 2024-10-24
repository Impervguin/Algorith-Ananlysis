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
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("LOCAL_POSTGRES_HOST"), os.Getenv("POSTGRES_DB"))
}

func NewPgsStorage(ctx context.Context) (*PgsStorage, error) {
	// fmt.Println(DSN())
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
