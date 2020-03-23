package usertoken

import (
	"fmt"
	"reflect"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//CreateToken 生成token字符串
func CreateToken(expired time.Duration, params map[string]interface{}) (token string, err error) {
	// jwt.MapClaims
	// clms := jwt.StandardClaims{
	// 	Issuer:    "znk_1007",
	// 	ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
	// }
	tmpExp := expired
	if tmpExp == 0 {
		tmpExp = time.Hour * 24 * 7
	}
	mclms := jwt.MapClaims{
		"exp": time.Now().Add(tmpExp).Unix(),
	}
	pm := reflect.ValueOf(params)
	if pm.Kind() == reflect.Map {
		for _, k := range pm.MapKeys() {
			v := pm.MapIndex(k)
			if v.CanInterface() {
				mclms[k.String()] = v.Interface()
			}
		}
	}
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, mclms)

	token, err = tk.SignedString(GetSecurityKeyByte())
	if err != nil {
		return
	}
	err = nil
	return
}

//ParseToken 解析token
func ParseToken(token string) (res map[string]interface{}, err error) {
	tk, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected login method %v", t.Header["alg"])
		}
		return GetSecurityKeyByte(), nil
	})
	if err != nil {
		return
	}
	clms, ok := tk.Claims.(jwt.MapClaims)
	if ok && tk.Valid {
		err = nil
		v := reflect.ValueOf(clms)
		if v.Kind() == reflect.Map {
			kMap := make(map[string]interface{})
			for _, k := range v.MapKeys() {
				val := v.MapIndex(k)
				if val.CanInterface() {
					kMap[k.String()] = val.Interface()
				}
			}
			res = kMap
		}
	} else {

	}
	return
}

//ExpiredDuration 两分钟响应超时失效
func ExpiredDuration() time.Duration {
	return time.Duration(time.Minute * 2)
}
