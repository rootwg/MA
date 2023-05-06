package router

import "net/http"

type Handlerfunc func(w http.ResponseWriter, r *http.Request)

type Router struct {
	handlerMap map[string]Handlerfunc
}

func (r *Router) Add(url string, handlerfunc Handlerfunc) {
	r.handlerMap[url] = handlerfunc
}
