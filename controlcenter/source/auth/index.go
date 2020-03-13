package auth

import (
	"fmt"

	"github.com/znk_fullstack/controlcenter/source/tools"
)

//Start 启动验证模块
func Start() {
	fp := tools.GetFilePathFromCurrent("view")
	fmt.Println("file path: ", fp)
	tools.LoadHTMLS(fp + "/html/*")
	tools.LoadStatic("assets", fp+"/assets")
	vGroup := tools.Gt.V1
	vGroup.POST("/login")
}
