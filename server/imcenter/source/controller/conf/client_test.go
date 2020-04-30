package imconf

import "testing"

func TestSetRpcCliConf(t *testing.T) {
	setRPCCliConf(EnvType(3), "127.0.0.1", "8080")
	c := getRPCCliConf(EnvType(3))
	if c.Port == "50051" {
		t.Fatal("set rpc client failed")
	}
	if c.Host != "127.0.0.1" {
		t.Fatal("set rpc client failed")
	}
}
