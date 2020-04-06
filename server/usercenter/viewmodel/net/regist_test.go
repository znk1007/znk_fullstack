package usernet

import "testing"

func TestCheckRegistTokenEmptyParams(t *testing.T) {
	resmap := map[string]interface{}{}
	tk, err := rs.checkRegistToken(resmap)
	if err != nil {
		t.Error("CheckRegistTokenEmptyParams: ", err.Error())
	}
	t.Log("TestCheckRegistTokenEmptyParams: ", tk)
}
