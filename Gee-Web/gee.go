package gee

import (
	// "gee/middlewares"
	"log"
	"net/http"
	"strings"
)

// type HandlerFunc func(http.ResponseWriter, *http.Request)
type HandlerFunc func(*Context)

type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

/*
 * 路由分组控制
 * 由于tries的路由规则，只需要在router上保持一个该路由前缀的engine结点就可以了
 * parent支持分组嵌套
*/
type RouterGroup struct {
	prefix string
	parent *RouterGroup
	middlewares []HandlerFunc
	engine *Engine
}

/*
 * engine需要支持group
*/
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func Default() *Engine {
	engine := New()
	engine.Use(Recovery())
	return engine
}

/*
 * 由于需要在全局的engine添加group，因此该engine保持的是同一个
*/
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	if prefix == "" {
		prefix = "/"
	}
	newGroup := &RouterGroup{
		prefix: group.prefix +  prefix,
		parent: group,
		engine: engine,
	}

	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, subPattern string, handler HandlerFunc) {
	pattern := ""
	if group.prefix == "" || group.prefix == "/" {
		pattern = subPattern
	}else{
		pattern = group.prefix + subPattern
	}
	
	log.Printf("Group %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handle HandlerFunc) {
	group.addRoute("POST", pattern, handle)
}

func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

// GET defines the method to add GET request
func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRouter("GET", pattern, handler)
}

// POST defines the method to add POST request
func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRouter("POST", pattern, handler)
}

// add to Engine RouterSet
func (e *Engine) addRouter(method string, pattern string, handler HandlerFunc) {
	e.router.addRoute(method, pattern, handler)
}

// 输出router树的所有路径
func (e *Engine) GetTreePath() {
	// e.router.getRoute("GET", "/c/dalas/123")
	e.router.GetTreePath()
}

func (e *Engine) Run(port string) (err error) {
	return http.ListenAndServe(port, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)

	//触发中间件
	middlewares := []HandlerFunc{}
	for _, group := range e.groups{
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c.handlers = middlewares
	e.router.handler(c)
}

func LogFmt() {
	// middlewares.LogFmt()
}