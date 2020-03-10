package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/znk_fullstack/controlcenter/source/cms/model"
)

//UserAuthState 用户验证状态
func UserAuthState(ctx *gin.Context) {
	var u model.AuthInfo
	if err := ctx.ShouldBind(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

}
