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
	})
	if err != nil {
		fmt.Println("jwt err: ", err.Error())
		return
	}
	fmt.Println("token = \n", tk)

	jwt.Parse(tk)
	res, expired, err := jwt.Result()
	if err != nil {
		fmt.Println("parse err: ", err.Error())
		return
	}
	fmt.Println("expired: ", expired)
	fmt.Println("res: ", res)
}
