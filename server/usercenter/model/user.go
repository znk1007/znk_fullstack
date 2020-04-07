package model

import (
	"time"

	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
	usergorm "github.com/znk_fullstack/server/usercenter/viewmodel/dao/gorm"
)

//UserDB 用户数据库模型
type UserDB struct {
	Identifier string `gorm:"primary_key"`
	User       userproto.User
}

//CreateUser 创建用户模型
func CreateUser(user *userproto.User) (exists bool, msg string) {
	userDB := &UserDB{
		Identifier: user.UserID,
	}
	exists = usergorm.DB().NewRecord(userDB)
	if !exists {
		user.CreatedAt = time.Now().String()
		user.UpdatedAt = time.Now().String()
		usergorm.DB().Create(userDB)
		exists = usergorm.DB().NewRecord(userDB)
		msg = "create user success"
	} else {
		msg = "user exists"
	}
	return
}

//FindUser 查询用户
func FindUser(userID string) (user *userproto.User) {
	userDB := &UserDB{}
	usergorm.DB().Where(&UserDB{Identifier: userID}).First(&userDB)
	user = &userDB.User
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
