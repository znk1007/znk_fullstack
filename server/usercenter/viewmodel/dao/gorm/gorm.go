package usergorm

import (
	"path"
	"runtime"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type gormConf struct {
	Username string `json: "username"`
}

func init() {
	gorm.Open("mysql")
}

func readGormConfig() {

}

//readFile 获取指定文件地址
func readFile(relativePath string) string {
	_, curPath, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(curPath) + "/" + relativePath)
}
