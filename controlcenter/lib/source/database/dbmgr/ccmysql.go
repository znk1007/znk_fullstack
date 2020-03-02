package ccmydb

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//CCMyDB MySql管理对象
type CCMyDB struct {
	db       *gorm.DB
	user     string
	password string
	dbName   string
}

//CreateMyDB 创建数据库管理对象
func CreateMyDB(user string, password string, dbName string) CCMyDB {
	return CCMyDB{
		user:     user,
		password: password,
		dbName:   dbName,
	}
}

//ConnectDB 连接MySql数据库
func (mydb CCMyDB) ConnectDB() error {
	authformat := mydb.user + ":" + mydb.password + "@/" + mydb.dbName
	db, err := gorm.Open("mysql", authformat+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return err
	}
	mydb.db = db
	return nil
}

//CloseDB 关闭数据库
func (mydb *CCMyDB) CloseDB() {
	if mydb.db == nil {
		return
	}
	mydb.db.Close()
}
