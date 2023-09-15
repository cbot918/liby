package httpy

import "fmt"

type router struct {
	route map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		route: make(map[string]HandlerFunc),
	}
}

func (r *router) addRoute(method string, path string, handler HandlerFunc) {
	key := method + "-" + path

	r.route[key] = handler
}

func (r *router) handle(key string, c *Context) {
	fn, ok := r.route[key]
	if ok {

		err := fn(c)
		if err != nil {
			fmt.Println("error: ", err)
		}

	} else {
		fmt.Fprintf(c.Writer, "request path not found")
	}
}
