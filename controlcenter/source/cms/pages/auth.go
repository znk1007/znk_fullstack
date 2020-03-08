package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/znk_fullstack/controlcenter/source/tools"
)

//CMSAuth 校验
func CMSAuth() {
	tools.Get(tools.NetHandler{
		Path: "/cms",
		HandlerFunc: func(c *gin.Context) {
			c.String(http.StatusOK, "校验成功")
		},
	})
}

func loginPage() {

}

func registerPage() {

}
