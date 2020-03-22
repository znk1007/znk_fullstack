package userredis

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path"
	"runtime"

	"github.com/rs/zerolog/log"
	userconf "github.com/znk_fullstack/server/usercenter/viewmodel/conf"
)

type redisConf struct {
	Env  userconf.Env `json:"env"`
	Host string       `json:"host"`
	Port string       `json:"port"`
}

var rc *redisConf

func init() {
	readConf()
}

func readConf() {
	rc = &redisConf{
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
}

//Start 启动Redis
func Start() error {
	if rc == nil {
		log.Info().Msg("read redis conf fail")
		return errors.New("read redis conf fail")
	}
	fmt.Println("rc = ", rc)
	switch userconf.GetEnv() {
	case userconf.Dev:
		fmt.Println()
	}
	return nil
}

//readFile 获取指定文件地址
func readFile(relativePath string) string {
	_, curPath, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(curPath) + "/" + relativePath)
}
