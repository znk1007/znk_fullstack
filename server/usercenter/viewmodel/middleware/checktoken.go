package usermiddleware

import (
	"errors"
	"time"

	"github.com/rs/zerolog/log"
	userjwt "github.com/znk_fullstack/server/usercenter/viewmodel/jwt"

	usercrypto "github.com/znk_fullstack/server/usercenter/viewmodel/crypto"
)

//CheckToken token校验
type CheckToken struct {
	uJWT *userjwt.UserJWT
}

//Initialize 初始化校验对象
func Initialize(expiredinterval time.Duration) CheckToken {
	return CheckToken{
		uJWT: userjwt.CreateUserJWT(expiredinterval),
	}
}

//Generate 生成token
func (check CheckToken) Generate(params map[string]interface{}) (token string, err error) {
	token, err = check.uJWT.Token(params)
	return
}

//Verify 校验token
func (check CheckToken) Verify(token string) (res map[string]interface{}, expired bool, err error) {
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
	key, ok := tk["appkey"].(string)
	if !ok || len(key) == 0 {
		log.Info().Msg("miss param `appkey`")
		err = errors.New("miss param `appkey`")
		res = nil
		return
	}
	if key != usercrypto.GetSecurityKeyString() {
		log.Info().Msg("appkey is error")
		err = errors.New("appkey is error")
		res = nil
		return
	}
	return
}
