package model

import (
	"testing"

	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
	userjwt "github.com/znk_fullstack/server/usercenter/viewmodel/jwt"
)

func TestUserJWT(t *testing.T) {
	userDB := &UserDB{
		ID:       "test",
		Password: "123",
		Abnormal: 1,
		User: &userproto.User{
			UserID:  "test",
			Account: "acc",
		},
	}
	userMap := map[string]interface{}{
		"user": userDB,
	}
	userJWT := userjwt.CreateUserJWT(1)
	tk, err := userJWT.Token(userMap)
	if err != nil {
		t.Fatal("user jwt err")
	}
	t.Logf("user token = %v", tk)
}
