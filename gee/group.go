package gee

import "log"

type RouterGroup struct {
	prefix      string
	parent      *RouterGroup
	engine      *Engine
	middlewares []HandlerFunc
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}

	// longer prefix group return faster
	if len(engine.groups) == 0 {
		engine.groups = append(engine.groups, newGroup)
	} else if len(engine.groups[len(engine.groups)-1].prefix) < len(newGroup.prefix) {
		engine.groups = append(engine.groups, newGroup)
	} else {
		for index := range engine.groups {
			if len(engine.groups[index].prefix) > len(newGroup.prefix) {
				group := engine.groups[index]
				engine.groups[index] = newGroup
				engine.groups = append(engine.groups, group)
				break
			}

		}
	}

	return newGroup
}

func (group *RouterGroup) Use(handler HandlerFunc) {
	group.middlewares = append(group.middlewares, handler)
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

func (group *RouterGroup) Get(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) Post(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}
