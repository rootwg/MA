package router

import "net/http"

//处理请求的类型接口
type Handlerfunc func(w http.ResponseWriter, r *http.Request)

//路由
type Router struct {
	//路由分组 切片
	gr []*GroupRouter
}

//构建路由分组
func (r *Router) Group(name string) *GroupRouter {
	g := &GroupRouter{
		groupName:  name,
		handlerMap: make(map[string]map[string]Handlerfunc),
	}
	r.gr = append(r.gr, g)
	return g
}
