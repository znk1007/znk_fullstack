
import (
	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
	usergorm "github.com/znk_fullstack/server/usercenter/viewmodel/dao/gorm"
)
package model
type UserDB struct {
	gorm.Model
	Identifier string `gorm:"not null;unique"`
	User userproto.User
}
//CreateUserDB 创建用户模型
func CreateUserDB(user *UserDB) {
	usergorm.DB().NewRecord(user)
}