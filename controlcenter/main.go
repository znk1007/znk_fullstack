package main

import (
	"fmt"

	ccdb "github.com/znk_fullstack/controlcenter/lib/source/database"
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
}
