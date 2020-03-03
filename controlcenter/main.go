package main

import (
	"fmt"

	ccdb "github.com/znk_fullstack/controlcenter/lib/source/database/dbmgr"
)

func main() {
	dbmgr := ccdb.CreateCCDB("mysql", "localhost:3306", "root", "znk1007!", "znk")
	err := dbmgr.ConnectDB()
	if err != nil {
		fmt.Println("connect db err: ", err.Error())
		panic(err)
	}
}
