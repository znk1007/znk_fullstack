package ccdb

import "github.com/jinzhu/gorm"

//CCUser 用户信息模型
type CCUser struct {
	gorm.Model
	UserID   string `gorm:"type:varchar(40);unique_index"`
	Username string `gorm:"type:varchar(100)"`
	Account  string `gorm:"type:varchar(50)"`
	Password string `gorm:"type:varchar(100)"`
}

//User 全局用户对象实例
var User CCUser
