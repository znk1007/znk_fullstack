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

//EnvConf 环境配置
type EnvConf struct {
	Et EnvType //Et 环境变量
}
