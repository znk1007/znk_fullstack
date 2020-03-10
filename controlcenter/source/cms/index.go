package cms

import (
	"github.com/znk_fullstack/controlcenter/source/cms/middleware"
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
	cmsGroup := tools.Gt.Router.Group("/")
	cmsGroup.POST("cms", middleware.UserAuthState)
}
