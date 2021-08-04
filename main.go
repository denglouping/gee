package main

import (
	"gee"
	"net/http"
)

func main () {

	test1 := gee.New()

	test1.GET("/test", func(ctx *gee.Context) {
		ctx.JSON(200,gee.H{
			"user": "123",
		})
	})

	test1.GET("/test2", func(ctx *gee.Context) {
		ctx.String(200,"%s", ctx.Query("test"))
	})

	test1.GET("/test3", func(ctx *gee.Context) {
		ctx.Data(200,[]byte("testtesttest"))
	})
	http.ListenAndServe(":8888",test1)

}