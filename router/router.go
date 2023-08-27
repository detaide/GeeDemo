package router

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type router struct {
	roots map[string]*Node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots: make(map[string]*Node),
		handlers : make(map[string]HandlerFunc),
	}
}

func patternFormat(pattern string) string {
	index1 := strings.Index(pattern , ":")
	index2 := strings.Index(pattern , "*")
	// fmt.Println(index1, index2)
	
	if index1 == -1 && index2 == -1 {
		return pattern
	}

	if index2 >= index1  {
		return pattern[:index2 -1]
	}
	return pattern[:index1 -1]
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = NewHeadNode()
	}
	//添加入roots匹配树
	r.roots[method].InsertPath(pattern)
	key := method + "-" + patternFormat(pattern)
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, path string) (string, map[string]string) {
	return  r.roots[method].MatchPattern(path)

}

func (r *router) GetTreePath() {
	for _, childRoot := range r.roots {
		childRoot.GetTreePath("")
	}
}

func (r *router) handler(c *Context) {
	
	relativePath, params := r.getRoute(c.Method, c.Path)

	fmt.Println(relativePath)

	//匹配路径添加路径函数
	if relativePath != "" {
		key := c.Method + "-" + relativePath
		c.Params = params
		c.handlers = append(c.handlers, r.handlers[key])
	}else{
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}

	c.Next()
}