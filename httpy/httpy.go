package httpy

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request) error

type Engine struct {
	Route map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{
		Route: make(map[string]HandlerFunc),
	}
}

func (e *Engine) AddRoute(method string, path string, handler HandlerFunc) {
	route := method + "-" + path

	e.Route[route] = handler
}

func (e *Engine) Get(path string, fn HandlerFunc) {
	e.AddRoute("GET", path, fn)
}

func (e *Engine) Post(path string, fn HandlerFunc) {
	e.AddRoute("POST", path, fn)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path

	fn, ok := e.Route[key]
	if ok {

		err := fn(w, r)
		if err != nil {
			fmt.Println("error: ", err)
		}

	} else {
		fmt.Fprintf(w, "request path not found")
	}

}
