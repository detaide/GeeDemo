package gee

import (
	"net/http"
)

// type HandlerFunc func(http.ResponseWriter, *http.Request)
type HandlerFunc func(*Context)

type Engine struct{
	router *router
}


func New() *Engine {
	return &Engine{ router : newRouter()}
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
func (e *Engine) addRouter(method string, pattern string, handler HandlerFunc){
	e.router.addRoute(method, pattern, handler)
}

//输出router树的所有路径
func (e *Engine) GetTreePath() {
	// e.router.getRoute("GET", "/c/dalas/123")
	e.router.GetTreePath()
}

func (e *Engine) Run(port string) (err error) {
	return http.ListenAndServe(port, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)
	e.router.handler(c)
}