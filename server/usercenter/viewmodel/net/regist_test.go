package usernet

import "testing"

func TestMakeDeviceEmptyParams(t *testing.T) {
	testregist = true
	succ, _ := rs.makeDevice("", "", "")
	if !succ {
		t.Error("makeDevice failed ")
		return
	}
	t.Log("makeDevice succ")
}

func TestRegist() {

}
