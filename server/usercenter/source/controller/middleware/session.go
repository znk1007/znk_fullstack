package usermiddleware

import (
	"errors"
	"strconv"
	"time"

	userjwt "github.com/znk_fullstack/server/usercenter/source/controller/jwt"
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
	uj *userjwt.UserJWT
}

//NewSession 创建session实例
func NewSession(expired int64) UserSession {
	return UserSession{
		uj: userjwt.NewUserJWT(expired),
	}
}

//SessionID 生成sessionID
func (us UserSession) SessionID(userID string, deviceID string) (sess string, err error) {
	sess, err = us.uj.Token(map[string]interface{}{
		"userID":   userID,
		"deviceID": deviceID,
	})
	return
}

//Parse 解析sessionID
func (us *UserSession) Parse(sess, userID, deviceID string) (expired bool, err error) {
	us.uj.Parse(sess, false)
	res, e := us.uj.Result()
	if e != nil {
		err = e
		expired = true
		return
	}
	ts, ok := res["timestamp"].(string)
	if !ok || len(ts) == 0 {
		err = errors.New("miss param `timestamp` or timestamp is empty")
		expired = true
		return
	}
	oldTS, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		err = errors.New("internal server error")
		expired = true
		return
	}
	nTS := time.Now().Unix()
	diff := nTS - oldTS
	if diff < int64(us.uj.ExpiredInterval()) {
		expired = false
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
	dID, ok := res["deviceID"].(string)
	if len(dID) == 0 {
		err = errors.New("miss param `deviceID` or deviceID is empty")
		expired = true
		return
	}
	if dID != deviceID {
		err = errors.New("device has been logout")
		expired = true
		return
	}
	return
}
