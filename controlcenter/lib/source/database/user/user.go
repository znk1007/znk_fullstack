package ccmydb

import "github.com/jinzhu/gorm"

//User 用户信息模型
type User struct {
	gorm.Model
	UserID   string `gorm:"type:varchar(40);unique_index"`
	Username string `gorm:"type:varchar(100)"`
	Account  string `gorm:"type:varchar(50)"`
	Password string `gorm:"type:varchar(100)"`
}
