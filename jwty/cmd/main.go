package main

import (
	"fmt"

	"github.com/cbot918/liby/jwty"
	"github.com/cbot918/liby/util"
)

func main(){
	// fmt.Println(
		j := jwty.New()
		bearer, err:= j.FastJwt(123,"yale@gmail.com"); util.Checke(err,"")
		fmt.Println(
			bearer,
		)
		fmt.Println(
			util.B64Encode(bearer),
		)
		
}