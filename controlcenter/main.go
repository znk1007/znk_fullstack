package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/znk_fullstack/controlcenter/source/cms"
	_ "github.com/znk_fullstack/controlcenter/source/config"
	ccdb "github.com/znk_fullstack/controlcenter/source/dao"
	"github.com/znk_fullstack/controlcenter/source/tools"
	_ "github.com/znk_fullstack/controlcenter/source/tools"
)

func init() {
	//日志配置
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.With().Caller().Logger()
}

func main() {

	ccdb.CreateUserTBL()
	cms.Start()
	tools.Listen()
}
