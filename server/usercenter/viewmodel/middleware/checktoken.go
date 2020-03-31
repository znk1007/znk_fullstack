package usermiddleware

import (
	"errors"

	"github.com/rs/zerolog/log"
	userjwt "github.com/znk_fullstack/server/usercenter/viewmodel/jwt"

	usercrypto "github.com/znk_fullstack/server/usercenter/viewmodel/crypto"
)

//CheckToken token校验
type CheckToken struct {
	uJWT userjwt.UserJWT
}

//Create 创建校验对象
func Create(expiredinterval int64) CheckToken {
	return CheckToken{
		uJWT: userjwt.CreateUserJWT(expiredinterval),
	}
}

//Do 校验token
func (check CheckToken) Do(token string) (res map[string]interface{}, expired bool, err error) {
	if len(token) == 0 {
		log.Info().Msg("miss param `sign` or `sign` is empty")
		res = nil
		err = errors.New("miss param `sign` or `sign` is empty")
		return
	}
	check.uJWT.Parse(token)
	tk, exp, e := check.uJWT.Result()
	expired = exp
	if e != nil {
		log.Info().Msg(err.Error())
		res = nil
		err = e
		return
	}
	key, ok := tk["appkey"]
	appkey := key.(string)
	if !ok {
		log.Info().Msg("miss param `appkey`")
		err = errors.New("miss param `appkey`")
		res = nil
		return
	}
	if len(appkey) == 0 {
		log.Info().Msg("appkey is empty")
		err = errors.New("appkey is empty")
		res = nil
		return
	}
	if appkey != usercrypto.GetSecurityKeyString() {
		log.Info().Msg("appkey is bad")
		err = errors.New("appkey is bad")
		res = nil
		return
	}
	return
}
