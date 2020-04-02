package main

import (
	userconf "github.com/znk_fullstack/server/usercenter/viewmodel/conf"
	// usergorm "github.com/znk_fullstack/server/usercenter/viewmodel/dao/gorm"
	userredis "github.com/znk_fullstack/server/usercenter/viewmodel/dao/redis"
)

func main() {
	userredis.ConnectRedis(userconf.Dev)
	// usergorm.ConnectMariaDB(userconf.Dev)
}
