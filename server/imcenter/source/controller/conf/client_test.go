package imconf

import "testing"

func TestSetrpcenv(t *testing.T) {
	setrpcenv(EnvType(3), "127.0.0.1", "8080")
	c := getUcEnv(EnvType(3))
	if c.Port == "50051" {
		t.Fatal("set rpc client failed")
	}
	if c.Host != "127.0.0.1" {
		t.Fatal("set rpc client failed")
	}
}
