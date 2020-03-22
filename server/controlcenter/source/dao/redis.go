package ccdb

import (
	"errors"
	"time"

	"github.com/go-redis/redis/v7"

	"github.com/znk_fullstack/controlcenter/source/config"
)

var rds *redis.Client

func init() {
	dbcf := config.GetDBConfig(config.Redis)
	ops := &redis.Options{
		Addr: dbcf.Host + ":" + dbcf.Port,
	}
	rds = redis.NewClient(ops)
}

//HSet 保存多条数据到哈希表中
func HSet(key string, values ...interface{}) {
	if rds == nil {
		return
	}
	rds.HSet(key, values...)
}

//HGet 根据key获取数据
func HGet(key string, fields string) (string, error) {
	if rds == nil {
		return "", errors.New("redis not exists")
	}
	return rds.HGet(key, fields).Result()
}

//HExists 键值数据是否存在
func HExists(key string, field string) bool {
	if rds == nil {
		return false
	}
	return rds.HExists(key, field).Val()
}

//Set 保存数据
func Set(key string, value interface{}, expiration time.Duration) {
	if rds == nil {
		return
	}
	rds.Set(key, value, expiration)
}

//Get 读取数据
func Get(key string) (string, error) {
	if rds == nil {
		return "", errors.New("redis not exists")
	}
	return rds.Get(key).Result()
}

//Del 删除数据
func Del(keys ...string) {
	if rds == nil {
		return
	}
	rds.Del(keys...)
}

//HDel 删除哈希表数据
func HDel(key string, fields ...string) {
	if rds == nil {
		return
	}
	rds.HDel(key, fields...)
}
