package cms

import "github.com/znk_fullstack/controlcenter/lib/source/cms/pages"

//Handler cms处理器
type Handler struct {
	ver     string
	verpath string
}

//Start 运行cms系统
func Start() {
	pages.AuthPage()
}
