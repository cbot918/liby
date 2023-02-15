package main

import (
	"fmt"

	"github.com/cbot918/liby/ready"
)

func main(){
	path := "./data.txt"
	r := ready.New()
	fmt.Println(r.Get(path))
}