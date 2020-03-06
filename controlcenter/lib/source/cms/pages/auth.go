package pages

import (
	"github.com/gin-gonic/gin"

	"github.com/znk_fullstack/controlcenter/lib/source/tools"
)

//AuthPage 验证页面
func AuthPage() {
	tools.Get(tools.NetHandler{
		Path: "/cms",
		HandlerFunc: func(c *gin.Context) {

		},
	})
}

func loginPage() {

}

func registerPage() {

}
