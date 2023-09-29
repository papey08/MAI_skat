package server

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

// OpenDB is a wrapper of sqlx.Open
func OpenDB(configStr string) error {
	var err error
	db, err = sqlx.Open("postgres", configStr)
	return err
}

// CloseDB closes database
func CloseDB() error {
	return db.Close()
}
