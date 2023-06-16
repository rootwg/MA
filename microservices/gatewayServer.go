package main

import (
	"fmt"
	"html/template"
	"studentDemo/sys/router"
)

type User struct {
	Name string
}

func main() {
	//modoule
	engine := router.NewEngine("8080")
	//添加分组user
	userGroup := engine.Group("user")
	userGroup.AddGet("/index", func(ctx *router.Context) {
		fmt.Println("index handler")
		user := &User{
			Name: "go微服务框架",
		}
		ctx.HtmlTemplate("login.html", template.FuncMap{}, user, "D:\\spaceWork\\go\\MA\\microservices\\sys\\router\\tpl\\login.html", "D:\\spaceWork\\go\\MA\\microservices\\sys\\router\\tpl\\header.html")
	})

	//userGroup.AddGet("/hello", func(ctx *router.Context) {
	//	fmt.Fprintln(ctx.W, "get user hello mszlu.com")
	//})
	//userGroup.AddGet("/hello2", func(ctx *router.Context) {
	//	fmt.Fprintln(ctx.W, "get user hello2  mszlu.com")
	//})
	//
	//orderGroup := engine.Group("order")
	//orderGroup.AddGet("/hello", func(ctx *router.Context) {
	//	fmt.Fprintln(ctx.W, "get order hello  mszlu.com")
	//})
	//orderGroup.AddGet("/hello3", func(ctx *router.Context) {
	//	fmt.Fprintln(ctx.W, "get order hello3  mszlu.com")
	//})
	//orderGroup.AddGet("/hello2/order", func(ctx *router.Context) {
	//	fmt.Fprintln(ctx.W, "get hello/order hello2  mszlu.com")
	//})

	//启动
	engine.Run()
}
