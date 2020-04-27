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
	usercrypto "github.com/znk_fullstack/server/usercenter/viewmodel/crypto"
)

//UserJWT 用户jwt验证对象
type UserJWT struct {
	expiredinterval int64
	parseSucc       bool
	err             error
	res             map[string]interface{}
}

//DefaultInterval 默认时间间隔
func DefaultInterval() int64 {
	return 60 * 2
}

//私钥
var privateKey *rsa.PrivateKey

//公钥
var publicKey *rsa.PublicKey

func init() {
	publicKey = loadRSAPublicKeyFromDisk("key/jwt.rsa.pub")
	privateKey = loadRSAPrivateKeyFromDisk("key/jwt.rsa")
}

//NewUserJWT 创建用户jwt验证 expired 纳秒级别
func NewUserJWT(expiredinterval int64) *UserJWT {
	return &UserJWT{
		expiredinterval: expiredinterval,
	}
}

func (uj *UserJWT) ExpiredInterval() int64 {
	return uj.expiredinterval
}

//Token token令牌
func (uj *UserJWT) Token(params map[string]interface{}) (token string, err error) {
	tmpExp := uj.expiredinterval
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
func (uj *UserJWT) Parse(token string, build bool) {
	uj.parseSucc = false
	uj.res = nil
	tk, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		t.Header["kid"] = usercrypto.GetSecurityKeyString()
		t.Header["alg"] = "RS512"
		return publicKey, nil
	})
	if err != nil {
		uj.err = err
		return
	}
	clms, ok := tk.Claims.(jwt.MapClaims)
	if !ok || !tk.Valid {
		uj.err = errors.New("internal error")
		return
	}
	if ok && tk.Valid && build {
		//是否生成map
		v := reflect.ValueOf(clms)
		if v.Kind() != reflect.Map {
			uj.err = errors.New("type error")
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
			uj.err = nil
			uj.res = kMap
			uj.parseSucc = true
		}
	}
}

//Result 结果
func (uj *UserJWT) Result() (res map[string]interface{}, err error) {
	res = uj.res
	err = uj.err
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
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, c)
	tk, err = token.SignedString(privateKey)
	return
}

//readFile 获取指定文件地址
func readFile(relativePath string) string {
	_, curPath, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(curPath) + "/" + relativePath)
}
