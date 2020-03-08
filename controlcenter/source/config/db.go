package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/rs/zerolog/log"

	"github.com/znk_fullstack/controlcenter/source/tools"
)

type itemsConfig struct {
	Items []itemConfig `json:"items"`
	Env   Env          `json:"env"`
}

type itemConfig struct {
	Env Env        `json:"env"`
	DBs []DBConfig `json:"dbs"`
}

//DBConfig 数据库配置
type DBConfig struct {
	Name     DBName `json:"name"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Dialect  string `json:"dialect"`
}

func (dbcf DBConfig) String() string {
	if len(dbcf.Dialect) == 0 {
		return "name=" + string(dbcf.Name) + "|host=" + dbcf.Host + "|port=" + dbcf.Port
	}
	return "name=" + string(dbcf.Name) + "|host=" + dbcf.Host + "|port=" + dbcf.Port + "|dialect=" + dbcf.Dialect
}

//DBName 数据库类名
type DBName string

const (
	//Redis redis数据库
	Redis DBName = "redis"
	//Gorm gorm连接库
	Gorm DBName = "gorm"
)

var items *itemsConfig

func init() {
	readDBItems()
}

//readDBItems 读取数据库配置
func readDBItems() {
	if items != nil {
		return
	}
	fp := tools.GetFilePathFromCurrent("json/db.json")
	bufs, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Info().Msg(err.Error())
		return
	}
	items = &itemsConfig{}
	err = json.Unmarshal(bufs, items)
	if err != nil {
		log.Info().Msg(err.Error())
	}
}

//GetDBConfig 获取数据库配置信息
func GetDBConfig(name DBName) DBConfig {
	if items == nil {
		readDBItems()
	}
	var dbcf DBConfig
	for _, item := range items.Items {
		if item.Env == CurEnv() {
			for _, db := range item.DBs {
				if db.Name == name {
					dbcf = db
					break
				}
			}
		}
	}
	return dbcf
}
