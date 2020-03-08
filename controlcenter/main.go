package main

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/znk_fullstack/controlcenter/source/cms"
	_ "github.com/znk_fullstack/controlcenter/source/config"
	ccdb "github.com/znk_fullstack/controlcenter/source/dao"
	_ "github.com/znk_fullstack/controlcenter/source/tools"
)

func main() {
	//日志配置
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.With().Caller().Logger()

	err := ccdb.ConnectDB("mysql", "localhost:3306", "root", "znk1007!", "znk")
	if err != nil {
		log.Info().Msg(err.Error())
		panic(err)
	}
	err = ccdb.CreateUserTBL()
	if err != nil {
		fmt.Println("create user table err: ", err)
	}
	cms.Start()
	// tools.Listen()
}
