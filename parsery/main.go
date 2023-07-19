package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const (
	file      = "main.ro"
	colorRed  = "\033[0;31m"
	colorNone = "\033[0m"
)

var (
	lg = fmt.Println
	lf = fmt.Printf
	lj = func(v any) {
		res, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			fmt.Println("lv: 想indent的資料格式過不了")
		}
		fmt.Printf("%+v\n", string(res))
	}
	lt = func(t []Token) {
		for _, item := range t {
			fmt.Printf("%+v\n", item)
		}
	}
)

func main() {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(os.Stdout, "%s %s %s\n", colorRed, "source:", colorNone)
	lg(string(content))

	fmt.Fprintf(os.Stdout, "%s %s %s\n", colorRed, "tokens:", colorNone)
	tokenizer := NewTokenizer(content)
	result, err := tokenizer.GetTokens()
	if err != nil {
		log.Fatal(err)
	}
	lt(result)

}
