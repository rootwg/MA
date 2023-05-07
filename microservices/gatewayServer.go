package main

import (
	"fmt"
	"net/http"
	"studentDemo/sys/router"
)

func main() {
	engine := router.NewEngine("8080")
	//添加分组user
	userGroup := engine.Group("user")
	userGroup.AddAny("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "any user hello mszlu.com")
	})
	//userGroup.AddGet("/hello2", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintln(w, "get user hello2  mszlu.com")
	//})
	userGroup.AddPost("/hello2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "post user hello2  mszlu.com")
	})

	//启动
	engine.Run()
}
