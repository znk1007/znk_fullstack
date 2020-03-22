package redis

import (
	"fmt"
	"path"
	"runtime"

	userconf "github.com/znk_fullstack/server/usercenter/viewmodel/conf"
)

type redisConf struct {
	Env  userconf.Env `json:"env"`
	Host string       `json:"host"`
	Port string       `json:"port"`
}

var rc redisConf

//Start 启动Redis
func Start() {
	switch userconf.GetEnv() {
	case userconf.Dev:
		fmt.Println()
	}
}

//readFile 获取指定文件地址
func readFile(relativePath string) string {
	_, curPath, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(curPath) + "/" + relativePath)
}
