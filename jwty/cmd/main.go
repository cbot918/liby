package main

import (
	"fmt"

	"github.com/cbot918/liby/jwty"
	"github.com/cbot918/liby/util"
)

func main(){
		j := jwty.New()
		bearer, err:= j.FastJwt(123,"yale@gmail.com"); util.Checke(err,"")
		fmt.Println( bearer, )
		// fmt.Println( "base64 encode: ", util.B64Encode(bearer), )
		// fmt.Println( "base64 decode: ", util.B64Decode(bearer), )
		
		j.DecodeJwt(bearer)

}