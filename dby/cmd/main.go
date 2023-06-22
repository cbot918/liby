package main

import (
	"database/sql"

	"github.com/cbot918/liby/dby"
	_ "github.com/lib/pq"
)

const (
	// host     = "localhost"
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "12345"
	dbname   = "testdb"

	db_type = "postgres"

	web_port = ":5455"
)

type Config struct {
	Host     string
	Port     int32
	User     string
	Password string
	Dbname   string
	WebPort  string
}

func NewConfig() *Config {
	return &Config{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Dbname:   dbname,
		WebPort:  web_port,
	}
}

func main() {
	// c := NewConfig()
	// connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.Dbname)

	connStr := "postgres://postgres:12345@localhost:5433/testdb?sslmode=disable"
	conn, err := sql.Open(db_type, connStr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	dby.Ping(conn)

}
