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
	expiredinterval int64
	isExp           bool
	parseSucc       bool
	err             error
	res             map[string]interface{}
}

//DefaultInterval 默认时间间隔
func DefaultInterval() int64 {
	return 1000 * 60 * 5
}

//CreateUserJWT 创建用户jwt验证 expired 纳秒级别
func CreateUserJWT(expiredinterval int64) UserJWT {
	return UserJWT{
		expiredinterval: expiredinterval,
	}
}

//Token token令牌 纳秒级别
func (userJWT UserJWT) Token(params map[string]interface{}) (token string, err error) {
	tmpExp := userJWT.expiredinterval
	if tmpExp == 0 {
		tmpExp = 1000 * 60 * 5
	}
	mclms := jwt.MapClaims{
		"timestamp": time.Now().Add(time.Duration(tmpExp)).UnixNano(),
	}
	for idx, val := range params {
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
	userJWT.parseSucc = false
	userJWT.err = err
	userJWT.res = nil
	userJWT.isExp = true
	if err != nil {
		log.Info().Msg(err.Error())
		userJWT.err = err
		return
	}
	clms, ok := tk.Claims.(jwt.MapClaims)
	if ok && tk.Valid {
		exp := clms["timestamp"].(string)
		oldTS, err := strconv.ParseInt(exp, 10, 64)
		if err != nil {
			log.Info().Msg(err.Error())
			userJWT.err = err
			return
		}
		nTS := time.Now().UTC().UnixNano()
		diss := nTS - oldTS
		if diss < userJWT.expiredinterval {
			userJWT.isExp = false
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
			userJWT.parseSucc = true
		} else {
			userJWT.err = errors.New("type error")
		}
	} else {
		userJWT.err = errors.New("parse error")
	}

}

//Result 结果
func (userJWT UserJWT) Result() (res map[string]interface{}, expired bool, err error) {
	expired = userJWT.isExp
	res = userJWT.res
	err = userJWT.err
	return
}
