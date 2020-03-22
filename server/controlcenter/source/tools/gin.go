package tools

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/znk_fullstack/controlcenter/source/config/certs"
)

type ginTool struct {
	serve  *http.Server
	router *gin.Engine
	V1     *gin.RouterGroup
}

//Gt gin 管理工具
var Gt ginTool

func init() {
	r := gin.Default()
	// tlsconf := &tls.Config{}
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		// TLSConfig: ,
	}

	Gt = ginTool{
		serve:  s,
		router: r,
		V1:     versionGenerator("v1", r),
	}
}

//versionGenerator 版本生成
func versionGenerator(ver string, r *gin.Engine) *gin.RouterGroup {
	return r.Group(ver)
}

//LoadHTMLS 加载HTML文件组
func LoadHTMLS(pattern string) {
	Gt.router.LoadHTMLGlob(pattern)
}

//LoadStatic 加载静态资源
func LoadStatic(relativePath string, root string) {
	Gt.router.Static(relativePath, root)
}

///Users/huangman/Desktop/golang/src/github.com/znk_fullstack/controlcenter/source/config/certs/server.crt
///Users/huangman/Desktop/golang/src/github.com/znk_fullstack/controlcenter/source/tools/config/certs/server.crt
//Listen 监听服务
func Listen() {
	crt := certs.GetCert("server.crt")
	k := certs.GetCert("server.key")
	fmt.Println("cert path: ", crt)
	fmt.Println("key path: ", k)
	if err := Gt.serve.ListenAndServeTLS(crt, k); err != nil {
		log.Info().Msg("listen server error: " + err.Error())
	}
}
