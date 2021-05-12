package gee

import (
	"net/http"
)

type HandlerFunc func(c *Context)

type Engine struct {
	// 用map 类型实现路由表
	// [method-url]HandlerFunc
	route *Router
}

func New() *Engine {
	return &Engine{route: newRouter()}
}

func (engine *Engine) addRouter(method string, pattern string, handle HandlerFunc) {
	engine.route.addRoute(method, pattern, handle)
}

// 调用Get或Post方法注册路由
func (engine *Engine) GET(pattern string, handle HandlerFunc) {
	engine.route.addRoute("GET", pattern, handle)
}

func (engine *Engine) POST(pattern string, handle HandlerFunc) {
	engine.route.addRoute("POST", pattern, handle)
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.route.handle(c)
}
