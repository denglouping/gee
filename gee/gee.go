package gee

import (
	"net/http"
)

type HandlerFunc func(ctx *Context)

type Engine struct {
	router *router
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

func (e *Engine) ServeHTTP (w http.ResponseWriter, req *http.Request) {
	ctx := newContext(w,req)
	e.router.handle(ctx)
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) error{
	return http.ListenAndServe(addr, engine)
}

func New() *Engine{
	engine := new(Engine)
	engine.router = newRouter()
	return engine

}
