package main

import (
	"fmt"
	"net/http"
	"studentDemo/sys/router"
)

func main() {
	engine := router.NewEngine()
	engine.AddRouter("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello gateway")
	})
	engine.Run()
}
