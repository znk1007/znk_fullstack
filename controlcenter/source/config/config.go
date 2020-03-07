package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
)

type itemsConfig struct {
	Items []itemConfig `json:"items"`
}

type itemConfig struct {
	Env Env        `json:"env"`
	DBs []DBConfig `json:"dbs"`
}

//DBConfig 数据库配置
type DBConfig struct {
	Name DBName `json:"name"`
	Host string `json:"host"`
	Port string `json:"port"`
}

//Env 环境
type Env string

const (
	//Dev 开发环境
	Dev Env = "dev"
	//Test 测试环境
	Test Env = "test"
	//Prod 生产环境
	Prod Env = "prod"
)

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
	fmt.Println("items: \n", items)
}

func readDBItems() {
	if items != nil {
		return
	}
	fp := GetFilePathFromCurrent("../config/db.json")
	fmt.Println("fp: ", fp)
	bufs, err := ioutil.ReadFile(fp)
	if err != nil {
		fmt.Println("read file err: ", err.Error())
		return
	}
	items = &itemsConfig{}
	err = json.Unmarshal(bufs, items)
	if err != nil {
		fmt.Println("unmarshal err: ", err.Error())
	}
}

//GetFilePathFromCurrent 获取指定文件地址
func GetFilePathFromCurrent(relativePath string) string {
	_, curPath, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(curPath) + "/" + relativePath)
}

//GetDBConfig 获取数据库配置信息
func GetDBConfig(env Env, name DBName) DBConfig {
	if items == nil {
		readDBItems()
	}
	var dbcf DBConfig
	for _, item := range items.Items {
		if item.Env == env {
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
