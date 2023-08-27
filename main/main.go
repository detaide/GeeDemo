package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.Default()
	
	r.GET("/", func(ctx *gee.Context) {
		ctx.String(http.StatusOK, "Hello World")
	})

	r.Run(":8080")

}