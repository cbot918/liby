package main

import "fmt"

type Token int

const (
	EOF Token = iota
	ADD
	SUB
	MUL
	QUO
)

var tokens = [...]string{
	ADD: "ADD",
}

var log = fmt.Println

func main() {
	log(tokens[1])

	const k = 0
	t := [...]string{
		k: "b",
	}
	log(t[0])
}
