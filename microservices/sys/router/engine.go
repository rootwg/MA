package router

import (
	"log"
	"net/http"
)

type Engine struct {
	router Router
}

func NewEngine() *Engine {
	return &Engine{
		router: Router{handlerMap: make(map[string]Handlerfunc)},
	}
}

/*
	加载路由，启动http服务
*/
func (e *Engine) Run() {
	for s, v := range e.router.handlerMap {
		http.HandleFunc(s, v)
	}
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func (e *Engine) AddRouter(url string, handlerfunc Handlerfunc) {
	e.router.handlerMap[url] = handlerfunc
}
