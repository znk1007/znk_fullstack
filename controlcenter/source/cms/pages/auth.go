package pages

import "github.com/gin-gonic/gin"

type authInfo struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

//CMSAuth 校验用户数据
func CMSAuth(ctx *gin.Context) {

}

func loginPage() {

}

func registerPage() {

}
