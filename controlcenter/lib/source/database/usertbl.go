package ccdb

import (
	"errors"

	"github.com/jinzhu/gorm"
)

//CCUserTBL 用户信息模型
type CCUserTBL struct {
	gorm.Model
	UserID   string `gorm:"type:varchar(40);unique_index"`
	Username string `gorm:"type:varchar(100)"`
	Account  string `gorm:"type:varchar(50)"`
	Password string `gorm:"type:varchar(100)"`
}

//CreateUserTBL 创建用户表
func CreateUserTBL() error {
	if dbConn == nil {
		return errors.New("Connect database first")
	}
	u := &CCUserTBL{}
	dbConn.db.CreateTable(u)
	return nil
}

//InsertUser 插入用户数据
func InsertUser(user *CCUserTBL) error {
	if dbConn == nil {
		return errors.New("Connect database first")
	}
	dbConn.db.Create(&user)
	return nil
}

//UpdateUser 更新用户信息
func UpdateUser() error {

}
