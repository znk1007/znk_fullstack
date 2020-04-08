package userjwt

import (
	"crypto/rsa"
	"errors"
	"io/ioutil"
	"path"
	"reflect"
	"runtime"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rs/zerolog/log"
	usercrypto "github.com/znk_fullstack/server/usercenter/viewmodel/crypto"
)

//UserJWT 用户jwt验证对象
type UserJWT struct {
	expiredinterval time.Duration
	isExp           bool
	parseSucc       bool
	err             error
	res             map[string]interface{}
}

//DefaultInterval 默认时间间隔
func DefaultInterval() time.Duration {
	return time.Second * 2
}

//CreateUserJWT 创建用户jwt验证 expired 纳秒级别
func CreateUserJWT(expiredinterval time.Duration) *UserJWT {
	return &UserJWT{
		expiredinterval: expiredinterval,
	}
}

//Token token令牌
func (userJWT *UserJWT) Token(params map[string]interface{}) (token string, err error) {
	tmpExp := userJWT.expiredinterval
	if tmpExp == 0 {
		tmpExp = DefaultInterval()
	}
	ts := time.Now().Add(time.Duration(tmpExp)).Unix()
	tsstr := strconv.FormatInt(ts, 10)
	mclms := jwt.MapClaims{
		"timestamp": tsstr,
	}
	for idx, val := range params {
		mclms[idx] = val
	}
	token, err = createToken(mclms)
	return
}

//Parse 解析jwt
func (userJWT *UserJWT) Parse(token string) {

	key := loadRSAPublicKeyFromDisk("key/jwt.rsa.pub")
	userJWT.parseSucc = false
	userJWT.res = nil
	userJWT.isExp = true
	tk, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		t.Header["kid"] = usercrypto.GetSecurityKeyString()
		t.Header["alg"] = "RS512"
		return key, nil
	})

	if err != nil {
		log.Info().Msg(err.Error())
		userJWT.err = err
		return
	}
	clms, ok := tk.Claims.(jwt.MapClaims)
	if !ok || !tk.Valid {
		userJWT.err = errors.New("parse error")
		return
	}
	if ok && tk.Valid {
		exp, ok := clms["timestamp"].(string)
		if !ok || len(exp) == 0 {
			userJWT.err = errors.New("miss param `timestamp`")
			return
		}
		oldTS, err := strconv.ParseInt(exp, 10, 64)
		if err != nil {
			userJWT.err = err
			return
		}
		nTS := time.Now().Unix()
		diff := nTS - oldTS
		if diff < int64(userJWT.expiredinterval) {
			userJWT.isExp = false
		}
		v := reflect.ValueOf(clms)
		if v.Kind() != reflect.Map {
			userJWT.err = errors.New("type error")
			return
		}
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
		}
	}
}

//Result 结果
func (userJWT UserJWT) Result() (res map[string]interface{}, expired bool, err error) {
	expired = userJWT.isExp
	res = userJWT.res
	err = userJWT.err
	return
}

//loadRSAPrivateKeyFromDisk 加载私钥
func loadRSAPrivateKeyFromDisk(location string) *rsa.PrivateKey {
	fp := readFile(location)
	keyData, e := ioutil.ReadFile(fp)
	if e != nil {
		panic(e.Error())
	}
	key, e := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if e != nil {
		panic(e.Error())
	}
	return key
}

//loadRSAPublicKeyFromDisk 加载公钥
func loadRSAPublicKeyFromDisk(location string) *rsa.PublicKey {
	fp := readFile(location)
	keyData, e := ioutil.ReadFile(fp)
	if e != nil {
		panic(e.Error())
	}
	key, e := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if e != nil {
		panic(e.Error())
	}
	return key
}

//createToken 生成token
func createToken(c jwt.Claims) (tk string, err error) {
	privateKey := loadRSAPrivateKeyFromDisk("key/jwt.rsa")
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, c)
	tk, err = token.SignedString(privateKey)
	return
}

//readFile 获取指定文件地址
func readFile(relativePath string) string {
	_, curPath, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(curPath) + "/" + relativePath)
}
