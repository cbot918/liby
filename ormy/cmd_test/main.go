package main

import (
	"fmt"
	"log"
	"ormy"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Name string
}

func main() {
	engine, _ := ormy.NewEngine("sqlite3", "ormy.db")
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success %d affected\n", count)
	row := s.Raw("select * from User where Name = (?)", "Tom").QueryRow()
	var u User
	if err := row.Scan(&u.Name); err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)
}
