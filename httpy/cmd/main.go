package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cbot918/liby/httpy"
)

var lg = fmt.Println
var lf = fmt.Printf

const port = ":8080"

type UserRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func main() {

	r := httpy.New()

	r = SetRouter(r)

	fmt.Println("listening ", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}

func SetRouter(r *httpy.Engine) *httpy.Engine {
	r.Get("/", func(c *httpy.Context) error {
		fmt.Println("request get in")

		return nil
	})

	r.Post("/", func(c *httpy.Context) error {

		user := &UserRequest{}
		if err := c.ReqBody(&user); err != nil {
			return err
		}
		lf("%#+v", user)

		c.ResFlex(200, "application/json", `{"hello":"world"}`)

		return nil
	})

	return r
}
