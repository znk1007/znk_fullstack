package usernet

import (
	"context"
	"net/http"

	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
	userpayload "github.com/znk_fullstack/server/usercenter/viewmodel/payload"
)

var ls *loginService

type loginResponse struct {
	res *userproto.LoginRes
	err error
}

//loginService 登录服务
type loginService struct {
	req     *userproto.LoginReq
	resChan chan loginResponse
	doing   map[string]bool
}

//handleLogin 处理登录请求
func (l *loginService) handleLogin() {
	acc := l.req.GetAccount()
	if len(acc) == 0 {
		return
	}
}

/*
状态码：code，
反馈消息：message，
时间戳：timestamp，
用户信息：user
*/
func (l *loginService) makeLoginToken(code int, err error) {
	switch code {
	case http.StatusOK:

	}
}

//Do 执行任务
func (l *loginService) Do() {
	go l.handleLogin()
}

//UserLogin 登录接口
func (l *loginService) UserLogin(ctx context.Context, req *userproto.LoginReq) (*userproto.LoginRes, error) {
	userpayload.Pool.WriteHandler(func(j chan userpayload.Job) {
		l.req = req
		j <- l
	})
	for {
		select {
		case res := <-l.resChan:
			return res.res, res.err
		}
	}
}
