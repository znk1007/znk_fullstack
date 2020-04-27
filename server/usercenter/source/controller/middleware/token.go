package usermiddleware

import (
	"errors"
	"net/http"
	"time"

	usercrypto "github.com/znk_fullstack/server/usercenter/source/controller/crypto"
	userjwt "github.com/znk_fullstack/server/usercenter/source/controller/jwt"
	netstatus "github.com/znk_fullstack/server/usercenter/source/controller/net/status"
	usertools "github.com/znk_fullstack/server/usercenter/source/controller/tools"
)

//Token token校验
type Token struct {
	uj         *userjwt.UserJWT
	freq       *usertools.Freq
	DeviceID   string
	DeviceName string
	Platform   string
	Result     map[string]interface{}
	SessionID  string
	UserID     string
	Password   string
}

//NewToken 初始化校验对象
func NewToken(expiredinterval int64, freqexp int64) *Token {
	return &Token{
		uj:   userjwt.NewUserJWT(expiredinterval),
		freq: usertools.NewFreq(freqexp),
	}
}

//Generate 生成token
func (verify *Token) Generate(params map[string]interface{}) (token string, err error) {
	token, err = verify.uj.Token(params)
	return
}

//Parse 解析校验token
func (verify *Token) Parse(acc, method, token string) (code int, err error) {
	code = http.StatusOK
	if len(acc) == 0 {
		err = errors.New("miss param `account` or `account` is empty")
		code = http.StatusBadRequest
		return
	}
	if len(token) == 0 {
		err = errors.New("miss param `data` or `data` is empty")
		code = http.StatusBadRequest
		return
	}
	if !verify.freq.Expired(acc, method, time.Now().Unix()) {
		err = errors.New("please request later")
		code = netstatus.RequestFrequence
		return
	}

	verify.uj.Parse(token, true)
	tkmap, e := verify.uj.Result()
	if e != nil {
		err = e
		code = http.StatusBadRequest
		return
	}
	//设备ID
	deviceID, ok := tkmap["deviceID"].(string)
	if !ok || len(deviceID) == 0 {
		err = errors.New("deviceID cannot be empty")
		code = http.StatusBadRequest
		return
	}
	//设备名
	deviceName, ok := tkmap["deviceName"].(string)
	if !ok || len(deviceName) == 0 {
		err = errors.New("deviceName cannot be empty")
		code = http.StatusBadRequest
		return
	}
	//平台类型
	platform, ok := tkmap["platform"].(string)
	if !ok || len(platform) == 0 {
		err = errors.New("platform cannot be empty")
		code = http.StatusBadRequest
		return
	}
	//秘钥
	key, ok := tkmap["appkey"].(string)
	if !ok || len(key) == 0 {
		err = errors.New("miss param `appkey`")
		code = http.StatusBadRequest
		return
	}
	if key != usercrypto.GetSecurityKeyString() {
		err = errors.New("appkey is error")
		code = http.StatusBadRequest
		return
	}
	verify.Result = tkmap
	verify.DeviceID = deviceID
	verify.DeviceName = deviceName
	verify.Platform = platform
	return
}
