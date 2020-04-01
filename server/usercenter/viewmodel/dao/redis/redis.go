package userredis

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/rs/zerolog/log"
	userconf "github.com/znk_fullstack/server/usercenter/viewmodel/conf"
)

//https://blog.csdn.net/itcats_cn/article/details/82391719
type redisConf struct {
	Test redisInfo `json:"test"`
	Dev  redisInfo `json:"dev"`
	Prod redisInfo `json:"prod"`
}

type redisInfo struct {
	Host     string   `json:"host"`
	Port     string   `json:"port"`
	Clusters []string `json:"cluster"`
}

var rds *redis.ClusterClient

//ConnectRedis 连接Redis
func ConnectRedis(envir userconf.Env) {
	conf := readRedisConfig()
	switch envir {
	case userconf.Dev:
		initCluster(conf.Dev.Clusters)
	case userconf.Prod:
		initCluster(conf.Prod.Clusters)
	case userconf.Test:
		initCluster(conf.Test.Clusters)
	}

}

//initCluster 初始化集群对象
func initCluster(addrs []string) {
	ops := &redis.ClusterOptions{
		Addrs: addrs,
	}
	rds = redis.NewClusterClient(ops)
	fmt.Println("redis: ", rds)
}

func readRedisConfig() *redisConf {
	rc := &redisConf{}
	fp := readFile("redis.json")
	bs, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Info().Msg(err.Error())
		panic(err.Error())
	}
	err = json.Unmarshal(bs, rc)
	if err != nil {
		panic(err.Error())
	}
	return rc
}

//Exists key是否存在
func Exists(key ...string) bool {
	if err := checkRds(); err != nil {
		log.Info().Msg(err.Error())
		return false
	}
	idx, err := rds.Exists(key...).Result()
	if err != nil {
		log.Info().Msg(err.Error())
		return false
	}
	return idx == 1
}

//Set 保存数据
func Set(key string, val interface{}, exp time.Duration) error {
	if err := checkRds(); err != nil {
		return err
	}
	rds.Set(key, val, exp)
	return nil
}

//Get 取值
func Get(key string) (val string, err error) {
	if err := checkRds(); err != nil {
		return "", err
	}

	val, err = rds.Get(key).Result()
	return
}

//Del 删除
func Del(key ...string) error {
	if err := checkRds(); err != nil {
		return err
	}
	rds.Del(key...)
	return redis.Nil
}

//HSet hash存值
//HSet accepts values in following formats:
//   - HMSet("myhash", "key1", "value1", "key2", "value2")
//   - HMSet("myhash", []string{"key1", "value1", "key2", "value2"})
//   - HMSet("myhash", map[string]interface{}{"key1": "value1", "key2": "value2"})
//
// Note that it requires Redis v4 for multiple field/value pairs support.
func HSet(key string, values ...interface{}) error {
	if err := checkRds(); err != nil {
		return err
	}
	rds.HSet(key, values...)
	return redis.Nil
}

//HGet hash读取
func HGet(key string, field string) (val string, err error) {
	if e := checkRds(); e != nil {
		val = ""
		err = e
		return
	}
	val, err = rds.HGet(key, field).Result()
	err = nil
	return
}

//HDel hash删除
func HDel(key string, field ...string) error {
	if err := checkRds(); err != nil {
		return err
	}
	rds.HDel(key, field...)
	return redis.Nil
}

// MSet is like Set but accepts multiple values:
//   - MSet("key1", "value1", "key2", "value2")
//   - MSet([]string{"key1", "value1", "key2", "value2"})
//   - MSet(map[string]interface{}{"key1": "value1", "key2": "value2"})
func MSet(values ...interface{}) error {
	if err := checkRds(); err != nil {
		return err
	}
	rds.MSet(values...)
	return redis.Nil
}

//MGet 取值
func MGet(key ...string) (slc []interface{}, err error) {
	if e := checkRds(); e != nil {
		slc = nil
		err = e
		return
	}
	slc, err = rds.MGet(key...).Result()
	return
}

//HMGet 取值
func HMGet(key string, field ...string) (slc []interface{}, err error) {
	if e := checkRds(); e != nil {
		slc = nil
		err = e
		return
	}
	slc, err = rds.HMGet(key, field...).Result()
	return
}

//checkRds 校验rds实例
func checkRds() error {
	if rds == nil {
		log.Info().Msg("redis client not init")
		return errors.New("redis client not init")
	}
	return nil
}

//readFile 获取指定文件地址
func readFile(relativePath string) string {
	_, curPath, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(curPath) + "/" + relativePath)
}
