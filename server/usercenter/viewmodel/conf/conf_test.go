package userconf

import "testing"

func TestEnv(t *testing.T) {
	SetEnv(Dev)
	rds := RedisSrvConf()
	t.Log("redis srv: " + rds)
	gm := GormSrvConf()
	t.Log("gorm srv: " + gm)
	rpc := RPCSrvConf()
	t.Log("rpc srv: " + rpc)
}
