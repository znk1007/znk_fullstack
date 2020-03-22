package auth

import (
	"github.com/znk_fullstack/controlcenter/source/auth/controller"
	"github.com/znk_fullstack/controlcenter/source/tools"
)

//Start 启动验证模块
func Start() {
	fp := tools.GetFilePathFromCurrent("view")
	tools.LoadHTMLS(fp + "/html/*")
	tools.LoadStatic("assets", fp+"/assets")
	vGroup := tools.Gt.V1
	vGroup.GET("/znksys",controller.Home)
	// vGroup.POST("/login")
}
