package usermiddleware

import (
	"errors"
	"net/http"

	usermodel "github.com/znk_fullstack/server/usercenter/model/user"
	usercrypto "github.com/znk_fullstack/server/usercenter/viewmodel/crypto"
	netstatus "github.com/znk_fullstack/server/usercenter/viewmodel/net/status"
)

//LoginVerify 登录验证
func LoginVerify(acc string, tk *Token) (code int, err error) {
	code = http.StatusOK
	exists := usermodel.UserExists(acc, tk.UserID)
	if !exists {
		err = errors.New("user not registed")
		code = netstatus.UserNotRegisted
		return
	}
	//用户是否被禁用
	code, err = userFrozen(acc, tk.UserID)
	if err != nil {
		return
	}
	res := tk.Result
	//校验userID
	userID, ok := res["userID"].(string)
	if !ok || len(userID) == 0 {
		code = http.StatusBadRequest
		err = errors.New("userID cannot be empty")
		return
	}
	psw, ok := res["password"].(string)
	if !ok || len(psw) == 0 {
		code = http.StatusBadRequest
		err = errors.New("password cannot be empty")
		return
	}
	//校验密码
	var old string
	old, err = usermodel.UserPassword(acc, userID)
	old, err = usercrypto.CBCDecrypt(old)
	if err != nil {
		code = http.StatusInternalServerError
		return
	}
	if len(old) == 0 {
		err = errors.New("internal server error")
		code = http.StatusInternalServerError
		return
	}
	if old != psw {
		code = http.StatusBadRequest
		err = errors.New("password no match")
		return
	}
	tk.UserID = userID
	tk.Password = psw
	return
}

//CommonRequestVerify 通用请求校验
func CommonRequestVerify(acc string, tk *Token) (code int, err error) {
	//用户是否被禁用
	code, err = LoginVerify(acc, tk)
	//校验用户是否退出登录
	var online int
	online, err = usermodel.UserOnline(acc, tk.UserID)
	if err != nil {
		code = http.StatusInternalServerError
		return
	}
	if online == 0 {
		err = errors.New("user has been logout")
		code = netstatus.UserLogout
		return
	}
	//校验sessionID
	res := tk.Result
	sessionID, ok := res["sessionID"].(string)
	if !ok || len(sessionID) == 0 {
		err = errors.New("sessionID cannot be empty")
		code = http.StatusBadRequest
		return
	}
	var expired bool
	expired, err = DefaultSess.Parse(sessionID, tk.UserID, tk.DeviceID)
	if err != nil {
		code = http.StatusBadRequest
		return
	}
	if expired {
		code = netstatus.SessionInvalidate
		err = errors.New("session invalidate, please login again")
		return
	}
	tk.SessionID = sessionID
	return
}

//userFrozen 用户是否被禁用
func userFrozen(acc, userID string) (code int, err error) {
	//用户是否被禁用
	code = http.StatusOK
	err = nil
	active, e := usermodel.UserActive(acc, userID)
	if e != nil {
		err = e
		code = http.StatusInternalServerError
		return
	}
	if active == 0 {
		err = errors.New("user has been frozen")
		code = netstatus.UserInactive
		return
	}
	return
}
