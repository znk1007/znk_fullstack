package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/rs/zerolog/log"

	"github.com/znk_fullstack/controlcenter/source/tools"
)

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

//EnvConfig 环境配置
type EnvConfig struct {
	Env Env `json:"env"`
}

var ec *EnvConfig

func init() {

}

func readEnv() {
	if ec != nil {
		return
	}
	fp := tools.GetFilePathFromCurrent("json/env.json")
	bs, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Info().Msg(err.Error())
		return
	}
	ec = &EnvConfig{}
	err = json.Unmarshal(bs, ec)
	if err != nil {
		log.Info().Msg(err.Error())
	}
}

//CurEnv 当前环境
func CurEnv() Env {
	return ec.Env
}
