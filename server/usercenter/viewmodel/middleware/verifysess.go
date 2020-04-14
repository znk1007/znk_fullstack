package usermiddleware

import (
	userjwt "github.com/znk_fullstack/server/usercenter/viewmodel/jwt"
)

const (
	sessionExpired = 60 * 60 * 24 * 7
)

var uJWT *userjwt.UserJWT

func init() {
	uJWT = userjwt.CreateUserJWT(sessionExpired)
}

//SessionID 生成sessionID
func SessionID() (sess string, err error) {
	sess, err = uJWT.Token(nil)
	return
}

//ParseSessionID 解析sessionID
func ParseSessionID(sess string) (expired bool, err error) {
	uJWT.Parse(sess, false)
	_, expired, err = uJWT.Result()
	return
}
