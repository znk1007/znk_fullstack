package usermiddleware

import (
	"errors"

	"github.com/rs/zerolog/log"
	userjwt "github.com/znk_fullstack/server/usercenter/viewmodel/jwt"

	usercrypto "github.com/znk_fullstack/server/usercenter/viewmodel/crypto"
)

//Token token校验
type Token struct {
	uJWT       *userjwt.UserJWT
	Expired    bool
	DeviceID   string
	DeviceName string
	Platform   string
	Result     map[string]interface{}
}

//NewToken 初始化校验对象
func NewToken(expiredinterval int) Token {
	return Token{
		uJWT: userjwt.NewUserJWT(expiredinterval),
	}
}

//Generate 生成token
func (verify Token) Generate(params map[string]interface{}) (token string, err error) {
	token, err = verify.uJWT.Token(params)
	return
}

//Verify 校验token
func (verify *Token) Verify(token string) (err error) {
	if len(token) == 0 {
		log.Info().Msg("miss param `token` or `token` is empty")
		err = errors.New("miss param `token` or `token` is empty")
		return
	}
	verify.uJWT.Parse(token, true)
	tk, exp, e := verify.uJWT.Result()
	if e != nil {
		log.Info().Msg(e.Error())
		err = e
		return
	}
	//设备ID
	deviceID, ok := tk["deviceID"].(string)
	if !ok || len(deviceID) == 0 {
		log.Info().Msg("deviceID cannot be empty")
		err = errors.New("deviceID cannot be empty")
		return
	}
	//设备名
	deviceName, ok := tk["deviceName"].(string)
	if !ok || len(deviceName) == 0 {
		log.Info().Msg("deviceName cannot be empty")
		err = errors.New("deviceName cannot be empty")
		return
	}
	//平台类型
	platform, ok := tk["platform"].(string)
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
	verify.Result = tk
	verify.DeviceID = deviceID
	verify.DeviceName = deviceName
	verify.Platform = platform
	verify.Expired = exp
	return
}
