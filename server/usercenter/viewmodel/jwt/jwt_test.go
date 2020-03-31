package userjwt

import (
	"testing"
)

func TestCreateToken(t *testing.T) {
	auth := map[string]interface{}{
		"userId":   "xxxxx",
		"password": "123456",
		"account":  "xxx",
		"email":    "xxxx@xxx.com",
		"phone":    "123456",
	}
	uJWT := CreateUserJWT(0)
	tkStr, e := uJWT.Token(auth)
	if e != nil {
		t.Error("create token failed: ", e)
		return
	}
	t.Log("token string: ", tkStr)

	uJWT.Parse(tkStr)
	res, _, e := uJWT.Result()
	if e != nil {
		t.Error("parse token failed: ", e)
		return
	}
	t.Log(res)
}
