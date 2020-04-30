package imconf

//UcEnv Rpc客户端环境配置
type UcEnv struct {
	Host string
	Port string
}

var climap map[EnvType]UcEnv

func init() {
	climap = map[EnvType]UcEnv{
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

//setrpcenv RPC用户中心配置
func setrpcenv(et EnvType, host, port string) {
	climap[et] = UcEnv{
		Host: host,
		Port: port,
	}
}

//getUcEnv 获取RPC用户中心配置
func getUcEnv(et EnvType) UcEnv {
	if c, ok := climap[et]; ok {
		return c
	}
	return climap[Dev]
}
