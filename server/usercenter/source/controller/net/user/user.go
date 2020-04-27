package usernet

import (
	"context"

	userGenID "github.com/znk_fullstack/server/usercenter/source/controller/generateId"
	userproto "github.com/znk_fullstack/server/usercenter/source/model/protos/generated"

	"google.golang.org/grpc"
)

type userSrv struct {
	rsrv      *rgstSrv
	lsrv      *lgnSrv
	uiSrv     *updateInfoSrv
	loSrv     *lgoSrv
	statusSrv *userStatusSrv
}

var usrv userSrv

func init() {
	usrv = userSrv{
		rsrv:      newRgstSrv(),
		lsrv:      newLgnSrv(),
		uiSrv:     newUpdateInfoSrv(),
		loSrv:     newLogSrv(),
		statusSrv: newStatusSrv(),
	}
}

//makeID 生成唯一ID
func makeID() string {
	return userGenID.GenerateID()
}

//RegisterRegistServer 注册到注册请求服务
func RegisterRegistServer(srv *grpc.Server) {
	// userproto.RegisterUserSrvServer(srv, usrv)
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
	u.uiSrv.writePswReq(req)
	return u.uiSrv.readPswRes(ctx)
}

func (u userSrv) Logout(ctx context.Context, req *userproto.UserLgoReq) (res *userproto.UserLgoRes, err error) {
	u.loSrv.write(req)
	return u.loSrv.read(ctx)
}

func (u userSrv) Status(ctx context.Context, req *userproto.UserStatusReq) (res *userproto.UserStatusRes, err error) {
	u.statusSrv.write(req)
	return u.statusSrv.read(ctx)
}
