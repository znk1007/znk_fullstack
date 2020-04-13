package usermodel

import (
	"time"

	"github.com/rs/zerolog/log"
	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
	usercrypto "github.com/znk_fullstack/server/usercenter/viewmodel/crypto"
	usertools "github.com/znk_fullstack/server/usercenter/viewmodel/tools"
)

//UserDB 用户数据库模型
type UserDB struct {
	ID       string `gorm:"primary_key"`
	Password string
	Abnormal int //账号是否异常
	Online   int //是否已登录
	User     *userproto.User
}

//CreateUser 创建用户
func CreateUser(acc, photo, userID, password string) (err error) {
	phone := ""
	if usertools.VerifyPhone(acc) {
		phone = acc
	}
	email := ""
	if usertools.VerifyEmail(acc) {
		email = acc
	}
	psd, e := usercrypto.CBCEncrypt(password)
	if e != nil {
		log.Info().Msg(e.Error())
		err = e
		return
	}
	user := &userproto.User{
		UserID:   userID,
		Account:  acc,
		Phone:    phone,
		Email:    email,
		Nickname: acc,
		Photo:    photo,
	}
	err = redisCreateUser(
		acc,
		userID,
		psd,
		phone,
		email,
		user.Nickname,
		photo,
		time.Now().String(),
	)
	if err != nil {
		return
	}
	_, err = gormCreateUser(
		user,
		password,
	)
	return
}

//FindUser 查询用户信息
func FindUser(acc, userID string) (user *userproto.User, err error) {
	phone, email, nickname, photo, updatedAt, e := redisGetUser(acc)
	if e != nil {
		err = e
		user = nil
		return
	}

	user = &userproto.User{
		UserID: userID,
	}
	return
}

//UserRegisted 用户是否走注册流程
func UserRegisted(acc string) (exs bool, ts int64, registed int) {
	exs, ts, registed = redisUserRegisted(acc)
	return
}

//SetUserRegisted 保存注册流程信息
func SetUserRegisted(acc string, ts string, registed int) (e error) {
	e = redisSetUserRegisted(acc, ts, registed)
	return
}
