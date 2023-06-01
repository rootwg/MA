package main

import (
	"fmt"
	"studentDemo/sys/router"
)

func main() {
	//modoule
	engine := router.NewEngine("8080")
	//添加分组user
	userGroup := engine.Group("user")
	userGroup.AddGet("/hello", func(ctx *router.Context) {
		fmt.Fprintln(ctx.W, "get user hello mszlu.com")
	})
	userGroup.AddGet("/hello2", func(ctx *router.Context) {
		fmt.Fprintln(ctx.W, "get user hello2  mszlu.com")
	})

	orderGroup := engine.Group("order")
	orderGroup.AddGet("/hello", func(ctx *router.Context) {
		fmt.Fprintln(ctx.W, "get order hello  mszlu.com")
	})
	orderGroup.AddGet("/hello3", func(ctx *router.Context) {
		fmt.Fprintln(ctx.W, "get order hello3  mszlu.com")
	})
	orderGroup.AddGet("/hello2/order", func(ctx *router.Context) {
		fmt.Fprintln(ctx.W, "get hello/order hello2  mszlu.com")
	})

	//启动
	engine.Run()
}
