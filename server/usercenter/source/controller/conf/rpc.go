package userconf

//RPCConf rpc配置
type RPCConf struct {
	Host string
	Port string
}

var rpcmap map[Env]RPCConf

func init() {
	rpcmap = map[Env]RPCConf{
		Dev: {
			Host: "localhost",
			Port: "50051",
		},
		Test: {
			Host: "localhost",
			Port: "50051",
		},
		Prod: {
			Host: "localhost",
			Port: "50051",
		},
	}
}

//getRPCConf 获取rpc配置
func getRPCConf(env Env) RPCConf {
	if r, ok := rpcmap[env]; ok {
		return r
	}
	return rpcmap[Dev]
}

//setRPCConf 设置rpc配置
func setRPCConf(env Env, host string, port string) {
	rpcmap[env] = RPCConf{
		Host: host,
		Port: port,
	}
}
