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

var dbConn *CCDB

//createCCDB 创建数据库管理对象
func createCCDB(dialect string, host string, user string, password string, dbName string) *CCDB {
	return &CCDB{
		user:     user,
		password: password,
		host:     host,
		dialect:  dialect,
		dbName:   dbName,
	}
}

//ConnectDB 连接MySql数据库
func ConnectDB(dialect string, host string, user string, password string, dbName string) error {
	if len(user) == 0 || len(password) == 0 || len(dialect) == 0 || len(host) == 0 || len(dbName) == 0 {
		return errors.New("user, password, dialect and dbName cannot be empty")
	}
	dbConn = createCCDB(dialect, host, user, password, dbName)
	authformat := user + ":" + password + "@(" + host + ")/" + dbName
	db, err := gorm.Open("mysql", authformat+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return err
	}
	dbConn.db = db
	return nil
}

//CloseDB 关闭数据库
func CloseDB() {
	if dbConn.db == nil {
		return
	}
	dbConn.db.Close()
}
