package router

import (
	"fmt"
	"log"
	"net/http"
)

const ANY = "ANY"

//启动引擎
type Engine struct {
	//这里也可以不使用指针，使用指针是担心后续传递engine.Router变成值传递而不是传递的指针
	*Router
	port string
}

//构建应用
func NewEngine(port string) *Engine {
	return &Engine{
		&Router{},
		port,
	}
}

//加载路由，启动http服务
func (e *Engine) Run() {
	http.HandleFunc("/", e.ServeHTTP)
	err := http.ListenAndServe(":"+e.port, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

//公用的http处理
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	gr := e.Router.gr
	//方法类型
	method := r.Method
	//请求地址
	url := r.RequestURI
	ctx := &Context{
		w, r,
	}
	for _, groupRouter := range gr {
		groupName := groupRouter.groupName
		for groupUrl, methodMap := range groupRouter.handlerMap {
			methodUrl := "/" + groupName + groupUrl
			if methodUrl == url {
				//优先处理ANY的请求
				handlerfunc, ok := methodMap[ANY]
				if ok {
					handlerfunc(ctx)
					return
				}
				//处理get\post\delet
				handlerfunc, ok = methodMap[method]
				if ok {
					handlerfunc(ctx)
					return
				}
				//如果根据方法类型获取不到则是非法的类型
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintln(w, r.RequestURI+" not found 405")
				return
			}
		}
	}
	//找不到handler 404
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, r.RequestURI+" not found 404")
}
