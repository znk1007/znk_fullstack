package usernet

import "testing"

func TestMakeDeviceEmptyParams(t *testing.T) {
	testregist = true
	resmap := map[string]interface{}{}
	succ, _ := rs.makeDevice(resmap)
	if !succ {
		t.Error("makeDevice failed ")
		return
	}
	t.Log("makeDevice succ")
}
