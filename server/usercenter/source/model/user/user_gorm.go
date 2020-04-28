package usermodel

import (
	"errors"

	usergorm "github.com/znk_fullstack/server/usercenter/source/controller/dao/gorm"
	userproto "github.com/znk_fullstack/server/usercenter/source/model/protos/generated"
)

// =======================mariadb===================================//

//gormCreateUser 创建用户模型
func gormCreateUser(user *userproto.User, password string) (exists bool, err error) {
	userDB := &UserDB{
		ID:       user.UserID,
		Password: password,
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

//gormFindUser 查询用户
func gormFindUser(userID string) (uDB UserDB, err error) {
	var userDB UserDB
	err = usergorm.DB().Model(
		&UserDB{
			ID: userID,
		},
	).First(&userDB).Error
	if err != nil {
		return
	}
	uDB = userDB
	return
}

//gormTotalUserCnt 总人数
func gormTotalUserCnt() (count int) {
	usergorm.DB().Model(&UserDB{}).Count(&count)
	return
}

//gormUserExists 用户是否存在
func gormUserExists(userID string) (exists bool) {
	var cnt int
	usergorm.DB().Model(
		&UserDB{
			ID: userID,
		},
	).Count(&cnt)
	exists = cnt > 0
	return
}

//gormUserOnline 用户是否在线
func gormUserOnline(userID string) (online int, err error) {
	var userDB UserDB
	userDB, err = gormFindUser(userID)
	if err != nil {
		return
	}
	online = userDB.Online
	return
}

//gormUpdateUserOnline 更新用户在线状态
func gormUpdateUserOnline(userID string, online int) (err error) {
	err = usergorm.DB().Model(
		&UserDB{
			ID:     userID,
			Active: 1,
		},
	).Update("online", online).Error
	return
}

//gormUserActive 用户是否激活中
func gormUserActive(userID string) (active int, err error) {
	var userDB UserDB
	userDB, err = gormFindUser(userID)
	if err != nil {
		return
	}
	active = userDB.Active
	return
}

//gormUpdateUserActive 更新激活状态
func gormUpdateUserActive(userID string, active int) (e error) {
	e = usergorm.DB().Model(&UserDB{ID: userID}).Update("active", active).Error
	return
}

//gormFindActiveUser 查询激活状态用户信息
func gormFindActiveUser(userID string) (user UserDB, err error) {
	var userDB UserDB
	err = usergorm.DB().Model(
		&UserDB{
			ID:     userID,
			Active: 1,
		},
	).First(&userDB).Error
	if err != nil {
		return
	}
	user = userDB
	return
}

//gormUserPassword 用户密码
func gormUserPassword(userID string) (psw string, err error) {
	var userDB UserDB
	userDB, err = gormFindUser(userID)
	if err != nil {
		return
	}
	psw = userDB.Password
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