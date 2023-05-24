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
	userGroup.AddAny("/hello", func(ctx *router.Context) {
		fmt.Fprintln(ctx.W, "any user hello mszlu.com")
	})
	//userGroup.AddGet("/hello2", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintln(w, "get user hello2  mszlu.com")
	//})
	userGroup.AddPost("/hello2", func(ctx *router.Context) {
		fmt.Fprintln(ctx.W, "post user hello2  mszlu.com")
	})

	//启动
	engine.Run()
}
