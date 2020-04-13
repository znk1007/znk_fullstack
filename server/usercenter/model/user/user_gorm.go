package usermodel

import (
	"errors"
	"time"

	"github.com/rs/zerolog/log"
	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
	usergorm "github.com/znk_fullstack/server/usercenter/viewmodel/dao/gorm"
)

// =======================mariadb===================================//

//gormCreateUser 创建用户模型
func gormCreateUser(user *userproto.User, password string) (exists bool, err error) {
	userDB := &UserDB{
		ID:       user.UserID,
		Password: password,
		Online:   0,
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
		user.Active = 1
		userDB.User = user
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
			ID: userID,
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
	return usergorm.DB().Model(&UserDB{ID: userID}).Update("user.active", active).Error
}

//UpdateUserOnline 更新用户在线状态
func UpdateUserOnline(userID string, online int32) error {
	return usergorm.DB().Model(
		&UserDB{
			ID: userID,
			User: &userproto.User{
				Active: 1,
			},
		},
	).Update("online", online).Error
}

//UpdateUserPhone 更新手机号
func UpdateUserPhone(userID string, phone string) error {
	return usergorm.DB().Model(
		&UserDB{
			ID: userID,
			User: &userproto.User{
				Active: 1,
			},
		},
	).Update("user.phone", phone).Error
}

//UpdateUserPassword 更新用户密码
func UpdateUserPassword(userID string, password string) (err error) {
	err = usergorm.DB().Model(
		&UserDB{
			ID: userID,
			User: &userproto.User{
				Active: 1,
			},
		},
	).Update("password", password).Error
	return
}

//UpdateUserNickname 更新昵称
func UpdateUserNickname(userID string, nickname string) (err error) {
	err = usergorm.DB().Model(
		&UserDB{
			ID: userID,
			User: &userproto.User{
				Active: 1,
			},
		},
	).Update("user.nickname", nickname).Error
	return
}

//UpdateUserPhoto 更新头像
func UpdateUserPhoto(userID string, photo string) (err error) {
	err = usergorm.DB().Model(
		&UserDB{
			ID: userID,
			User: &userproto.User{
				Active: 1,
			},
		},
	).Update("user.photo", photo).Error
	return
}
