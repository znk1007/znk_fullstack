package userjwt

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rs/zerolog/log"
	usercrypto "github.com/znk_fullstack/server/usercenter/viewmodel/crypto"
)

//UserJWT 用户jwt验证对象
type UserJWT struct {
	expired  int64
	params   map[string]interface{}
	isExp    bool
	isParsed bool
	err      error
	res      map[string]interface{}
}

//CreateUserJWT 创建用户jwt验证 expired 纳秒级别
func CreateUserJWT(expired int64, params map[string]interface{}) UserJWT {
	return UserJWT{
		expired: expired,
		params:  params,
	}
}

//Token token令牌 纳秒级别
func (userJWT UserJWT) Token() (token string, err error) {
	tmpExp := userJWT.expired
	if tmpExp == 0 {
		tmpExp = 1000 * 60 * 5
	}
	mclms := jwt.MapClaims{
		"exp": time.Now().Add(time.Duration(tmpExp)).UnixNano(),
	}
	for idx, val := range userJWT.params {
		mclms[idx] = val
	}
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, mclms)
	token, err = tk.SignedString(usercrypto.GetSecurityKeyByte())
	return
}

//Parse 解析jwt
func (userJWT UserJWT) Parse(token string) {
	tk, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected login method %v", t.Header["alg"])
		}
		return usercrypto.GetSecurityKeyByte(), nil
	})
	userJWT.isParsed = true
	if err != nil {
		log.Info().Msg(err.Error())
		userJWT.err = err
		return
	}
	clms, ok := tk.Claims.(jwt.MapClaims)
	if ok && tk.Valid {
		exp := clms["exp"].(string)
		oldTS, err := strconv.ParseInt(exp, 10, 64)
		if err != nil {
			log.Info().Msg(err.Error())
			userJWT.err = err
			return
		}
		nTS := time.Now().UnixNano()
		diss := nTS - oldTS
		if diss > userJWT.expired {
			userJWT.isExp = true
		} else {
			userJWT.isExp = false
		}
		if userJWT.isExp {
			userJWT.err = errors.New("token is expired")
			userJWT.res = nil
			return
		}
		v := reflect.ValueOf(clms)
		if v.Kind() == reflect.Map {
			kMap := make(map[string]interface{})
			for _, k := range v.MapKeys() {
				val := v.MapIndex(k)
				if val.CanInterface() {
					kMap[k.String()] = val.Interface()
				}
			}
			userJWT.err = nil
			userJWT.res = kMap
		}
	} else {
		userJWT.res = nil
		userJWT.err = errors.New("parse error")
	}
}

//Expired 是否已过期
func (userJWT UserJWT) Expired() (exp bool, err error) {
	if !userJWT.isParsed {
		exp = false
		err = errors.New("call Parse method first")
		return
	}
	exp = userJWT.isExp
	err = userJWT.err
	return
}
