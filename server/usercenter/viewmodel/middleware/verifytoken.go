package usermiddleware

import (
	"errors"

	"github.com/rs/zerolog/log"
	userjwt "github.com/znk_fullstack/server/usercenter/viewmodel/jwt"

	usercrypto "github.com/znk_fullstack/server/usercenter/viewmodel/crypto"
)

//VerifyToken token校验
type VerifyToken struct {
	uJWT *userjwt.UserJWT
}

//NewVerifyToken 初始化校验对象
func NewVerifyToken(expiredinterval int) VerifyToken {
	return VerifyToken{
		uJWT: userjwt.CreateUserJWT(expiredinterval),
	}
}

//Generate 生成token
func (verify VerifyToken) Generate(params map[string]interface{}) (token string, err error) {
	token, err = verify.uJWT.Token(params)
	return
}

//Verify 校验token
func (verify VerifyToken) Verify(token string) (res map[string]interface{}, deviceID string, platform string, expired bool, err error) {
	expired = true
	if len(token) == 0 {
		log.Info().Msg("miss param `token` or `token` is empty")
		err = errors.New("miss param `token` or `token` is empty")
		return
	}
	verify.uJWT.Parse(token, true)
	tk, exp, e := verify.uJWT.Result()
	expired = exp
	if e != nil {
		log.Info().Msg(e.Error())
		err = e
		return
	}
	var ok bool
	//设备ID
	deviceID, ok = tk["deviceID"].(string)
	if !ok || len(deviceID) == 0 {
		log.Info().Msg("deviceID cannot be empty")
		err = errors.New("deviceID cannot be empty")
		return
	}
	//平台类型
	platform, ok = tk["platform"].(string)
	if !ok || len(platform) == 0 {
		log.Info().Msg("platform cannot be empty")
		err = errors.New("platform cannot be empty")
		return
	}
	//秘钥
	key, ok := tk["appkey"].(string)
	if !ok || len(key) == 0 {
		log.Info().Msg("miss param `appkey`")
		err = errors.New("miss param `appkey`")
		return
	}
	if key != usercrypto.GetSecurityKeyString() {
		log.Info().Msg("appkey is error")
		err = errors.New("appkey is error")
		return
	}
	res = tk
	return
}
