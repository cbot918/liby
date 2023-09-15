package httpy

import (
	"net/http"
)

type HandlerFunc func(c *Context) error

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

func (e *Engine) Get(path string, fn HandlerFunc) {
	e.router.addRoute("GET", path, fn)
}

func (e *Engine) Post(path string, fn HandlerFunc) {
	e.router.addRoute("POST", path, fn)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path

	context := NewContext(w, r)

	e.router.handle(key, context)

}
