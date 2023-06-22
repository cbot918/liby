package main

import (
	parsery "github.com/cbot918/liby/parsery/internal"
)

const (
	file = "./data/arith.txt"
)

func main() {
	parsery.New(file).Run()
}
