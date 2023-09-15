package main

import (
	"log"
	"net/http"

	"github.com/cbot918/liby/httpy"
)

func main() {
	r := httpy.New()

	r.GET("/", func(c *httpy.Context) error {
		return c.ResFlex(200, "text/html", `<h1>Helloooo</h1>`)
	})

	r.GET("/index", func(c *httpy.Context) error {
		return c.ResString(200, "/index")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *httpy.Context) error {
			return c.ResString(200, "/v1/")
		})

		v1.GET("/hello", func(c *httpy.Context) error {
			// expect /hello?name=geektutu
			res := "/v1/hello" + " " + c.ReqQuery("name") + " " + c.Path
			return c.ResString(200, res)
		})
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *httpy.Context) error {
			// expect /hello/geektutu
			res := "/v1/hello/" + c.ReqParam("name") + " " + c.Path
			return c.ResString(200, res)
		})
		v2.POST("/login", func(c *httpy.Context) error {
			return c.ResJson(http.StatusOK, map[string]interface{}{
				"username": "yale918",
				"password": "12345",
			})
		})
	}

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
