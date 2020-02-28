package gindemo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type gindemo struct {
	router *gin.Engine
}

var demo gindemo

func init() {
	demo = gindemo{
		router: gin.Default(),
	}
}

//Router 路由
func Router() *gin.Engine {
	return demo.router
}

//SlashGet 请求
func SlashGet() {
	demo.router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world")
	})
}

//Run 运行
func Run() {
	demo.router.Run()
}
