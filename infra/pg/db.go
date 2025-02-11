package pg

import (
	"database/sql"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type DB struct {
	*bun.DB
}

func NewDB() (*DB, error) {
	db, err := newDB()
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func newDB() (*bun.DB, error) {
	dsn := "postgres://postgres:postgres@db:5432/sns?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	sqldb.SetMaxOpenConns(10)
	sqldb.SetMaxIdleConns(10)
	sqldb.SetConnMaxLifetime(time.Second * 60)
	db := bun.NewDB(sqldb, pgdialect.New())
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil

}
