package tools

import (
	"path"
	"runtime"
)

//GetFilePathFromCurrent 获取指定文件地址
func GetFilePathFromCurrent(relativePath string) string {
	_, curPath, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(curPath) + "/" + relativePath)
}
