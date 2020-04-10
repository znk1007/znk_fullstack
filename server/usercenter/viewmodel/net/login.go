package usernet

import (
	"context"

	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
)

var ls *loginService

//loginService 登录服务
type loginService struct {
	req     *userproto.LoginReq
	resChan chan *userproto.LoginRes
}

//Do 执行任务
func (l *loginService) Do() {

}

//UserLogin 登录接口
func (l *loginService) UserLogin(ctx context.Context, req *userproto.LoginReq) (res *userproto.LoginRes, err error) {

	return nil, nil
}
