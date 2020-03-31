package usermiddleware

import (
	"errors"

	"github.com/rs/zerolog/log"
	userjwt "github.com/znk_fullstack/server/usercenter/viewmodel/jwt"

	usercrypto "github.com/znk_fullstack/server/usercenter/viewmodel/crypto"
)

//CheckToken token校验
type CheckToken struct {
	UJWT userjwt.UserJWT
}

//Create 创建校验对象
func Create(expiredinterval int64) CheckToken {
	return CheckToken{
		UJWT: userjwt.CreateUserJWT(expiredinterval),
	}
}

//Do 校验token
func (check CheckToken) Do(md map[string]interface{}, expiredinterval int64) (expired bool, ts string, err error) {
	var token string
	if val, ok := md["token"]; ok {
		token = val.(string)
	}
	if len(token) == 0 {
		log.Info().Msg("miss param `sign` or `sign` is empty")
		expired = true
		err = errors.New("miss param `sign` or `sign` is empty")
		return
	}
	check.UJWT.Parse(token)
	tk, exp, e := check.UJWT.Result()
	expired = exp
	if e != nil {
		log.Info().Msg(err.Error())
		exp = true
		err = e
		return
	}
	key, ok := tk["appkey"]
	appkey := key.(string)
	if !ok {
		log.Info().Msg("miss param `appkey`")
		err = errors.New("miss param `appkey`")
		return
	}
	if len(appkey) == 0 {
		log.Info().Msg("appkey is empty")
		err = errors.New("appkey is empty")
		return
	}
	if appkey != usercrypto.GetSecurityKeyString() {
		log.Info().Msg("appkey is bad")
		err = errors.New("appkey is bad")
		return
	}
	var tsVal interface{}
	tsVal, ok = tk["timestamp"]
	if !ok {
		log.Info().Msg("miss param `timestamp`")
		err = errors.New("miss param `timestamp`")
		return
	}
	ts = tsVal.(string)
	return
}
