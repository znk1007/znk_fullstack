package usermodel

import (
	"errors"
	"time"

	"github.com/rs/zerolog/log"
	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
	usercrypto "github.com/znk_fullstack/server/usercenter/viewmodel/crypto"
	usertools "github.com/znk_fullstack/server/usercenter/viewmodel/tools"
)

//Permission 权限
type Permission int

const (
	super   Permission = iota
	admin   Permission = 1
	user    Permission = 2
	visitor Permission = 3
)

//UserDB 用户数据库模型
type UserDB struct {
	ID       string `gorm:"primary_key"`
	Password string
	Active   int        //是否激活状态
	Online   int        //用户是否已登录
	Per      Permission //用户权限
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

//UserExists 用户是否存在
func UserExists(acc, userID string) (exists bool) {
	exists = redisUserExists(acc)
	if !exists {
		exists = gormUserExists(userID)
	}
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

//SetUserPassword 更新密码
func SetUserPassword(acc, userID, psw string) (err error) {
	err = redisSetUserPassword(acc, psw)
	if err == nil {
		err = gormUpdateUserPassword(userID, psw)
	}
	return
}

//SetUserPhone 更新手机号
func SetUserPhone(acc, userID, phone string) (err error) {
	err = redisSetUserPhone(acc, phone)
	if err == nil {
		err = gormUpdateUserPhone(userID, phone)
	}
	return
}

//SetUserNickname 更新昵称
func SetUserNickname(acc, userID, nickname string) (err error) {
	err = redisSetUserNickname(acc, nickname)
	if err == nil {
		err = gormUpdateUserNickname(userID, nickname)
	}
	return
}

//FindUser 查询用户信息
func FindUser(acc, userID string) (user *userproto.User, err error) {
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
	}
	return
}
