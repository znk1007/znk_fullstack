package usernet

import (
	"context"

	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
	userGenID "github.com/znk_fullstack/server/usercenter/viewmodel/generateId"

	"google.golang.org/grpc"
)

var usrv *userSrv

func init() {
	usrv = &userSrv{
		rsrv: newRgstSrv(),
		lsrv: newLgnSrv(),
	}
}

type userSrv struct {
	rsrv *rgstSrv
	lsrv *lgnSrv
}

//makeID 生成唯一ID
func makeID() string {
	return userGenID.GenerateID()
}

//registerRegistServer 注册到注册请求服务
func registerRegistServer(srv *grpc.Server) {
	userproto.RegisterUserSrvServer(srv, usrv)
}

//注册
func (u *userSrv) Regist(ctx context.Context, req *userproto.UserRgstReq) (res *userproto.UserRgstRes, err error) {
	u.rsrv.write(req)
	return u.rsrv.read()
}

//登录
func (u *userSrv) Login(ctx context.Context, req *userproto.UserLgnReq) (res *userproto.UserLgnRes, err error) {
	u.lsrv.write(req)
	return u.lsrv.read()
}
