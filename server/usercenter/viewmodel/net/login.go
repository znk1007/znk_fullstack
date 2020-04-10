package usernet

import (
	"context"

	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
)

//loginService 登录服务
type loginService struct {
	req     *userproto.LoginReq
	resChan chan *userproto.LoginRes
}

//Do 执行任务
func Do() {

}

//UserLogin 登录接口
func (ls loginService) UserLogin(ctx context.Context, req *userproto.LoginReq) (res *userproto.LoginRes, err error) {
	return nil, nil
}
