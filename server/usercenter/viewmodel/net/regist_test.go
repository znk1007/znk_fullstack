package usernet

import "testing"

func TestverifyDeviceEmptyParams(t *testing.T) {
	e := saveCurrentDevice("", "", "")
	if e != nil {
		t.Error("verifyDevice failed ")
		return
	}
	t.Log("verifyDevice succ")
}

func TestRegist(t *testing.T) {
	test := map[string]interface{}{
		"key1": "test1",
		"key2": "test2",
		"key3": "test3",
	}
	t1 := test["key1"]
	if t1 != "test1" {
		t.Fatal("t1 is equal to test1")
	}
	delete(test, "key1")
	t1 = test["key1"]
	if t1 != nil {
		t.Fatal("t1 is not equal to test1")
	}

	t2 := test["key2"]
	if t2 == nil {
		t.Fatal("t2 is not nil")
	}

}
