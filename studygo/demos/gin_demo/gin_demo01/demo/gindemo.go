package gindemo

import (
	"github.com/gin-gonic/gin"
)

type ginWebDemo struct {
	ginRouter *gin,
}

func init() {
	gd = ginWebDemo{
		ginRouter: gin.Default(),
	}
}

//StartEngine 启动服务
func (g GinDemo) StartEngine() {

}
