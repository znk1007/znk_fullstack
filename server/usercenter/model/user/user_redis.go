package usermodel

import (
	"strconv"
	"time"

	userredis "github.com/znk_fullstack/server/usercenter/viewmodel/dao/redis"
)

// ================================================redis===========================================//

//accOnline 用户是否在线
func accOnline(acc string) (on int) {
	key := "user_online_" + acc
	val, e := userredis.Get(key)
	if e != nil {
		on = 0
		return
	}
	online, e := strconv.Atoi(val)
	if e != nil {
		on = 0
		return
	}
	on = online
	return
}

//setAccOnline 设置用户在线状态
func setAccOnline(acc string, online int) {
	key := "user_online_" + acc
	if online == 0 {
		userredis.Del(key)
		return
	}
	userredis.Set(key, 1, time.Duration(time.Hour*24*7))
}

func setUserActive() {

}

//setAccUserInfo 保存用户基本信息
func setAccUserInfo(acc, userID, phone, email, nickname, photo string) {
	key := "user_info_" + acc
	userredis.HSet(key, "userID", userID, "phone", phone, "email", email, "nickname", nickname, "photo", photo)
}

//accUserID 获取用户ID
func accUserID(acc string) (userID string) {
	key := "user_info_" + acc
	val, err := userredis.HGet(key, "userID")
	if err == nil {
		userID = val
	}
	return
}

//AccUserInfo 获取用户基本信息
func accUserInfo(acc string) (phone, email, nickname, photo string, err error) {
	key := "user_info_" + acc
	vals, e := userredis.HMGet(key, "phone", "email", "nickname", "photo")
	err = e
	if e == nil && len(vals) >= 4 {
		phone, _ = vals[0].(string)
		email, _ = vals[1].(string)
		nickname, _ = vals[2].(string)
		photo, _ = vals[3].(string)
	}
	return
}

//AccRegisted 账号信息是否已注册
func AccRegisted(acc string) (exs bool, ts int64, registed int) {
	key := "user_regist_" + acc
	exs = userredis.Exists(key)
	if exs {
		infos, err := userredis.HMGet(acc, "ts", "registed")
		if err != nil || (infos != nil && len(infos) < 2) {
			return
		}
		tsstr, ok := infos[0].(string)
		if ok {
			ts, err = strconv.ParseInt(tsstr, 10, 64)
		}
		registed, _ = infos[1].(int)
	}
	return
}

//SetAccRegisted 保存注册信息
func SetAccRegisted(acc string, ts string, registed int) (e error) {
	key := "user_regist_" + acc
	e = userredis.HSet(key, "ts", ts, "registed", registed)
	return
}
