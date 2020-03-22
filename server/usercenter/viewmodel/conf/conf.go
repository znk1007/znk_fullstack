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
)

var e Env

func init()  {
	e = Dev
}

//SetEnv 配置环境
func SetEnv(env Env)  {
	e = env
}
//GetEnv 获取环境
func GetEnv() Env {
	return e
}


