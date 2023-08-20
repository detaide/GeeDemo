package main

import (
	// "fmt"
	"fmt"
	"gee"
	"net/http"
)

// func main(){
// 	fmt.Println("Serve start")
// 	r := gee.New()
// 	r.GET("/", func(c *gee.Context){
// 		c.HTML(http.StatusOK, "<h1>Hello World<h1/>")
// 	})

// 	r.GET("/hello", func (c *gee.Context) {
// 		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
// 	})

// 	r.POST("/Login", func (c *gee.Context) {
// 		c.JSON(http.StatusOK, gee.H{
// 			"username" : c.PostForm("username"),
// 			"password" : c.PostForm("password"),
// 		})
// 	})

// 	r.Run(":8088")
// }

func main() {
	// head := gee.NewHeadNode()
	// head.InsertPath("/a/:name/:username/:password")
	// head.InsertPath("/a/b/:name/:username")

	// head.GetTreePath("")
	r := gee.New()
	r.GET("/" , func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello World</h1>")
	})

	r.GET("/a/b", func(ctx *gee.Context) {
		ctx.String(http.StatusOK, "u'r in %s", ctx.Path)
	})
	r.GET("/c/:name/:password", func(ctx *gee.Context) {
		output := ""
		fmt.Println(ctx.Params)
		for k, v := range ctx.Params {
			output += fmt.Sprintf("Key : %s, value : %s\n" , k, v)
		}
		ctx.StringMap(http.StatusOK, output)
	})
	
	r.GET("/d/*", func(ctx *gee.Context) {
		ctx.JSON(http.StatusOK, gee.H{ "filepath" : ctx.Params["*"]})
	})
	// r.GetTreePath() 
	r.Run(":8088")
	
}