package router

import (
	"html/template"
	"log"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

//为html模板设置值 并返回
func (c *Context) HtmlTemplate(name string, funcMap template.FuncMap, data any, fileName ...string) {
	t := template.New(name)
	t.Funcs(funcMap)

	files, err := t.ParseFiles(fileName...)
	if err != nil {
		log.Panicln(err)
		return
	}
	c.W.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = files.Execute(c.W, data)
	if err != nil {
		log.Panicln(err)
		return
	}
}
