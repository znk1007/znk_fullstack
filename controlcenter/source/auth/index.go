package auth

import (
	"fmt"

	"github.com/znk_fullstack/controlcenter/source/tools"
)

//Start 启动验证模块
func Start() {
	fp := tools.GetFilePathFromCurrent("view")
	fmt.Println("file path: ", fp)
	tools.Gt.Router.LoadHTMLGlob(fp + "/html/*")
	tools.Gt.Router.Static("assets", fp+"/assets")

}
