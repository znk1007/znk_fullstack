package usermodel

import (
	"errors"

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
		Active:   1,
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
		userDB.User = user
		err = usergorm.DB().Create(userDB).Error
		exists = usergorm.DB().NewRecord(userDB)
	}
	return
}

//gormFindUser 查询用户信息
func gormFindUser(userID string, online int) (user *userproto.User, err error) {
	userDB := &UserDB{}
	err = usergorm.DB().Model(
		&UserDB{
			ID:     userID,
			Online: online,
			Active: 1,
		},
	).First(&userDB).Error
	if err != nil {
		return
	}
	user = userDB.User
	return
}

//gormTotalUserCnt 总人数
func gormTotalUserCnt() (count int) {
	usergorm.DB().Model(&UserDB{}).Count(&count)
	return
}

//gormUpdateUserActive 更新激活状态
func gormUpdateUserActive(userID string, active int32) error {
	return usergorm.DB().Model(&UserDB{ID: userID}).Update("user.active", active).Error
}

//gormUpdateUserOnline 更新用户在线状态
func gormUpdateUserOnline(userID string, online int32) (err error) {
	err = usergorm.DB().Model(
		&UserDB{
			ID:     userID,
			Active: 1,
		},
	).Update("online", online).Error
	return
}

//gormUpdateUserPhone 更新手机号
func gormUpdateUserPhone(userID string, phone string) (err error) {
	err = usergorm.DB().Model(
		&UserDB{
			ID:     userID,
			Active: 1,
		},
	).Update("user.phone", phone).Error
	return
}

//gormUpdateUserPassword 更新用户密码
func gormUpdateUserPassword(userID string, password string) (err error) {
	err = usergorm.DB().Model(
		&UserDB{
			ID:     userID,
			Active: 1,
		},
	).Update("password", password).Error
	return
}

//UpdateUserNickname 更新昵称
func gormUpdateUserNickname(userID string, nickname string) (err error) {
	err = usergorm.DB().Model(
		&UserDB{
			ID:     userID,
			Active: 1,
		},
	).Update("user.nickname", nickname).Error
	return
}

//UpdateUserPhoto 更新头像
func gormUpdateUserPhoto(userID string, photo string) (err error) {
	err = usergorm.DB().Model(
		&UserDB{
			ID:     userID,
			Active: 1,
		},
	).Update("user.photo", photo).Error
	return
}

//gormUserActive 用户是否激活
func gormUserActive(userID string) (userDB *UserDB, err error) {
	u := &UserDB{}
	err = usergorm.DB().Model(
		&UserDB{
			ID: userID,
		},
	).First(u).Error
	if err != nil {
		return
	}
	userDB = u
	return
}
