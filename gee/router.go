package gee

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router{
	return &router{
		handlers: make(map[string]HandlerFunc),
	}
}

func (rt *router) addRoute (method string, path string, handler HandlerFunc) {
	rt.handlers[method + path] = handler
}

func (rt *router) handle (ctx *Context){
	if handler, ok := rt.handlers[ctx.Method + ctx.Path]; ok {
		handler(ctx)
	}else {
		ctx.String(404, "error :%s ", "page not found")
	}
}