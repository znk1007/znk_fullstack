package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/znk_fullstack/controlcenter/source/auth"
	_ "github.com/znk_fullstack/controlcenter/source/config"
	"github.com/znk_fullstack/controlcenter/source/tools"
	_ "github.com/znk_fullstack/controlcenter/source/tools"
)

func init() {
	//日志配置
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.With().Caller().Logger()
}

func main() {

	// ccdb.CreateUserTBL()
	auth.Start()
	tools.Listen()
}
