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

func Get() {

}

//Listen 监听服务
func Listen() {
	gt.serve.ListenAndServe()
}
