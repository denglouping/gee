package gee

import (
	"net/http"
	"strings"
)

type router struct {
	handlers map[string]HandlerFunc
	roots    map[string]*node
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}

	return parts
}

func (rt *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	rt.handlers[key] = handler

	if _, ok := rt.roots[method]; !ok {
		rt.roots[method] = new(node)
	}
	rt.roots[method].insert(pattern, parts, 0)

}

func (rt *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)

	if _, ok := rt.roots[method]; !ok {
		return nil, nil
	}
	n := rt.roots[method].search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' && len(part) > 1 {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
			}
		}
		return n, params
	}

	return nil, nil

}

func (rt *router) handle(ctx *Context) {
	n, params := rt.getRoute(ctx.Method, ctx.Path)
	if n != nil {
		ctx.Params = params
		key := ctx.Method + "-" + n.pattern
		rt.handlers[key](ctx)
	} else {
		ctx.String(http.StatusNotFound, "404 NOT FOUND: %s\n", ctx.Path)
	}

	ctx.Next()
}
