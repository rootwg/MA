package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
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
		dir, _ := os.Getwd()
		fmt.Printf("========")
		fmt.Printf(dir)
		ctx.HtmlTemplate("login.html", template.FuncMap{}, user, dir+"/tpl/login.html", dir+"/tpl/header.html")
	})
	userGroup.AddGet("/userInfo", func(ctx *router.Context) {
		_ = ctx.JSON(http.StatusOK, &User{
			Name: "测试json",
		})
	})
	//启动
	fmt.Printf("=====启动成功======")
	engine.Run()
}
