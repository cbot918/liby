package main

import (
	"log"
)

const (
	file = "poc/1.jsx"
)

func main() {
	c, err := NewHtmlJsxCompiler(file)
	if err != nil {
		log.Fatal(err)
	}
	lg(string(c.Content))
	_, _ = c.Tokenize(c.Content)
}
