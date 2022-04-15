package middlewares

import (
	"fmt"
	"gee"
	"log"
	"time"
)

func Logger() gee.HandlerFunc {
	return func(ctx *gee.Context) {
		fmt.Println("---------------------------------------------------------------------------------------------")
		startTime := time.Now()
		ctx.Next()
		log.Printf("[%d] %s in %v\n", ctx.StatusCode, ctx.Request.RequestURI, time.Since(startTime))
	}
}

func SayHi() gee.HandlerFunc {
	return func(ctx *gee.Context) {
		log.Println("say Hi~~~")
		ctx.Next()
	}
}
