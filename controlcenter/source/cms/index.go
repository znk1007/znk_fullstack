package cms

import (
	"fmt"

	"github.com/znk_fullstack/controlcenter/source/cms/controller"
	"github.com/znk_fullstack/controlcenter/source/tools"
)

//Handler cms处理器
type Handler struct {
	ver     string
	verpath string
}

//Start 运行cms系统
func Start() {
	firstVersion()
}

func firstVersion() {
	fp := tools.GetFilePathFromCurrent("view/cms")
	fmt.Println("file path: ", fp)
	tools.Gt.Router.LoadHTMLGlob(fp + "/html/*")
	tools.Gt.Router.Static("/static", fp+"/")

	vGroup := tools.Gt.Router.Group("/v1")

	cmsGroup := vGroup.Group("/cms")
	cmsGroup.GET("", controller.Home)
	cmsGroup.POST("/regist", controller.Regist)
	cmsGroup.POST("/login", controller.Login)
}
