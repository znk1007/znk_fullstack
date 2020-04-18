package usermiddleware

import (
	"errors"

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
func SessionID(userID string) (sess string, err error) {
	sess, err = uJWT.Token(map[string]interface{}{
		"userID": userID,
	})
	return
}

//ParseSessionID 解析sessionID
func ParseSessionID(sess string, userID string) (expired bool, err error) {
	uJWT.Parse(sess, false)
	res, exp, e := uJWT.Result()
	if e != nil {
		err = e
		expired = true
		return
	}
	orgUID, ok := res["userID"].(string)
	if !ok || len(orgUID) == 0 {
		err = errors.New("miss param `userID` or userID is empty")
		expired = true
		return
	}
	if orgUID != userID {
		err = errors.New("login user invalidate")
		expired = true
		return
	}
	expired = exp
	return
}
