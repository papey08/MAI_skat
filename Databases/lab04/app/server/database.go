package server

import "database/sql"

var db *sql.DB

func OpenDB(configStr string) error {
	var err error
	db, err = sql.Open("postgres", configStr)
	return err
}

func CloseDB() error {
	return db.Close()
}
