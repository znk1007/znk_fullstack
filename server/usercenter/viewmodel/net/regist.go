package usernet

import (
	"context"
	"errors"

	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
	userpayload "github.com/znk_fullstack/server/usercenter/viewmodel/payload"
	"google.golang.org/grpc"
)

var rs *registService

func init() {
	rs = &registService{
		reqChan: make(chan *userproto.RegistReq),
		resChan: make(chan *userproto.RegistRes),
	}
}

//RegistService 注册
type registService struct {
	reqChan chan *userproto.RegistReq
	resChan chan *userproto.RegistRes
}

type registJob struct {
	req *userproto.RegistReq
}

func (s *registService) Do() {

}

//RegisterRegistServer 注册到注册请求服务
func RegisterRegistServer(srv *grpc.Server) {
	userproto.RegisterRegistSrvServer(srv, rs)
}

//UserReigst 注册
func (s *registService) UserReigst(ctx context.Context, req *userproto.RegistReq) (*userproto.RegistRes, error) {
	rs.reqChan <- req
	userpayload.Pool.Write(rs)
	return nil, errors.New("regist failed")
}
