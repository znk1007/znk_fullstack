package usergorm

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/zerolog/log"
	userconf "github.com/znk_fullstack/server/usercenter/source/controller/conf"
)

var gdb *gorm.DB

//ConnectMariaDB 连接MariaDB
func ConnectMariaDB() {
	initMariaDB()
}

//DB 数据库句柄
func DB() *gorm.DB {
	return gdb
}

//Close 关闭数据库
func Close() {
	if gdb != nil {
		gdb.Close()
		gdb = nil
	}
}

func initMariaDB() {
	gc := userconf.GormSrvConf()
	var err error
	//"user:password@tcp(addr)/dbname?charset=utf8&parseTime=True&loc=Local"
	url := gc.Username + ":" + gc.Password + "@tcp(" + gc.Host + ":" + gc.Port + ")/" + gc.DB + "?charset=utf8&parseTime=True&loc=Local"
	gdb, err = gorm.Open(gc.Dialect, url)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
}

//checkDB 校验db句柄
func checkDB() error {
	if gdb == nil {
		return errors.New("db object can't be nil")
	}
	return nil
}

/*
type User struct {
  gorm.Model
  Name         string
  Age          sql.NullInt64
  Birthday     *time.Time
  Email        string  `gorm:"type:varchar(100);unique_index"`
  Role         string  `gorm:"size:255"` // 设置字段大小为255
  MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
  Num          int     `gorm:"AUTO_INCREMENT"` // 设置 num 为自增类型
  Address      string  `gorm:"index:addr"` // 给address字段创建名为addr的索引
  IgnoreMe     int     `gorm:"-"` // 忽略本字段
}
*/
