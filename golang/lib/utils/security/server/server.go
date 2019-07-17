package security

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"

	"google.golang.org/grpc/credentials"
)

/**
 * 服务端证书配置
 */
type ServerConfig struct {
	CAFile   string
	CertFile string
	KeyFile  string
}

/**
 * 生成安全传输验证
 */
func (s *ServerConfig) GenerateServerCredentialsByCA() (credentials.TransportCredentials, error) {
	cert, err := tls.LoadX509KeyPair(s.CertFile, s.KeyFile)
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(s.CAFile)
	if err != nil {
		return nil, err
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		return nil, errors.New("certPool.AppendCertsFromPEM err")
	}
	cdt := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})
	return cdt, nil
}

/**
 * 获取安全传输验证
 */
func (s *ServerConfig) GenerateServerTLSCredentials() (credentials.TransportCredentials, error) {
	cdt, err := credentials.NewServerTLSFromFile(s.CertFile, s.KeyFile)
	if err != nil {
		return nil, err
	}
	return cdt, nil
}
