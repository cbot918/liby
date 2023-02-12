package main

import (
	"fmt"

	"github.com/cbot918/liby/cmdy"
)

func main(){
	c := cmdy.New()
	fmt.Println(
		c.RunAndGet([]string{"ls"}),
	)
	
}