package userconf

import "testing"

func TestEnv(t *testing.T) {
	SetEnv(Dev)
	rds := RedisSrvConf()
	t.Log(rds)
	gm := GormSrvConf()
	t.Log(gm)
	rpc := RPCSrvConf()
	t.Log(rpc)
}
