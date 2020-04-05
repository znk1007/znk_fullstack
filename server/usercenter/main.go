package main

import (
	// userconf "github.com/znk_fullstack/server/usercenter/viewmodel/conf"
	// usergorm "github.com/znk_fullstack/server/usercenter/viewmodel/dao/gorm"
	"fmt"
	"time"

	_ "github.com/znk_fullstack/server/usercenter/viewmodel/dao/redis"
	userjwt "github.com/znk_fullstack/server/usercenter/viewmodel/jwt"
)

func main() {
	// userredis.ConnectRedis(userconf.Dev)
	// usergorm.ConnectMariaDB(userconf.Dev)
	jwt := userjwt.CreateUserJWT(time.Duration(time.Second))
	tk, err := jwt.Token(map[string]interface{}{
		"key1": "test1",
		"key2": "test2",
		"key3": "test3",
	})
	if err != nil {
		fmt.Println("jwt err: ", err.Error())
		return
	}
	fmt.Println("token = ", tk)

	jwt.Parse(tk)
	res, expired, err := jwt.Result()
	if err != nil {
		fmt.Println("parse err 1: ", err.Error())
		return
	}
	fmt.Println("expired 1: ", expired)
	fmt.Println("res 1: ", res)

	testtk := "eyJhbGciOiJIUzI1NiJ9.eyJrZXkxIjoidGVzdDEiLCJrZXkyIjoidGVzdDIiLCJ0aW1lc3RhbXAiOiIxNTg2MDYxNjQwODk2In0.ickL9kjygLN1aZ5RS_TZ2jkGJiGyufBF9XSvZ0ah57A"
	jwt.Parse(testtk)
	res, expired, err = jwt.Result()
	if err != nil {
		fmt.Println("parse err 2: ", err.Error())
		return
	}
	fmt.Println("expired 2: ", expired)
	fmt.Println("res 2: ", res)
}
