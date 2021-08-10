package gee

import (
	"log"
	"time"
)

func TimeLogger() HandlerFunc {
	return func(ctx *Context) {
		t := time.Now()

		ctx.Next()

		log.Printf("[%d] %s in %v", ctx.StatusCode, ctx.Req.RequestURI, time.Since(t))
	}
}

func AccessLogger() HandlerFunc {
	return func(ctx *Context) {
		t := time.Now()

		ctx.Next()

		log.Printf("access [%d] %s in %v", ctx.StatusCode, ctx.Req.RequestURI, time.Since(t))
	}
}
