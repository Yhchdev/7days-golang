package gee

import "net/http"

type H map[string]interface{}

type Context struct {
	//origin object
	Writer http.ResponseWriter
	Req    *http.Request
	// request Info
	Method string
	Path   string
	// Response Info
	StatusCode int
}
