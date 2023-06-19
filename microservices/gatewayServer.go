package main

import (
	"fmt"
	"html/template"
	"net/http"
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
			Name: "测试html",
		}
		ctx.HtmlTemplate("login.html", template.FuncMap{}, user, "D:\\spaceWork\\go\\MA\\microservices\\sys\\router\\tpl\\login.html", "D:\\spaceWork\\go\\MA\\microservices\\sys\\router\\tpl\\header.html")
	})
	userGroup.AddGet("/userInfo", func(ctx *router.Context) {
		_ = ctx.JSON(http.StatusOK, &User{
			Name: "测试json",
		})
	})
	//启动
	engine.Run()
}
