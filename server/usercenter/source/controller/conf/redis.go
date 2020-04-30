package userconf

//RedisConf redis配置
type RedisConf struct {
	Host     string
	Port     string
	Clusters []string
	Password string
}

var redismap map[Env]RedisConf

func init() {
	redismap = map[Env]RedisConf{
		Dev: {
			Host:     "localhost",
			Port:     "6379",
			Clusters: nil,
			Password: "man_znk-1007",
		},
		Test: {
			Host:     "47.105.85.107",
			Port:     "6378",
			Clusters: nil,
			Password: "man_znk-1007",
		},
		Prod: {
			Host:     "47.105.85.107",
			Port:     "6378",
			Clusters: nil,
			Password: "man_znk-1007",
		},
	}
}

//getRedisConf 获取当前redis服务配置
func getRedisConf(env Env) RedisConf {
	if r, ok := redismap[env]; ok {
		return r
	}
	return redismap[Dev]
}

//setRedisConf 设置redis服务配置
func setRedisConf(env Env, host string, port string, clusters []string, password string) {
	rds := RedisConf{
		Host:     host,
		Port:     port,
		Clusters: clusters,
		Password: password,
	}
	redismap[env] = rds
}
