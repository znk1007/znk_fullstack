package redisz

import (
	"sync"
	"time"

	"github.com/znk1007/fullstack/lib/utils/database/redisz/redigo/redis"
)

// Manager 操作客户端
var Manager *RedisManager

// RedisManager redis客户端管理
type RedisManager struct {
	asyncPool *redis.AsyncPool
	syncPool  *redis.Pool
	isSync    bool
	Addr      string
}

var (
	zMap  map[string]*RedisManager
	mutex *sync.RWMutex
)

func init() {
	zMap = make(map[string]*RedisManager)
	mutex = new(sync.RWMutex)
	Manager = GetAsyncRedisz("127.0.0.1:6379")
}

// newSyncPool 初始化同步连接池
func newSyncPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     100,
		MaxActive:   100,
		IdleTimeout: time.Minute * 1,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
	}
}

// newAsyncPool 初始化异步连接池
func newAsyncPool(addr string) *redis.AsyncPool {
	return &redis.AsyncPool{
		Dail: func() (redis.AsyncConn, error) {
			return redis.AsyncDial("tcp", addr)
		},
		MaxGetCount: 1000,
	}
}

// GetSyncRedisz 获取同步客户端
func GetSyncRedisz(addr string) *RedisManager {
	var z *RedisManager
	var ok bool
	mutex.RLock()
	z, ok = zMap[addr]
	mutex.RUnlock()
	if !ok {
		mutex.Lock()
		z, ok = zMap[addr]
		if !ok {
			z = &RedisManager{
				syncPool: newSyncPool(addr),
				isSync:   true,
				Addr:     addr,
			}
			zMap[addr] = z
		}
		mutex.Unlock()
	}
	return z
}

// GetAsyncRedisz 获取异步客户端
func GetAsyncRedisz(addr string) *RedisManager {
	var z *RedisManager
	var ok bool
	mutex.RLock()
	z, ok = zMap[addr]
	mutex.RUnlock()
	if !ok {
		mutex.Lock()
		z, ok = zMap[addr]
		if !ok {
			z = &RedisManager{
				asyncPool: newAsyncPool(addr),
				isSync:    false,
				Addr:      addr,
			}
			zMap[addr] = z
		}
		mutex.Unlock()
	}
	return z
}

// Exists 指定key是否存在
func (z *RedisManager) Exists(key string) (int64, error) {
	var conn redis.Conn
	if z.isSync {
		conn = z.syncPool.Get()
	} else {
		conn = z.asyncPool.Get()
	}
	defer conn.Close()
	reply, err := conn.Do("EXISTS", key)
	if err == nil && reply == nil {
		return 0, nil
	}
	return redis.Int64(reply, err)
}

// Expire 设置指定key数据有效期限
func (z *RedisManager) Expire(key string, expire int64) error {
	var conn redis.Conn
	if z.isSync {
		conn = z.syncPool.Get()
	} else {
		conn = z.asyncPool.Get()
	}
	defer conn.Close()
	_, err := conn.Do("EXPIRE", key, expire)
	return err
}

// SetStringWithExpire 保存字符串并设置有效期限
func (z *RedisManager) SetStringWithExpire(key string, val string, timeoutSeconds int) (string, error) {
	var conn redis.Conn
	if z.isSync {
		conn = z.syncPool.Get()
	} else {
		conn = z.asyncPool.Get()
	}
	defer conn.Close()
	return redis.String(conn.Do("SET", key, val, "EX", timeoutSeconds))
}

// SetString 根据key保存字符串
func (z *RedisManager) SetString(key string, val string) (string, error) {
	var conn redis.Conn
	if z.isSync {
		conn = z.syncPool.Get()
	} else {
		conn = z.asyncPool.Get()
	}
	defer conn.Close()
	return redis.String(conn.Do("SET", key, val))
}

// GetString 获取指定key的字符串数据
func (z *RedisManager) GetString(key string) (string, error) {
	var conn redis.Conn
	if z.isSync {
		conn = z.syncPool.Get()
	} else {
		conn = z.asyncPool.Get()
	}
	defer conn.Close()
	reply, err := conn.Do("GET", key)
	if err == nil && reply == nil {
		return "", nil
	}
	return redis.String(reply, err)
}

// GetBool 获取指定key的布尔数据
func (z *RedisManager) GetBool(key string) (bool, error) {
	var conn redis.Conn
	if z.isSync {
		conn = z.syncPool.Get()
	} else {
		conn = z.asyncPool.Get()
	}
	defer conn.Close()
	reply, err := conn.Do("GET", key)
	if err == nil && reply == nil {
		return false, nil
	}
	return redis.Bool(reply, err)
}

// GetValues 获取指定key转换前数据
func (z *RedisManager) GetValues(key string) ([]interface{}, error) {
	var conn redis.Conn
	if z.isSync {
		conn = z.syncPool.Get()
	} else {
		conn = z.asyncPool.Get()
	}
	defer conn.Close()
	reply, err := conn.Do("GET", key)
	if err == nil && reply == nil {
		return nil, nil
	}
	return redis.Values(reply, err)
}

// HGetString 获取指定hashID及查询域的字符串数据
func (z *RedisManager) HGetString(hashID string, field string) (string, error) {
	var conn redis.Conn
	if z.isSync {
		conn = z.syncPool.Get()
	} else {
		conn = z.asyncPool.Get()
	}
	defer conn.Close()
	reply, err := conn.Do("HGET", hashID, field)
	if err == nil && reply == nil {
		return "", nil
	}
	return redis.String(reply, err)
}

// HGetAllStringMap 获取指定hashID所有数据
func (z *RedisManager) HGetAllStringMap(hashID string) (map[string]string, error) {
	var conn redis.Conn
	if z.isSync {
		conn = z.syncPool.Get()
	} else {
		conn = z.asyncPool.Get()
	}
	defer conn.Close()
	return redis.StringMap(conn.Do("HGETALL", hashID))
}

// HGetAllStructValue 获取指定hashID的对象数据
func (z *RedisManager) HGetAllStructValue(hashID string, dst interface{}) error {
	var conn redis.Conn
	if z.isSync {
		conn = z.syncPool.Get()
	} else {
		conn = z.asyncPool.Get()
	}
	defer conn.Close()
	v, err := redis.Values(conn.Do("HGETALL", hashID))
	if err != nil {
		return err
	}
	if err = redis.ScanStruct(v, dst); err != nil {
		return err
	}
	return nil
}

// HSet 根据hashID保存哈希数据
func (z *RedisManager) HSet(hashID string, field string, val string) error {
	var conn redis.Conn
	if z.isSync {
		conn = z.syncPool.Get()
	} else {
		conn = z.asyncPool.Get()
	}
	defer conn.Close()
	_, err := conn.Do("HSET", hashID, field, val)
	return err
}

// HMSet 批量保存哈希数据
func (z *RedisManager) HMSet(args ...interface{}) error {
	var conn redis.Conn
	if z.isSync {
		conn = z.syncPool.Get()
	} else {
		conn = z.asyncPool.Get()
	}
	defer conn.Close()
	_, err := conn.Do("HMSET", args...)
	return err
}

// Del 删除指定key的数据
func (z *RedisManager) Del(key string) (int64, error) {
	var conn redis.Conn
	if z.isSync {
		conn = z.syncPool.Get()
	} else {
		conn = z.asyncPool.Get()
	}
	defer conn.Close()
	reply, err := conn.Do("DEL", key)
	if err == nil && reply == nil {
		return 0, nil
	}
	return redis.Int64(reply, err)
}
