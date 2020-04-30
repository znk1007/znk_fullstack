package imconf

//RPCCliConf Rpc客户端环境配置
type RPCCliConf struct {
	Host string
	Port string
}

var climap map[EnvType]RPCCliConf

func init() {
	climap = map[EnvType]RPCCliConf{
		Dev: {
			Host: "localhost",
			Port: "50051",
		},
		Test: {
			Host: "localhost",
			Port: "50051",
		},
		Prod: {
			Host: "47.105.85.107",
			Port: "50051",
		},
	}
}

//setRPCCliConf RPC用户中心配置
func setRPCCliConf(et EnvType, host, port string) {
	climap[et] = RPCCliConf{
		Host: host,
		Port: port,
	}
}

//getRPCCliConf 获取RPC用户中心配置
func getRPCCliConf(et EnvType) RPCCliConf {
	if c, ok := climap[et]; ok {
		return c
	}
	return climap[Dev]
}
