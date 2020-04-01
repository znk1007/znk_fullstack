package main

import (
	userconf "github.com/znk_fullstack/server/usercenter/viewmodel/conf"
	_ "github.com/znk_fullstack/server/usercenter/viewmodel/dao/redis"
	userredis "github.com/znk_fullstack/server/usercenter/viewmodel/dao/redis"
)

func main() {
	userredis.ConnectRedis(userconf.Dev)

}
