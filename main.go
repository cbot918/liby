package main

import (
	"github.com/cbot918/liby/cmdy"
)

func main(){
	c := cmdy.New()
	
	str := []string{"touch 1.test","touch 2.test","touch 3.test"}
	c.Run(str)

}