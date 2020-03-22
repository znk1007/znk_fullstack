package redis

import (
	"path"
	"runtime"
)

type conf struct {
}

func init() {

}

func Start() {

}

//readFile 获取指定文件地址
func readFile(relativePath string) string {
	_, curPath, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(curPath) + "/" + relativePath)
}
