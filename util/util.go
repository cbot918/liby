package util

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
)

func Checke(err error, message string) {
	if err != nil {
		log.Fatal(err)
		fmt.Println(message)
	}
}

func Logg(message interface{}) {
	fmt.Println(message)
}

func Logb(message byte){
	fmt.Println(string(message))
}

func Loggj(obj interface{}){
	bytes, _ := json.MarshalIndent(obj, "\t", "\t")
	fmt.Println(string(bytes))
}

func Type(obj interface{}){
	fmt.Printf("%T\n", obj)
}

func LenString(obj []string){
	fmt.Println(len(obj))
}

func LenInt(obj []int){
	fmt.Println(len(obj))
}

func B64Encode(input string) string{
	return base64.StdEncoding.EncodeToString([]byte(input))
}

func B64Decode(input string) string {
	result, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		fmt.Println("liby util B64Decode error")
		log.Fatal(err)
	}
	return string(result)
}