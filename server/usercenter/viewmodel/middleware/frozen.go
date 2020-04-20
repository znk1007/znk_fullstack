package usermiddleware

import (
	"errors"
	"net/http"

	usermodel "github.com/znk_fullstack/server/usercenter/model/user"
	netstatus "github.com/znk_fullstack/server/usercenter/viewmodel/net/status"
)

//UserFrozen 用户是否被禁用
func UserFrozen(acc string, userID string) (code int, err error) {
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
		code = netstatus.UserFrozen
		return
	}
	return
}
