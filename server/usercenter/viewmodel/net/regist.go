package usernet

import (
	"context"
	"errors"
	"fmt"

	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
)

var (
	reqChan chan *userproto.RegistReq
	resChan chan *userproto.RegistRes
)

var rs *registService

func init() {
	reqChan = make(chan *userproto.RegistReq)
	resChan = make(chan *userproto.RegistRes)
	rs = &registService{}
	go rs.run()
}

//RegistService 注册
type registService struct {
}

func (rs *registService) run() {
	for {
		select {
		case req := <-reqChan:
			fmt.Println(req)
		}
	}
}

//UserReigst 注册
func (rs *registService) UserReigst(context.Context, *userproto.RegistReq) (*userproto.RegistRes, error) {
	return nil, errors.New("regist failed")
}
