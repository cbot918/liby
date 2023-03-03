package main

import (
	"fmt"

	"github.com/cbot918/liby/jwty"
	"github.com/cbot918/liby/util"
)

func main(){
	j := jwty.New()

	id := 123
	email:= "yale@gmail.com"
	fmt.Printf("\nid: %d\nemail: %s\n\n",id,email)
	token, err:= j.FastJwt(id,email); util.Checke(err,"")
	fmt.Printf("token: %s\n\n", token, )
	
	// fmt.Println( "base64 encode: ", util.B64Encode(bearer), )
	// fmt.Println( "base64 decode: ", util.B64Decode(bearer), )
	
	decodedId, decodedEmail := j.GetIdAndEmail(j.DecodeJwt(token))
	fmt.Println("decodedId: ",decodedId)
	fmt.Printf("decodedEmail: %s\n\n", decodedEmail)

}