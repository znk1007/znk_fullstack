package main

import (
	userconf "github.com/znk_fullstack/server/usercenter/viewmodel/conf"
	usergorm "github.com/znk_fullstack/server/usercenter/viewmodel/dao/gorm"
	userredis "github.com/znk_fullstack/server/usercenter/viewmodel/dao/redis"
	usernet "github.com/znk_fullstack/server/usercenter/viewmodel/net"
)

func main() {

	userconf.SetEnv(userconf.Dev)
	userredis.ConnectRedis()
	usergorm.ConnectMariaDB()
	usernet.RunRPC()
}

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
