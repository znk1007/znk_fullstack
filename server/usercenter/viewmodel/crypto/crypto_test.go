package usercrypto

import "testing"

func TestCBCCrypto(t *testing.T) {
	org := "测试溜溜溜儿"
	res, err := CBCEncrypt(org)
	if err != nil {
		t.Error("cbc encrypt err: ", err)
		return
	}
	org, err = CBCDecrypt(res)
	if err != nil {
		t.Error("cbc decrypt err: ", err)
		return
	}
	t.Log("org: ", org)
}
