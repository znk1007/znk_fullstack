package imconf

//EnvType 环境类型
type EnvType int

const (
	//Dev 开发环境
	Dev EnvType = iota
	//Test 测试环境
	Test EnvType = 1
	//Prod 生成环境
	Prod EnvType = 2
)

//curEt 当前环境
var curEt EnvType

//SetCurEnvConf 设置当前环境
func SetCurEnvConf(et EnvType) {
	curEt = et
}

//SetUcEnv 配置用户中心环境
func SetUcEnv(et EnvType, host, port string) {
	setrpcenv(et, host, port)
}

//GetUcEnv 获取用户中心环境配置
func GetUcEnv() UcEnv {
	return getUcEnv(curEt)
}
