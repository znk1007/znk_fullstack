package main

import (
	"fmt"

	"github.com/znk_fullstack/controlcenter/source/cms"
	_ "github.com/znk_fullstack/controlcenter/source/config"
	ccdb "github.com/znk_fullstack/controlcenter/source/dao"
	"github.com/znk_fullstack/controlcenter/source/tools"
)

func main() {

	err := ccdb.ConnectDB("mysql", "localhost:3306", "root", "znk1007!", "znk")
	if err != nil {
		fmt.Println("connect db err: ", err.Error())
		panic(err)
	}
	err = ccdb.CreateUserTBL()
	if err != nil {
		fmt.Println("create user table err: ", err)
	}
	cms.Start()
	tools.Listen()
}
