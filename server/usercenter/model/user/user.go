package usermodel

import (
	"errors"
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
	Active   int //是否激活状态
	Online   int //用户是否已登录
	User     *userproto.User
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
	psw, e := usercrypto.CBCEncrypt(password)
	if e != nil {
		log.Info().Msg(e.Error())
		err = e
		return
	}
	createdAt := time.Now().String()
	user := &userproto.User{
		UserID:    userID,
		Account:   acc,
		Phone:     phone,
		Email:     email,
		Nickname:  acc,
		Photo:     photo,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}

	err = redisCreateUser(
		acc,
		userID,
		psw,
		phone,
		email,
		user.GetNickname(),
		photo,
		createdAt,
		createdAt,
	)
	if err != nil {
		return
	}
	var exists bool
	exists, err = gormCreateUser(
		user,
		password,
	)
	if !exists {
		err = errors.New("internal server error")
		return
	}
	if err == nil {
		redisSetUserActive(acc, 1)
	}
	return
}

//UserOnline 用户是否已登录
func UserOnline(acc, userID string) (online int, err error) {
	online, err = redisUserOnline(acc)
	if err != nil {
		online, err = gormUserOnline(userID)
	}
	return
}

//SetUserOnline 更新用户在线状态
func SetUserOnline(acc, userID string, online int) (err error) {
	err = redisSetUserOnline(acc, online)
	if err == nil {
		err = gormUpdateUserOnline(userID, online)
	}
	return
}

//UserActive 用户是否激活状态
func UserActive(acc, userID string) (active int, err error) {
	active, err = redisUserActive(acc)
	if err != nil {
		active, err = gormUserActive(userID)
	}
	return
}

//SetUserActive 更新用户激活状态
func SetUserActive(acc, userID string, active int) (err error) {
	err = redisSetUserActive(acc, active)
	if err != nil {
		err = gormUpdateUserActive(userID, active)
	}
	return
}

//UserPassword 账户密码
func UserPassword(acc, userID string) (psw string, err error) {
	psw, err = redisUserPassword(acc)
	if len(psw) == 0 {
		psw, err = gormUserPassword(userID)
	}
	return
}

//FindUser 查询用户信息
func FindUser(acc, userID string) (user *userproto.User, err error) {
	per, e := redisUserPermission(acc)
	if e != nil {
		per = userproto.Permission_user
	}
	//用户是否被禁用
	active, e := UserActive(acc, userID)
	if e != nil {
		user = nil
		err = e
		return
	}
	if active == 0 {
		user = nil
		err = errors.New("user has been frozen")
		return
	}

	//redis 中的数据
	user, e = redisGetUser(acc)
	if e != nil || user == nil {
		//mariadb中的数据
		user, e = gormFindActiveUser(userID)
		if e != nil {
			err = e
			return
		}
		user.Permission = per
	}
	user.Permission = per
	return
}
