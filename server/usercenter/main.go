package main

import (
	// userconf "github.com/znk_fullstack/server/usercenter/viewmodel/conf"
	// usergorm "github.com/znk_fullstack/server/usercenter/viewmodel/dao/gorm"

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
