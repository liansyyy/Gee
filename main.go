package main

import (
	"gee"
	"middlewares"
	"net/http"
)

func main() {

	engine := gee.New()

	engine.Use(middlewares.Logger(), middlewares.Recovery())

	engine.GET("/", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "Hello, Gee!")
	})

	engine.GET("/panic", func(ctx *gee.Context) {
		panic("panic on purpose!")
	})

	engine.GET("/nihao", func(ctx *gee.Context) {
		ctx.String(http.StatusOK, "nihao, %s\n", ctx.QueryString("name"))
	})

	engine.POST("/login", func(ctx *gee.Context) {
		ctx.JSON(http.StatusOK, gee.H{
			"username": ctx.PostForm("username"),
			"password": ctx.PostForm("password"),
		})
	})

	v1 := engine.Group("/v1")
	{
		v1.Use(middlewares.SayHi())
	}

	v1.GET("/apple/*name", func(ctx *gee.Context) {
		ctx.String(http.StatusOK, "v1,name is : %s\n", ctx.Param("name"))
	})

	engine.Run(":9999")
}
