package imnetclient

import (
	imconf "github.com/znk_fullstack/server/imcenter/source/controller/conf"
	userproto "github.com/znk_fullstack/server/imcenter/source/model/protos/generated"
	"google.golang.org/grpc"
)

const (
	appkey = "fullstack-manznk"
)

//RegistUserCenter 注册用户中心
func RegistUserCenter(acc, pwd string) (err error) {
	ucEnv := imconf.GetUcEnv()
	addr := ucEnv.Host + ":" + ucEnv.Port
	creds, e := imconf.TLSCredentials()
	if e != nil {
		err = e
		return
	}
	conn, e := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(creds),
		grpc.WithBlock(),
	)
	if e != nil {
		err = e
		return
	}
	return
}

func NewUcSrv(cc *grpc.ClientConn) {

	userproto.NewUserSrvClient(cc)
}
