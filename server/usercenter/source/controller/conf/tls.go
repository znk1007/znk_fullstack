package userconf

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"path"
	"runtime"

	"google.golang.org/grpc/credentials"
)

type tlsConf struct {
	ca         []byte
	srvPemfile string
	srvKeyfile string
}

var tc tlsConf

func init() {
	cafile := readFile("key/ca.pem")
	bs, err := ioutil.ReadFile(cafile)
	if err != nil {
		panic("must contain a ca file")
	}
	tc = tlsConf{
		ca:         bs,
		srvPemfile: readFile("key/server.pem"),
		srvKeyfile: readFile("key/server.key"),
	}
}

//readFile 获取指定文件地址
func readFile(relativePath string) string {
	_, curPath, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(curPath) + "/" + relativePath)
}

//CATLSCredentials ca证书tls安全认证
func CATLSCredentials() (credentials.TransportCredentials, error) {
	cert, err := tls.LoadX509KeyPair(tc.srvPemfile, tc.srvKeyfile)
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(tc.ca); !ok {
		return nil, errors.New("set x509 pem failed")
	}
	tcl := credentials.NewTLS(
		&tls.Config{
			Certificates: []tls.Certificate{
				cert,
			},
			ClientAuth: tls.RequireAndVerifyClientCert,
			ClientCAs:  certPool,
		},
	)
	return tcl, nil
}

//TLSCredentials tls安全验证
func TLSCredentials() (credentials.TransportCredentials, error) {
	tcl, err := credentials.NewServerTLSFromFile(tc.srvPemfile, tc.srvKeyfile)
	if err != nil {
		return nil, err
	}
	return tcl, nil
}
