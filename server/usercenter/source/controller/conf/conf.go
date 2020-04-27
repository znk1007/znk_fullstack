package userconf

//Env 环境
type Env string

const (
	//Dev 开发环境
	Dev Env = "dev"
	//Test 测试环境
	Test Env = "test"
	//Prod 生产环境
	Prod Env = "prod"
	//Custom 自定义
	Custom Env = "custom"
)

var e Env

func init() {
	e = Dev
}

//SetEnv 配置环境
func SetEnv(env Env) {
	e = env
}

//RedisSrvConf Redis服务配置
func RedisSrvConf() RedisConf {
	return getRedisConf(e)
}

//SetRedisSrvConf 设置Redis服务配置
func SetRedisSrvConf(env Env, host string, port string, clusters []string, password string) {
	setRedisConf(env, host, port, clusters, password)
}

//GormSrvConf gorm服务配置
func GormSrvConf() GormConf {
	return getGormConf(e)
}

//SetGormSrvConf 设置gorm配置
func SetGormSrvConf(env Env, host string, port string, username string, password string, dialect string, db string) {
	SetGormSrvConf(env, host, port, username, password, dialect, db)
}

//RPCSrvConf 获取当前rpc服务配置
func RPCSrvConf() RPCConf {
	return getRPCConf(e)
}

//SetRPCSrvConf 设置rpc服务配置
func SetRPCSrvConf(env Env, host string, port string) {
	setRPCConf(env, host, port)
}
