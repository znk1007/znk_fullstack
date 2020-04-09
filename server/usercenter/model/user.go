package model

import (
	"errors"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
	usergorm "github.com/znk_fullstack/server/usercenter/viewmodel/dao/gorm"
	userredis "github.com/znk_fullstack/server/usercenter/viewmodel/dao/redis"
)

//UserDB 用户数据库模型
type UserDB struct {
	Identifier string `gorm:"primary_key"`
	Password   string
	User       *userproto.User
}

//CreateUser 创建用户模型
func CreateUser(user *userproto.User, password string) (exists bool, err error) {
	userDB := &UserDB{
		Identifier: user.UserID,
		Password:   password,
		User:       user,
	}
	//用户是否已存在
	exists = usergorm.DB().NewRecord(userDB)
	if exists {
		exists = true
		err = errors.New("user has been registed")
		return
	}
	if len(password) == 0 {
		log.Info().Msg("password cannot be empty")
		exists = false
		err = errors.New("password cannot be empty")
		return
	}
	if !exists {
		user.CreatedAt = time.Now().String()
		user.UpdatedAt = time.Now().String()
		user.Online = 0
		user.Active = 1
		err = usergorm.DB().Create(userDB).Error
		exists = usergorm.DB().NewRecord(userDB)
	}
	return
}

//FindUser 查询用户
func FindUser(userID string) (user *userproto.User, err error) {
	userDB := &UserDB{}
	err = usergorm.DB().Model(
		&UserDB{
			Identifier: userID,
			User: &userproto.User{
				Active: 1,
			},
		},
	).First(&userDB).Error
	if err != nil || userDB == nil {
		return
	}
	user = userDB.User
	return
}

//TotalUserCount 总用户数
func TotalUserCount() int {
	var count int
	usergorm.DB().Model(&UserDB{}).Count(&count)
	return count
}

//UpdateUserActive 更新激活状态
func UpdateUserActive(userID string, active int32) error {
	return usergorm.DB().Model(&UserDB{Identifier: userID}).Update("user.active", active).Error
}

//UpdateUserOnline 更新用户在线状态
func UpdateUserOnline(userID string, online int32) error {
	return usergorm.DB().Model(
		&UserDB{
			Identifier: userID,
			User: &userproto.User{
				Active: 1,
			},
		},
	).Update("user.online", online).Error
}

//UpdateUserPhone 更新手机号
func UpdateUserPhone(userID string, phone string) error {
	return usergorm.DB().Model(
		&UserDB{
			Identifier: userID,
			User: &userproto.User{
				Active: 1,
			},
		},
	).Update("user.online", phone).Error
}

//AccRegisted 账号信息是否已注册
func AccRegisted(acc string) (exs bool, ts int64, registed int) {
	exs = userredis.Exists(acc)
	ts = -1
	registed = 0
	if exs {
		infos, err := userredis.HMGet(acc, "ts", "registed")
		if err != nil || (infos != nil && len(infos) < 2) {
			return
		}
		tsstr, ok := infos[0].(string)
		if ok {
			ts, err = strconv.ParseInt(tsstr, 10, 64)
			if err != nil {
				ts = -1
			}
		}
		rgd, ok := infos[1].(int)
		if ok {
			registed = rgd
		}
	}
	return
}

//SetAccRegisted 保存注册信息
func SetAccRegisted(acc string, ts string, registed int) (e error) {
	e = userredis.HSet(acc, "ts", ts, "registed", registed)
	return
}
