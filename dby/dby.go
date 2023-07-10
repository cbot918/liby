package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DBy struct {
	DB *sql.DB
}

func NewDBy(dbType string, url string) (*DBy, error) {
	db, err := sql.Open(dbType, url)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &DBy{
		DB: db,
	}, nil
}

func (db *DBy) Exec(q string) (sql.Result, error) {
	result, err := db.DB.Exec(q)
	return result, err
}
