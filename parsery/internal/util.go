package parsery

import "fmt"

var (
	log  = fmt.Println
	logf = fmt.Printf
)

func logs(input []rune) {
	fmt.Println(string(input))
}

func logc(input rune) {
	fmt.Println(string(input))
}
