package util

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
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

func Logb(message byte) {
	fmt.Println(string(message))
}

func Loggj(obj interface{}) {
	bytes, _ := json.MarshalIndent(obj, "\t", "\t")
	fmt.Println(string(bytes))
}

func Type(obj interface{}) {
	fmt.Printf("%T\n", obj)
}

func LenString(obj []string) {
	fmt.Println(len(obj))
}

func LenInt(obj []int) {
	fmt.Println(len(obj))
}

func B64Encode(input string) string {
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

func Request(url string) string {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("res body: %s", string(body))
	return string(body)
}

func GetUuid() string {
	return uuid.New().String()
}

func GetUuidFill(source string, targetLen int) (res string) {
	res = source
	for len(res) < targetLen {
		res += "0"
	}
	return
}

func GetParseTime(input string) string {
	// Parse the input timestamp string
	parsedTime, err := time.Parse("2006-01-02 15:04:05.99 -0700 MST", input)
	if err != nil {
		fmt.Println("Failed to parse timestamp:", err)
		return ""
	}
	// Format the timestamp to the desired format
	formattedTime := parsedTime.UTC().Format("2006-01-02 15:04:05.99Z")
	return formattedTime
}
