package ccdb

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
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
func CreateUserTBL() {
	if dbConn == nil {
		log.Panic().Err(errors.New("Connect database first"))
		return
	}
	u := &CCUserTBL{}
	exists := dbConn.db.HasTable(u)
	if !exists {
		dbConn.db.CreateTable(u)
	}
}

//UpsertUser 插入用户数据
func UpsertUser(user *CCUserTBL) {
	if dbConn == nil {
		log.Panic().Err(errors.New("Connect database first"))
		return
	}
	var u *CCUserTBL
	dbConn.db.Where("userID = ?", user.UserID).First(&u)
	if u != nil && (u.UserID == user.UserID) {
		dbConn.db.Model(&u).Updates(CCUserTBL{Phone: user.Phone, Username: user.Username})
	}
	dbConn.db.Create(&user)
}

//UpdateUser 更新用户信息
func UpdateUser(userID string, values map[string]interface{}) {
	if dbConn == nil {
		log.Panic().Err(errors.New("Connect database first"))
		return
	}
	var user *CCUserTBL
	dbConn.db.Model(&user).Where("userId = ?", userID).Updates(values)
}

//FindUser 查询用户第一条数据
func FindUser(userID string) *CCUserTBL {
	if dbConn == nil {
		log.Panic().Err(errors.New("Connect database first"))
		return nil
	}
	var u *CCUserTBL
	dbConn.db.Where("userID = ?", userID).First(&u)
	if u == nil {
		log.Info().Msg("user not exists")
		return nil
	}
	return u
}
