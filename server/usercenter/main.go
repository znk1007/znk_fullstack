package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	userconf "github.com/znk_fullstack/server/usercenter/viewmodel/conf"
	usergorm "github.com/znk_fullstack/server/usercenter/viewmodel/dao/gorm"
	userredis "github.com/znk_fullstack/server/usercenter/viewmodel/dao/redis"
	usernet "github.com/znk_fullstack/server/usercenter/viewmodel/net"
)

func main() {
	//日志配置
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.With().Caller().Logger()
	//环境配置
	userconf.SetEnv(userconf.Dev)
	//redis
	userredis.ConnectRedis()
	//mariadb
	usergorm.ConnectMariaDB()
	//rpc
	usernet.RunRPC()
}

/*
sess, err := usermiddleware.SessionID()
if err != nil {
	fmt.Println("get sessionID err: ", err.Error())
	return
}
fmt.Println("sessionID: ", sess)
time.Sleep(time.Duration(time.Second * 2))
expired, err := usermiddleware.ParseSessionID(sess)
if err != nil {
	fmt.Println("parse sessionID err: ", err.Error())
	return
}
fmt.Println("expired: ", expired)
*/

/*

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
*/

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
