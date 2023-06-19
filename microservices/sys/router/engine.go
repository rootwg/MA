package router

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
)

const ANY = "ANY"

//启动引擎
type Engine struct {
	//这里也可以不使用指针，使用指针是担心后续传递engine.Router变成值传递而不是传递的指针
	*Router
	port string    //端口
	pool sync.Pool //池化
}

//构建应用
func NewEngine(port string) *Engine {
	engine := &Engine{
		Router: &Router{},
		port:   port,
	}

	engine.pool.New = func() any {
		return engine.allocateContext()
	}
	return engine
}

func (e *Engine) allocateContext() any {
	return &Context{engin: e}
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
	ctx := e.pool.Get().(*Context)
	ctx.W = w
	ctx.R = r
	e.handleHttpRequest(ctx)
	e.pool.Put(ctx)
}
func (e *Engine) handleHttpRequest(ctx *Context) {
	gr := e.Router.gr
	//方法类型
	method := ctx.R.Method
	//请求地址 //user/hello
	url := ctx.R.RequestURI
	for _, groupRouter := range gr {
		//user
		groupName := groupRouter.groupName
		//获取方法url
		mothedUrl := SubStringLast(url, "/"+groupName)
		//获取方法url是否存在
		node := groupRouter.treeNode.Get(mothedUrl)
		if node != nil {
			methodMap := groupRouter.handlerMap[mothedUrl]
			//user/hello
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
			ctx.W.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(ctx.W, ctx.R.RequestURI+" not found 405")
			return
		}
	}
	//找不到handler 404
	ctx.W.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(ctx.W, ctx.R.RequestURI+" not found 404")
}
func SubStringLast(str string, substr string) string {
	//先查找有没有  0
	index := strings.Index(str, substr)
	if index == -1 {
		return ""
	}
	//5
	len := len(substr)
	//5 -最后  /hello
	return str[index+len:]
}
