package usermiddleware

import (
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"
	userjwt "github.com/znk_fullstack/server/usercenter/viewmodel/jwt"

	usercrypto "github.com/znk_fullstack/server/usercenter/viewmodel/crypto"
	netstatus "github.com/znk_fullstack/server/usercenter/viewmodel/net/status"
)

//Token token校验
type Token struct {
	uJWT       *userjwt.UserJWT
	Expired    bool
	DeviceID   string
	DeviceName string
	Platform   string
	Result     map[string]interface{}
	Password   string
	SessionID  string
	UserID     string
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

//VerifyByPswAndSess 校验token，包含password&sessionID
func (verify *Token) VerifyByPswAndSess(token string) (code int, err error) {
	code = http.StatusOK
	err = verify.VerifyByPsw(token)
	res := verify.Result
	//校验sessionID
	sessionID, ok := res["sessionID"].(string)
	if !ok || len(sessionID) == 0 {
		err = errors.New("sessionID cannot be empty")
		code = http.StatusBadRequest
		return
	}
	var expired bool
	expired, err = DefaultSess.Parse(sessionID, verify.UserID)
	if err != nil {
		code = http.StatusBadRequest
		return
	}
	if expired {
		code = netstatus.SessionInvalidate
		err = errors.New("session invalidate, please login again")
		return
	}
	verify.SessionID = sessionID
	return
}

//VerifyByPsw 校验token，包含password
func (verify *Token) VerifyByPsw(token string) (err error) {
	err = verify.Verify(token)
	if err != nil {
		return
	}
	res := verify.Result
	//校验userID
	userID, ok := res["userID"].(string)
	if !ok || len(userID) == 0 {
		err = errors.New("userID cannot be empty")
		return
	}
	psw, ok := res["password"].(string)
	if !ok || len(psw) == 0 {
		err = errors.New("password cannot be empty")
		return
	}
	var old string
	old, err = usercrypto.CBCDecrypt(psw)
	if old != psw {
		err = errors.New("password is error")
		return
	}
	verify.UserID = userID
	verify.Password = psw
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
