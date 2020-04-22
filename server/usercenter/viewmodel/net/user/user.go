package usernet

import (
	"context"

	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
	userGenID "github.com/znk_fullstack/server/usercenter/viewmodel/generateId"

	"google.golang.org/grpc"
)

var usrv userSrv

func init() {
	usrv = userSrv{
		rsrv:  newRgstSrv(),
		lsrv:  newLgnSrv(),
		upSrv: newUpdatePswSrv(),
	}
}

type userSrv struct {
	rsrv  *rgstSrv
	lsrv  *lgnSrv
	upSrv *updatePswSrv
}

//makeID 生成唯一ID
func makeID() string {
	return userGenID.GenerateID()
}

//RegisterRegistServer 注册到注册请求服务
func RegisterRegistServer(srv *grpc.Server) {
	userproto.RegisterUserSrvServer(srv, usrv)
}

//注册
func (u userSrv) Regist(ctx context.Context, req *userproto.UserRgstReq) (res *userproto.UserRgstRes, err error) {
	u.rsrv.write(req)
	return u.rsrv.read(ctx)
}

//登录
func (u userSrv) Login(ctx context.Context, req *userproto.UserLgnReq) (res *userproto.UserLgnRes, err error) {
	u.lsrv.write(req)
	return u.lsrv.read(ctx)
}

func (u userSrv) UpdatePassword(ctx context.Context, req *userproto.UserUpdatePswReq) (res *userproto.UserUpdatePswRes, err error) {
	u.upSrv.write(req)
	return u.upSrv.read(ctx)
}

func (u userSrv) Logout(ctx context.Context, req *userproto.UserLgoReq) (res *userproto.UserLgoRes, err error) {
	return
}
