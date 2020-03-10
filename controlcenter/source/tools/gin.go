package tools

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ginTool struct {
	serve  *http.Server
	Router *gin.Engine
}

//NetHandler 网络处理对象
type NetHandler struct {
	Method      string
	Path        string
	HandlerFunc gin.HandlerFunc
}

//Gt gin 管理工具
var Gt ginTool

func init() {
	r := gin.Default()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	Gt = ginTool{
		serve:  s,
		Router: r,
	}
}

//Listen 监听服务
func Listen() {
	Gt.serve.ListenAndServe()
}
