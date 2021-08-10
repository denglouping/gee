package main

import (
	"gee"
	"log"
	"net/http"
)

func main() {

	test1 := gee.New()

	//test1.GET("/test", func(ctx *gee.Context) {
	//	ctx.JSON(200, gee.H{
	//		"user": "123",
	//	})
	//})
	//
	//test1.GET("/test2", func(ctx *gee.Context) {
	//	ctx.String(200, "%s", ctx.Query("test"))
	//})
	//
	//test1.GET("/test3", func(ctx *gee.Context) {
	//	ctx.Data(200, []byte("testtesttest"))
	//})
	//
	//test1.GET("/test4/*abc", func(ctx *gee.Context) {
	//	ctx.Data(200, []byte(ctx.Params["abc"]))
	//})

	test1.Get("/test/hello", func(ctx *gee.Context) {
		ctx.String(200, "HELLO WORLD")
	})

	testGroup1 := test1.Group("/test/hello")
	testGroup1.Use(gee.AccessLogger())

	testGroup := test1.Group("/test")
	testGroup.Use(gee.TimeLogger())

	err := http.ListenAndServe(":8888", test1)
	if err != nil {
		log.Printf("error happened: %s", err.Error())
	}

}
