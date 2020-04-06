package usernet

import "testing"

func TestCheckRegistTokenEmptyParams(t *testing.T) {
	testregist = true
	resmap := map[string]interface{}{}
	succ := rs.checkRegistToken(resmap)
	if !succ {
		t.Error("CheckRegistTokenEmptyParams failed ")
		return
	}
	t.Log("TestCheckRegistTokenEmptyParams succ")
}
