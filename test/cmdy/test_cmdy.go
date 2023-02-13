package main

import (
	"fmt"

	"github.com/cbot918/liby/cmdy"
)

func main(){
	c := cmdy.New()

	// test func
	c.Run([]string{"ls"})
	
	// test func
	fmt.Println(
		c.RunAndGet([]string{"uname -a"}),
	)
	
}