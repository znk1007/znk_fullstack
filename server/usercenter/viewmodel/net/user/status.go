package usernet

import (
	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
	usermiddleware "github.com/znk_fullstack/server/usercenter/viewmodel/middleware"
	userpayload "github.com/znk_fullstack/server/usercenter/viewmodel/payload"
)

const (
	statusExpired = 60 * 5
)

type userStatusRes struct {
	res *userproto.UserStatusRes
	err error
}

//userStatusSrv 用户状态
type userStatusSrv struct {
	req   *userproto.UserStatusReq
	res   chan userStatusRes
	doing map[string]bool
	token *usermiddleware.Token
	pool  userpayload.WorkerPool
}

func newStatusSrv() *userStatusSrv {
	srv := &userStatusSrv{
		res:   make(chan userStatusRes),
		doing: make(map[string]bool),
		token: usermiddleware.NewToken(),
	}
	srv.pool.Run()
	return srv
}
