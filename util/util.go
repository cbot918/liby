package util

import (
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

func Type(obj interface{}){
	fmt.Printf("%T\n", obj)
}

func LenString(obj []string){
	fmt.Println(len(obj))
}

func LenInt(obj []int){
	fmt.Println(len(obj))
}