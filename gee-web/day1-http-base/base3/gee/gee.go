package gee

import (
	"fmt"
	"net/http"
)

type handlerFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	route map[string]handlerFunc
}

func New() *Engine {
	return &Engine{route: make(map[string]handlerFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, handle handlerFunc) {
	key := method + "-" + pattern
	engine.route[key] = handle
}

func (engine *Engine) Get(pattern string, handle handlerFunc) {
	engine.addRoute("GET", pattern, handle)
}

func (engine *Engine) Post(pattern string, handle handlerFunc) {
	engine.addRoute("POST", pattern, handle)
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path

	if handle, ok := engine.route[key]; ok {
		handle(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL.Path)
	}
}
