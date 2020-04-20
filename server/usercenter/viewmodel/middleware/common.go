package usermiddleware

import (
	"errors"
	"net/http"

	usermodel "github.com/znk_fullstack/server/usercenter/model/user"
	usercrypto "github.com/znk_fullstack/server/usercenter/viewmodel/crypto"
	netstatus "github.com/znk_fullstack/server/usercenter/viewmodel/net/status"
)

//CommonVerify 通用校验
func CommonVerify(acc string, tk Token, sess bool) (code int, err error) {
	//用户是否被禁用
	code, err = userFrozen(acc, tk.UserID)
	if err != nil {
		return
	}

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
	expired, err = DefaultSess.Parse(sessionID, verify.UserID, sessionID)
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

//userFrozen 用户是否被禁用
func userFrozen(acc, userID string) (code int, err error) {
	//用户是否被禁用
	code = http.StatusOK
	active, e := usermodel.UserActive(acc, userID)
	if e != nil {
		err = e
		code = http.StatusInternalServerError
		return
	}
	if active == 0 {
		err = errors.New("user has been frozen")
		code = netstatus.UserActive
		return
	}
	return
}

//sessionValid 会话是否过期
func sessionValid(acc, userID string) {

}
