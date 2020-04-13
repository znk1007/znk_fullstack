package userredis

import (
	"errors"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/rs/zerolog/log"
	userconf "github.com/znk_fullstack/server/usercenter/viewmodel/conf"
)

//https://blog.csdn.net/itcats_cn/article/details/82391719

//clstrRds 集群客户端
var clstrRds *redis.ClusterClient

//nclstrRds 非集群客户端
var nclstrRds *redis.Client

//ConnectRedis 连接Redis
func ConnectRedis() {
	initRedis()
}

//initRedis 初始化对象
func initRedis() {
	rdsConf := userconf.RedisSrvConf()
	psw := rdsConf.Password
	if rdsConf.Clusters != nil && len(rdsConf.Clusters) > 0 {
		ops := &redis.ClusterOptions{
			Addrs:    rdsConf.Clusters,
			Password: psw,
		}
		clstrRds = redis.NewClusterClient(ops)
	} else {
		ops := &redis.Options{
			Addr:     rdsConf.Host + ":" + rdsConf.Port,
			Password: rdsConf.Password,
		}
		nclstrRds = redis.NewClient(ops)
	}
	log.Info().Msg("connect redis success")
}

//Exists key是否存在
func Exists(key ...string) bool {
	var idx int64
	var err error
	if nclstrRds != nil {
		if err := checkRds(); err != nil {
			log.Info().Msg(err.Error())
			return false
		}
		idx, err = nclstrRds.Exists(key...).Result()
		if err != nil {
			log.Info().Msg(err.Error())
			return false
		}
	} else {
		if err := checkRds(); err != nil {
			log.Info().Msg(err.Error())
			return false
		}
		idx, err = clstrRds.Exists(key...).Result()
		if err != nil {
			log.Info().Msg(err.Error())
			return false
		}
	}

	return idx == 1
}

//Set 保存数据
func Set(key string, val interface{}, exp time.Duration) (e error) {
	if err := checkRds(); err != nil {
		e = err
		return
	}
	if clstrRds != nil {
		e = clstrRds.Set(key, val, exp).Err()
	} else {
		e = nclstrRds.Set(key, val, exp).Err()
	}
	return
}

//Get 取值
func Get(key string) (val string, err error) {
	if err := checkRds(); err != nil {
		return "", err
	}

	if clstrRds != nil {
		val, err = clstrRds.Get(key).Result()
	} else {
		val, err = nclstrRds.Get(key).Result()
	}
	return
}

//Del 删除
func Del(keys ...string) (e error) {
	if err := checkRds(); err != nil {
		e = err
		return
	}
	if clstrRds != nil {
		e = clstrRds.Del(keys...).Err()
	} else {
		e = nclstrRds.Del(keys...).Err()
	}
	return
}

//HSet hash存值
//HSet accepts values in following formats:
//   - HMSet("myhash", "key1", "value1", "key2", "value2")
//   - HMSet("myhash", []string{"key1", "value1", "key2", "value2"})
//   - HMSet("myhash", map[string]interface{}{"key1": "value1", "key2": "value2"})
//
// Note that it requires Redis v4 for multiple field/value pairs support.
func HSet(key string, values ...interface{}) (e error) {
	if err := checkRds(); err != nil {
		e = err
		return err
	}
	if clstrRds != nil {
		e = clstrRds.HSet(key, values...).Err()
	} else {
		e = nclstrRds.HSet(key, values...).Err()
	}
	return
}

//HSetNX 设置不存在的field，如果 key 不存在，一个新哈希表被创建并执行 HSETNX 命令。
func HSetNX(key string, field string, value interface{}) (succ bool, e error) {
	if err := checkRds(); err != nil {
		succ = false
		e = err
		return
	}
	if clstrRds != nil {
		succ, e = clstrRds.HSetNX(key, field, value).Result()
	} else {
		succ, e = nclstrRds.HSetNX(key, field, value).Result()
	}
	return
}

//HGet hash读取
func HGet(key string, field string) (val string, err error) {
	if e := checkRds(); e != nil {
		val = ""
		err = e
		return
	}
	if clstrRds != nil {
		val, err = clstrRds.HGet(key, field).Result()
	} else {
		val, err = nclstrRds.HGet(key, field).Result()
	}
	return
}

//HDel hash删除
func HDel(key string, field ...string) (e error) {
	if err := checkRds(); err != nil {
		e = err
		return
	}
	if clstrRds != nil {
		e = clstrRds.HDel(key, field...).Err()
	} else {
		e = nclstrRds.HDel(key, field...).Err()
	}
	return redis.Nil
}

// MSet is like Set but accepts multiple values:
//   - MSet("key1", "value1", "key2", "value2")
//   - MSet([]string{"key1", "value1", "key2", "value2"})
//   - MSet(map[string]interface{}{"key1": "value1", "key2": "value2"})
func MSet(values ...interface{}) (e error) {
	if err := checkRds(); err != nil {
		e = err
		return
	}
	if clstrRds != nil {
		e = clstrRds.MSet(values...).Err()
	} else {
		e = nclstrRds.MSet(values...).Err()
	}
	return
}

//MGet 取值
func MGet(key ...string) (slc []interface{}, err error) {
	if e := checkRds(); e != nil {
		slc = nil
		err = e
		return
	}
	if clstrRds != nil {
		slc, err = clstrRds.MGet(key...).Result()
	} else {
		slc, err = nclstrRds.MGet(key...).Result()
	}
	return
}

//HMGet 取值
func HMGet(key string, field ...string) (slc []interface{}, err error) {
	if e := checkRds(); e != nil {
		slc = nil
		err = e
		return
	}
	if clstrRds != nil {
		slc, err = clstrRds.HMGet(key, field...).Result()
	} else {
		slc, err = nclstrRds.HMGet(key, field...).Result()
	}
	return
}

//checkRds 校验rds实例
func checkRds() error {
	if clstrRds == nil && nclstrRds == nil {
		log.Info().Msg("redis client not init")
		return errors.New("redis client not init")
	}
	return nil
}

//readFile 获取指定文件地址
// func readFile(relativePath string) string {
// 	_, curPath, _, _ := runtime.Caller(1)
// 	return path.Join(path.Dir(curPath) + "/" + relativePath)
// }
