package router

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"unicode"
)

type Context struct {
	W          http.ResponseWriter
	R          *http.Request
	engin      *Engine
	queryCache url.Values //参数
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

/*
	处理json数据
*/
func (c *Context) JSON(status int, data any) error {
	c.W.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.W.WriteHeader(status)
	rsp, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = c.W.Write(rsp)
	if err != nil {
		return err
	}
	return nil
}

/*
	处理xml数据
*/
func (c *Context) XML(status int, data any) error {
	c.W.Header().Set("Content-Type", "application/xml;charset=utf-8")
	c.W.WriteHeader(status)
	err := xml.NewEncoder(c.W).Encode(data)
	if err != nil {
		return err
	}
	return nil
}

/**
文件下载处理
*/
func (c *Context) File(filePath string) {
	http.ServeFile(c.W, c.R, filePath)
}

/**
指定下载后文件名
*/
func (c *Context) FileAttachment(filepath, filename string) {
	if isASCII(filename) {
		c.W.Header().Set("Content-Disposition", `attachment; filename="`+filename+`"`)
	} else {
		c.W.Header().Set("Content-Disposition", `attachment; filename*=UTF-8''`+url.QueryEscape(filename))
	}
	http.ServeFile(c.W, c.R, filepath)
}
func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

/**
访问文件系统
*/
func (c *Context) FileFromFS(filepath string, fs http.FileSystem) {
	defer func(old string) {
		c.R.URL.Path = old
	}(c.R.URL.Path)

	c.R.URL.Path = filepath

	http.FileServer(fs).ServeHTTP(c.W, c.R)
}

/**
重定向页面
*/
func (c *Context) Redirect(status int, location string) {
	if (status < http.StatusMultipleChoices || status > http.StatusPermanentRedirect) && status != http.StatusCreated {
		panic(fmt.Sprintf("Cannot redirect with status code %d", status))
	}
	http.Redirect(c.W, c.R, location, status)
}
func (c *Context) DefaultQuery(key, defaultValue string) string {
	array, ok := c.GetQueryArray(key)
	if !ok {
		return defaultValue
	}
	return array[0]
}

func (c *Context) GetQuery(key string) string {
	c.initQueryCache()
	return c.queryCache.Get(key)
}
func (c *Context) QueryArray(key string) (values []string) {
	c.initQueryCache()
	values, _ = c.queryCache[key]
	return
}
func (c *Context) GetQueryArray(key string) (values []string, ok bool) {
	c.initQueryCache()
	values, ok = c.queryCache[key]
	return
}

/**
获取参数map
*/
func (c *Context) initQueryCache() {
	if c.queryCache == nil {
		if c.R != nil {
			c.queryCache = c.R.URL.Query()
		} else {
			c.queryCache = url.Values{}
		}
	}
}
