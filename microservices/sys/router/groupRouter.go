package router

import "net/http"

//路由分组
type GroupRouter struct {
	//分组名称
	groupName string
	//路由url，请求处理handler
	handlerMap map[string]map[string]Handlerfunc
}

//对分组进行添加路由
func (g *GroupRouter) add(url string, method string, handlerfunc Handlerfunc) {
	_, ok := g.handlerMap[url]
	if !ok {
		g.handlerMap[url] = make(map[string]Handlerfunc)
	}
	g.handlerMap[url][method] = handlerfunc
}
func (g *GroupRouter) AddAny(name string, handlerFunc Handlerfunc) {
	g.add(name, ANY, handlerFunc)
}
func (g *GroupRouter) AddGet(name string, handlerFunc Handlerfunc) {
	g.add(name, http.MethodGet, handlerFunc)
}
func (g *GroupRouter) AddPost(name string, handlerFunc Handlerfunc) {
	g.add(name, http.MethodPost, handlerFunc)
}
