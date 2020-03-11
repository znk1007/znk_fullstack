package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/znk_fullstack/controlcenter/source/cms/model"
)

//UserAuthState 用户验证状态
func Auth(c *gin.Context) {
	var u model.AuthInfo
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(u.Platform) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("platform cannot be empty")})
		return
	}
	if u.Platform == "web" {

	}
}
