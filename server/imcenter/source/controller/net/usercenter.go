package netim

import (
	userproto "github.com/znk_fullstack/server/imcenter/source/model/protos/generated"
	"google.golang.org/grpc"
)

type ucSrv struct {
}

func NewUCSrv(cc *grpc.ClientConn) {
	userproto.NewUserSrvClient(cc)
}
