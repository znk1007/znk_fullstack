package usermodel

import (
	"strconv"
	"time"

	userredis "github.com/znk_fullstack/server/usercenter/source/controller/dao/redis"
	userproto "github.com/znk_fullstack/server/usercenter/source/model/protos/generated"
)

const (
	userInfoPrefix   = "user_info_"     //保存用户信息key前缀
	userOnlinePrefix = "user_online_"   //保存用户在线状态
	userRegistPrefix = "user_registed_" //保存用户注册状态
)

//redisCreateUser redis保存用户数据
func redisCreateUser(acc, userID, password, phone, email, nickname, photo, createdAt, updatedAt string) (err error) {
	key := userInfoPrefix + acc
	err = userredis.HSet(
		key,
		"userID", userID,
		"password", password,
		"phone", phone,
		"email", email,
		"nickname", nickname,
		"photo", phone,
		"updatedAt", updatedAt,
		"createdAt", createdAt,
		"active", "1",
		"permission", string(User),
	)
	return
}

//redisUserExists 用户是否存在
func redisUserExists(acc string) (exists bool) {
	key := userInfoPrefix + acc
	exists = userredis.Exists(key)
	return
}

//redisGetUser 获取用户基本信息
func redisGetUser(acc string) (user UserDB, err error) {
	key := userInfoPrefix + acc
	vals, e := userredis.HMGet(key, "phone", "email", "nickname", "photo", "updatedAt", "createdAt", "permission")
	err = e
	if e == nil && len(vals) > 5 {
		phone, _ := vals[0].(string)
		email, _ := vals[1].(string)
		nickname, _ := vals[2].(string)
		photo, _ := vals[3].(string)
		updatedAt, _ := vals[4].(string)
		createdAt, _ := vals[5].(string)
		per, _ := vals[6].(Permission)
		u := &userproto.User{
			Phone:     phone,
			Email:     email,
			Nickname:  nickname,
			Photo:     photo,
			UpdatedAt: updatedAt,
			CreatedAt: createdAt,
		}
		user = UserDB{
			Per:  per,
			User: u,
		}
	}
	return
}

//redisGetUserID 获取用户ID
func redisGetUserID(acc string) (userID string, err error) {
	key := userInfoPrefix + acc
	var val string
	val, err = userredis.HGet(key, "userID")
	if err == nil {
		userID = val
	}
	return
}

//redisSetUserPassword 更新用户密码
func redisSetUserPassword(acc, psw string) (err error) {
	key := userInfoPrefix + acc
	err = userredis.HSet(key, psw)
	return
}

//redisUserPassword 获取密码
func redisUserPassword(acc string) (password string, err error) {
	key := userInfoPrefix + acc
	var val string
	val, err = userredis.HGet(key, "password")
	if err == nil {
		password = val
	}
	return
}

//redisUserOnline 用户是否在线
func redisUserOnline(acc string) (on int, err error) {
	key := userOnlinePrefix + acc
	var val string
	val, _ = userredis.Get(key)
	if len(val) == 0 {
		on = 0
		return
	}
	on, err = strconv.Atoi(val)
	return
}

//redisSetUserOnline 设置用户在线状态
func redisSetUserOnline(acc string, online int) (err error) {
	key := userOnlinePrefix + acc
	if online == 0 {
		err = userredis.Del(key)
		return
	}
	err = userredis.Set(key, "1", time.Duration(time.Hour*24*7))
	return
}

//redisSetUserActive 用户是否激活状态
func redisSetUserActive(acc string, active int) (e error) {
	key := userInfoPrefix + acc
	e = userredis.HSet(key, "active", string(active))
	return
}

//redisUserActive 用户是否激活状态
func redisUserActive(acc string) (active int, err error) {
	key := userInfoPrefix + acc
	var val string
	val, err = userredis.HGet(key, "active")
	if err != nil {
		active = 0
		return
	}
	active, err = strconv.Atoi(val)
	return
}

//redisSetUserPermission 设置用户权限
func redisSetUserPermission(acc string, per Permission) (err error) {
	key := userInfoPrefix + acc
	err = userredis.HSet(key, "permission", string(per))
	return
}

//redisUserPermission 获取用户权限
func redisUserPermission(acc string) (permission Permission, err error) {
	key := userInfoPrefix + acc
	val, e := userredis.HGet(key, "permission")
	if e != nil {
		err = e
		return
	}
	per, e := strconv.Atoi(val)
	permission = Permission(per)
	return
}

//redisSetUserPhone 更新手机号
func redisSetUserPhone(acc, phone string) (err error) {
	key := userInfoPrefix + acc
	err = userredis.HSet(key, "phone", phone)
	return
}

//redisSetUserNickname 更新昵称
func redisSetUserNickname(acc, nickname string) (err error) {
	key := userInfoPrefix + acc
	err = userredis.HSet(key, "nickname", nickname)
	return
}

//redisUserRegisted 账号信息是否已注册
func redisUserRegisted(acc string) (exs bool, ts int64, registed int) {
	key := userRegistPrefix + acc
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

//redisSetUserRegisted 保存注册信息
func redisSetUserRegisted(acc string, ts string, registed int) (e error) {
	key := userRegistPrefix + acc
	e = userredis.HSet(key, "ts", ts, "registed", registed)
	return
}