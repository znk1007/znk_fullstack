package usergorm

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"runtime"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/zerolog/log"
	userconf "github.com/znk_fullstack/server/usercenter/viewmodel/conf"
)

type gormConf struct {
	Dev  gormInfo `json:"dev"`
	Test gormInfo `json:"test"`
	Prod gormInfo `json:"prod"`
}

type gormInfo struct {
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Dialect  string `json:"dialect"`
	Name     string `json:"name"`
}

var gdb *gorm.DB

//ConnectMariaDB 连接MariaDB
func ConnectMariaDB(envir userconf.Env) {
	initMariaDB(envir)
}

//Close 关闭数据库
func Close() {
	if gdb != nil {
		gdb.Close()
		gdb = nil
	}
}

func readGormConf() *gormConf {
	gc := &gormConf{}
	file := readFile("gorm.json")
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		log.Info().Msg(err.Error())
		panic(err.Error())
	}
	err = json.Unmarshal(bs, gc)
	if err != nil {
		log.Info().Msg(err.Error())
		panic(err.Error())
	}
	return gc
}

func initMariaDB(envir userconf.Env) {
	gc := readGormConf()
	var err error
	//"user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	switch envir {
	case userconf.Dev:
		url := gc.Dev.Username + ":" + gc.Dev.Password + "@/" + gc.Dev.Name + "?charset=utf8&parseTime=True&loc=Local"
		gdb, err = gorm.Open(gc.Dev.Dialect, url)
	case userconf.Test:
		url := gc.Test.Username + ":" + gc.Test.Password + "@/" + gc.Test.Name + "?charset=utf8&parseTime=True&loc=Local"
		gdb, err = gorm.Open(gc.Test.Dialect, url)
	case userconf.Prod:
		url := gc.Prod.Username + ":" + gc.Prod.Password + "@/" + gc.Prod.Name + "?charset=utf8&parseTime=True&loc=Local"
		gdb, err = gorm.Open(gc.Prod.Dialect, url)
	}
	if err != nil {
		log.Info().Msg(err.Error())
		panic(err)
	}
}

//readFile 获取指定文件地址
func readFile(relativePath string) string {
	_, curPath, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(curPath) + "/" + relativePath)
}
