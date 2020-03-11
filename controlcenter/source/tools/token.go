package tools

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

//CreateToken 生成token字符串
func CreateToken() (token string, succ bool) {
	clms := jwt.StandardClaims{
		Issuer:    "znk_1007",
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
	}
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, clms)

	token, err := tk.SignedString(GetSecurityKeyByte())
	if err != nil {
		succ = false
		return
	}
	succ = true
	return
}

func ParseToken(token string) {

}
