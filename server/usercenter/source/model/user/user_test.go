package usermodel

import (
	"testing"

	userjwt "github.com/znk_fullstack/server/usercenter/source/controller/jwt"
	userproto "github.com/znk_fullstack/server/usercenter/source/model/protos/generated"
)

func TestUserJWT(t *testing.T) {
	userDB := &UserDB{
		ID:       "test",
		Password: "123",
		User: &userproto.User{
			UserID:  "test",
			Account: "acc",
		},
	}
	userMap := map[string]interface{}{
		"user": userDB,
	}
	userJWT := userjwt.NewUserJWT(1)
	tk, err := userJWT.Token(userMap)
	if err != nil {
		t.Fatal("user jwt err")
	}
	t.Logf("user token = %v", tk)
}
