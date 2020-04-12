package main

import (
	"fmt"

	userconf "github.com/znk_fullstack/server/usercenter/viewmodel/conf"
	usergorm "github.com/znk_fullstack/server/usercenter/viewmodel/dao/gorm"
	userredis "github.com/znk_fullstack/server/usercenter/viewmodel/dao/redis"
	usermiddleware "github.com/znk_fullstack/server/usercenter/viewmodel/middleware"
	usernet "github.com/znk_fullstack/server/usercenter/viewmodel/net"
)

func main() {
	cnt := 900000
	fa := usermiddleware.NewFreqAccess(1, cnt/2)
	for idx := 0; idx < cnt; idx++ {
		freq := fa.AccessCtrl("test", func() bool {
			return true
		})
		if freq {
			fmt.Println("access too frequence")
		}
	}
	return
	userconf.SetEnv(userconf.Dev)
	userredis.ConnectRedis()
	usergorm.ConnectMariaDB()
	usernet.RunRPC()
}

/*
userDB := &model.UserDB{
	ID:       "test",
	Password: "123",
	Abnormal: 1,
	User: &userproto.User{
		UserID:  "test",
		Account: "acc",
	},
}
userMap := map[string]interface{}{
	"user": userDB,
}
userJWT := userjwt.CreateUserJWT(1)
tk, err := userJWT.Token(userMap)
if err != nil {
	fmt.Println("user jwt err: ", err)
}
fmt.Printf("user token = %v", tk)
*/

/*
	res := map[string]interface{}{
		"test":  1,
		"test2": "2",
	}
	t1, ok := res["test"].(string)
	if !ok || len(t1) ==0 {
		fmt.Println("test1 not ok")

	} else {
		fmt.Println("test1 ok: ", t1)
	}
	t2, ok := res["test2"].(string)
	if !ok {
		fmt.Println("test2 not ok")

	} else {
		fmt.Println("test2 ok: ", t2)
	}
*/
