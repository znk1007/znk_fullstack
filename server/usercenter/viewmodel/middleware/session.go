package usermiddleware

import (
	"errors"

	userjwt "github.com/znk_fullstack/server/usercenter/viewmodel/jwt"
)

const (
	//DefaultSessExp 默认超时
	defaultSessExp = 60 * 60 * 24 * 7
)

//DefaultSess 默认session
var DefaultSess UserSession

func init() {
	DefaultSess = NewSession(defaultSessExp)
}

//UserSession session对象
type UserSession struct {
	uJWT *userjwt.UserJWT
}

//NewSession 创建session实例
func NewSession(expired int) UserSession {
	return UserSession{
		uJWT: userjwt.NewUserJWT(expired),
	}
}

//SessionID 生成sessionID
func (us UserSession) SessionID(userID string) (sess string, err error) {
	sess, err = us.uJWT.Token(map[string]interface{}{
		"userID": userID,
	})
	return
}

//Parse 解析sessionID
func (us UserSession) Parse(sess string, userID string) (expired bool, err error) {
	us.uJWT.Parse(sess, false)
	res, exp, e := us.uJWT.Result()
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
		err = errors.New("user invalidate")
		expired = true
		return
	}
	expired = exp
	return
}
