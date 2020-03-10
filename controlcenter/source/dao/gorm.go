package ccdb

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rs/zerolog/log"

	"github.com/znk_fullstack/controlcenter/source/config"
)

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

func init() {
	dbConn = createCCDB()
}

//createCCDB 创建数据库管理对象
func createCCDB() *CCDB {
	if dbConn != nil {
		return dbConn
	}
	dbcf := config.GetDBConfig(config.Gorm)
	user := dbcf.Username
	password := dbcf.Password
	dialect := dbcf.Dialect
	host := dbcf.Host
	dbName := dbcf.Name
	if len(user) == 0 || len(password) == 0 || len(dialect) == 0 || len(host) == 0 || len(dbName) == 0 {
		log.Panic().Err(errors.New("user, password, dialect and dbName cannot be empty"))
	}
	authformat := user + ":" + password + "@(" + host + ")/" + string(dbName)
	db, err := gorm.Open(dialect, authformat+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil || db == nil {
		log.Panic().Err(err)
		dbConn = nil
		return nil
	}
	conn := &CCDB{
		db: db,
	}
	return conn
}

//CloseDB 关闭数据库
func CloseDB() {
	if dbConn.db == nil {
		return
	}
	dbConn.db.Close()
	dbConn = nil
}
