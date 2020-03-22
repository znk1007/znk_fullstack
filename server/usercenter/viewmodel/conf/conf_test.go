package userconf

import "testing"

func TestEnv(t *testing.T)  {
	SetEnv(Dev)
	e := GetEnv()
	t.Log("env: "+ e)
}