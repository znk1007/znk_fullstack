package usernet

import (
	"context"
	"errors"

	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
)

//RegistService 注册
type RegistService struct {
}

//UserReigst 注册
func (rs *RegistService) UserReigst(context.Context, *userproto.RegistReq) (*userproto.RegistRes, error) {
	return nil, errors.New("regist failed")
}
