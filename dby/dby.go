package dby

import (
	"database/sql"
	"fmt"
)

func NewDbConn(dbType string, connStr string) *sql.DB {
	conn, err := sql.Open(dbType, connStr)
	if err != nil {
		panic(err)
	}
	return conn
}

func Ping(conn *sql.DB) {
	err := conn.Ping()
	if err != nil {
		fmt.Println("dby:conn.Ping failed")
		panic(err)
	}
}
