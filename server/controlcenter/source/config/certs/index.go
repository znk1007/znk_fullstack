package certs

import (
	"path"
	"runtime"
)

//GetCert 获取证书
func GetCert(name string) string {
	return getFilePathFromCurrent("/" + name)
}

//getFilePathFromCurrent 获取指定文件地址
func getFilePathFromCurrent(relativePath string) string {
	_, curPath, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(curPath) + "/" + relativePath)
}
