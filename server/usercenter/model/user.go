package model

import (
	"time"

	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
	usergorm "github.com/znk_fullstack/server/usercenter/viewmodel/dao/gorm"
)

//UserDB 用户数据库模型
type UserDB struct {
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
	Identifier string     `gorm:"primary_key"`
	User       userproto.User
}

//CreateUser 创建用户模型
func CreateUser(user *userproto.User) (exists bool, msg string) {
	userDB := &UserDB{
		Identifier: user.UserID,
	}
	exists = usergorm.DB().NewRecord(userDB)
	if !exists {
		userDB.CreatedAt = time.Now()
		userDB.UpdatedAt = time.Now()
		userDB.DeletedAt = nil
		user.CreatedAt = userDB.CreatedAt.String()
		user.UpdatedAt = userDB.UpdatedAt.String()
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

//UpdateUserFrozen 更新冻结状态
func UpdateUserFrozen(userID string, frozen int32) {

}
