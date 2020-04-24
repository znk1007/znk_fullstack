package usernet

import (
	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
	usermiddleware "github.com/znk_fullstack/server/usercenter/viewmodel/middleware"
	userpayload "github.com/znk_fullstack/server/usercenter/viewmodel/payload"
)

type reqType int

const (
	//更新密码
	psw      reqType = iota
	phone    reqType = 1
	nickname reqType = 2
)

const (
	updateInfoExpired = 60 * 2
)

//updatePswRes 更新密码响应
type updatePswRes struct {
	res *userproto.UserUpdatePswRes
	err error
}

//updatePhoneRes 更新手机响应
type updatePhoneRes struct {
	res *userproto.UserUpdatePhoneRes
	err error
}

//updateNicknameRes 更新昵称响应
type updateNicknameRes struct {
	res *userproto.UserUpdateNicknameRes
	err error
}

//userUpdateInfoSrv 更新信息服务
type userUpdateInfoSrv struct {
	phoneReq    *userproto.UserUpdatePhoneReq
	phoneRes    chan updatePhoneRes
	nicknameReq *userproto.UserUpdateNicknameReq
	nicknameRes chan updateNicknameRes
	pswReq      *userproto.UserUpdatePswReq
	pswRes      chan updatePswRes
	doing       map[string]bool
	token       *usermiddleware.Token
	pool        userpayload.WorkerPool
	rt          reqType
}

func newUpdateInfoSrv() *userUpdateInfoSrv {
	srv := &userUpdateInfoSrv{
		phoneRes:    make(chan updatePhoneRes),
		nicknameRes: make(chan updateNicknameRes),
		pswRes:      make(chan updatePswRes),
		doing:       make(map[string]bool),
		token:       usermiddleware.NewToken(300),
	}
	srv.pool.Run()
	return srv
}
