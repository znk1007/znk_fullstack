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
	Active   int //是否激活状态
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
		psd,
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
	_, err = gormCreateUser(
		user,
		password,
	)
	if err == nil {
		redisSetUserActive(acc, 1)
	}
	return
}

//UserActive 用户是否激活状态
func UserActive(acc, userID string) (active int, err error) {
	active, err = redisUserActive(acc)

	return
}

//FindUser 查询用户信息
func FindUser(acc, userID string) (user *userproto.User, err error) {

	per, e := redisUserPermission(acc)
	if e != nil {
		per = userproto.Permission_user
	}
	phone, email, nickname, photo, createdAt, updatedAt, e := redisGetUser(acc)

	if e != nil {
		u, e := gormFindUser(userID)
		if e != nil {
			err = e
			return
		}
		phone = u.GetPhone()
		email = u.GetEmail()
		nickname = u.GetNickname()
		photo = u.GetPhoto()
		createdAt = u.GetCreatedAt()
		updatedAt = u.GetUpdatedAt()
	}

	if err != nil {
		return
	}

	user = &userproto.User{
		Phone:      phone,
		Email:      email,
		Nickname:   nickname,
		Photo:      photo,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
		UserID:     userID,
		Permission: per,
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
