package main

import (
	"fmt"

	"github.com/cbot918/liby/loady"
	"github.com/cbot918/liby/util"
)

const (
	port = ":6379"
)

func main() {
	r := loady.New()
	r.Run("welcome to loady")

	r.Concurrent(testcase, 5)
	// r.Test(testcase)

}

func testcase() {
	res := util.Request("http://localhost" + port)
	fmt.Println(res)
}
