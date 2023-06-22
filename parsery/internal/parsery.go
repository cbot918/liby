package parsery

import (
	"fmt"
	"os"
)

type Parsery struct {
	File    string
	Content []rune
}

func New(file string) *Parsery {
	return &Parsery{
		File: file,
	}
}

func (p *Parsery) Run() {

	tokens := NewTokenizer(p.getContent()).GetTokens()
	// fmt.Println(tokens)

	_ = NewParser(tokens)
	// fmt.Println(ast)

}

func (p *Parsery) getContent() []rune {
	file, err := os.ReadFile(p.File)
	if err != nil {
		fmt.Println("read file failed: ", err)
		os.Exit(1)
	}
	return []rune(string(file))
}

func (p *Parsery) PrintContent() {
	fmt.Println(string(p.Content))
}
