package ccdb

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

//CCUserTBL 用户信息模型
type CCUserTBL struct {
	gorm.Model
	UserID   string `gorm:"type:varchar(40);unique_index"`
	Username string `gorm:"type:varchar(100)"`
	Account  string `gorm:"type:varchar(50)"`
	Password string `gorm:"type:varchar(100)"`
	Phone    string `gorm:"type:varchar(20)"`
	Photo    string `gorm:"type:varchar(200)"`
}

//CreateUserTBL 创建用户表
func CreateUserTBL() error {
	if dbConn == nil {
		return errors.New("Connect database first")
	}
	u := &CCUserTBL{}
	exists := dbConn.db.HasTable(u)
	fmt.Println("table exists")
	if !exists {
		dbConn.db.CreateTable(u)
	}
	return nil
}

//UpsertUser 插入用户数据
func UpsertUser(user *CCUserTBL) error {
	if dbConn == nil {
		return errors.New("Connect database first")
	}
	var u *CCUserTBL
	dbConn.db.Where("userID = ?", user.UserID).First(&u)
	if u != nil && (u.UserID == user.UserID) {
		dbConn.db.Model(&u).Updates(CCUserTBL{Phone: user.Phone, Username: user.Username})
	}
	dbConn.db.Create(&user)
	return nil
}

//UpdateUser 更新用户信息
func UpdateUser(userID string, values map[string]interface{}) error {
	if dbConn == nil {
		return errors.New("Connect database first")
	}
	var user *CCUserTBL
	dbConn.db.Model(&user).Where("userId = ?", userID).Updates(values)
	return nil
}

//FindUser 查询用户第一条数据
func FindUser(userID string) (*CCUserTBL, error) {
	if dbConn == nil {
		return nil, errors.New("Connect database first")
	}
	var u *CCUserTBL
	dbConn.db.Where("userID = ?", userID).First(&u)
	if u == nil {
		return nil, errors.New("user not exists")
	}
	return u, nil
}
