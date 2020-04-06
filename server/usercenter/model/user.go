
import userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
package model
type UserDB struct {
	gorm.Model
	UserID string `gorm:"not null;unique"`
	User userproto.User
}
//CreateUserDB 创建用户模型
func CreateUserDB(user *userproto.User) {
	
}