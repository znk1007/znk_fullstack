package usermodel

import (
	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
)

//UserDB 用户数据库模型
type UserDB struct {
	ID       string `gorm:"primary_key"`
	Password string
	Abnormal int //账号是否异常
	User     *userproto.User
}
