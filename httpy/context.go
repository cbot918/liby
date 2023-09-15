package httpy

import (
	"encoding/json"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	// response info
	StatusCode int
}

func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

func (c *Context) ReqQuery(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) ReqBody(v any) error {
	if err := json.NewDecoder(c.Req.Body).Decode(v); err != nil {
		return err
	}
	return nil
}

func (c *Context) ResSetStatuscode(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) ResAddHeader(key string, val string) {
	c.Writer.Header().Add(key, val)
}

func (c *Context) ResString(code int, message string) {
	c.ResAddHeader("Content-Type", "text/plain")
	c.ResSetStatuscode(code)
	c.Writer.Write([]byte(message))
}

func (c *Context) ResJson(code int, obj interface{}) {
	c.ResAddHeader("Content-Type", "application/json")
	encoder := json.NewEncoder(c.Writer)
	err := encoder.Encode(obj)
	if err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) ResFlex(code int, typ string, message string) {
	c.ResAddHeader("Content-Type", typ)
	c.ResSetStatuscode(code)
	c.Writer.Write([]byte(message))
}
