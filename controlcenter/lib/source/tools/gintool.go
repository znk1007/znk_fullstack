package tools

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ginTool struct {
	serve  *http.Server
	router *gin.Engine
}

//NetHandler 网络处理对象
type NetHandler struct {
	Method      string
	Path        string
	HandlerFunc gin.HandlerFunc
}

var gt ginTool

func init() {
	r := gin.Default()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	gt = ginTool{
		serve:  s,
		router: r,
	}
}

//Group 路由组
func Group(path string, handlers []NetHandler) {
	g := gt.router.Group(path)
	for _, handler := range handlers {
		g.Handle(handler.Method, handler.Path, handler.HandlerFunc)
	}
}

//Handler 请求处理
func Handler(handler NetHandler) {
	gt.router.Handle(handler.Method, handler.Path, handler.HandlerFunc)
}

//Get Get请求
func Get(handler NetHandler) {
	gt.router.GET(handler.Method, handler.HandlerFunc)
}

//Post post请求
func Post(handler NetHandler) {
	gt.router.POST(handler.Method, handler.HandlerFunc)
}

//Listen 监听服务
func Listen() {
	gt.serve.ListenAndServe()
}
