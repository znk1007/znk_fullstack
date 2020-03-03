package ccdb

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// import _ "github.com/jinzhu/gorm/dialects/mysql"
// import _ "github.com/jinzhu/gorm/dialects/postgres"
// import _ "github.com/jinzhu/gorm/dialects/sqlite"
// import _ "github.com/jinzhu/gorm/dialects/mssql"

//CCDB 数据库管理对象
type CCDB struct {
	db       *gorm.DB
	user     string
	password string
	dialect  string
	host     string
	dbName   string
}

//CreateCCDB 创建数据库管理对象
func CreateCCDB(dialect string, host string, user string, password string, dbName string) CCDB {
	return CCDB{
		user:     user,
		password: password,
		host:     host,
		dialect:  dialect,
		dbName:   dbName,
	}
}

//ConnectDB 连接MySql数据库
func (the CCDB) ConnectDB() error {
	if len(the.user) == 0 || len(the.password) == 0 || len(the.dialect) == 0 || len(the.host) == 0 || len(the.dbName) == 0 {
		return errors.New("user, password, dialect and dbName cannot be empty")
	}
	authformat := the.user + ":" + the.password + "@(" + the.host + ")/" + the.dbName
	db, err := gorm.Open("mysql", authformat+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return err
	}
	the.db = db
	return nil
}

//CloseDB 关闭数据库
func (the CCDB) CloseDB() {
	if the.db == nil {
		return
	}
	the.db.Close()
}
