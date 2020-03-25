package userjwt

import "testing"
package userjwt
func TestCreateToken(t *testing.T) {
	auth := map[string]interface{}{
		"userId":   "xxxxx",
		"password": "123456",
		"account":  "xxx",
		"email":    "xxxx@xxx.com",
		"phone":    "123456",
	}
	tkStr, e := CreateToken(0, auth)
	if e != nil {
		t.Error("create token failed: ", e)
		return
	}
	t.Log("token string: ", tkStr)

	res, e := ParseToken(tkStr)
	if e != nil {
		t.Error("parse token failed: ", e)
		return
	}
	t.Log(res)
}