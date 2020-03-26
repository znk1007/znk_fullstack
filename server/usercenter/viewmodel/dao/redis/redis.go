package userredis

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/rs/zerolog/log"
	userconf "github.com/znk_fullstack/server/usercenter/viewmodel/conf"
)

type redisConf struct {
	Env  userconf.Env `json:"env"`
	Host string       `json:"host"`
	Port string       `json:"port"`
}

var rds *redis.Client

func init() {
	readConf()
}

func readConf() {
	rc := &redisConf{
		Env:  userconf.Dev,
		Host: "localhost",
		Port: "6379",
	}
	fp := readFile("redis.json")
	bs, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Info().Msg(err.Error())
		return
	}
	err = json.Unmarshal(bs, rc)
	if err != nil {
		log.Info().Msg(err.Error())
		return
	}
	fmt.Println("config redis succ")
	rds = redis.NewClient(&redis.Options{
		Addr: rc.Host + ":" + rc.Port,
	})
}

//Set 保存数据
func Set(key string, val interface{}, exp time.Duration) error {
	if rds == nil {
		log.Info().Msg("redis client not init")
		return errors.New("redis client not init")
	}
	rds.Set(key, val, exp)
	return nil
}

//Get 取值
func Get(key string) (val string, err error) {
	if rds == nil {
		log.Info().Msg("redis client not init")
		val = ""
		err = errors.New("redis client not init")
		return
	}

	val, err = rds.Get(key).Result()
	return
}

//readFile 获取指定文件地址
func readFile(relativePath string) string {
	_, curPath, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(curPath) + "/" + relativePath)
}
