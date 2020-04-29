package imseckey

import (
	"path"
	"runtime"

	"google.golang.org/grpc/credentials"
)

var fm filem

type filem struct {
	srvpem string
	srvkey string
}

func init() {
	fm = filem{
		srvpem: filepath("key/server.pem"),
		srvkey: filepath("key/server.key"),
	}
}

//TLSCredentials ca安全验证配置
func TLSCredentials() (tc credentials.TransportCredentials, err error) {
	tc, err = credentials.NewClientTLSFromFile(fm.srvpem, fm.srvkey)
	return
}

//filepath 文件路径
func filepath(repath string) string {
	_, curpath, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(curpath) + "/" + repath)
}
