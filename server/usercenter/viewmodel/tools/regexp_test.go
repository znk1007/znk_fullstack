package tools

import "testing"

func TestValidateEmail(t *testing.T) {
	em1 := "111@qq.com"
	succ := VerifyEmail(em1)
	if !succ {
		t.Error("invalidate email")
	}
}

func TestInvalidateEmail(t *testing.T) {
	em1 := "13011111111.com"
	succ := VerifyEmail(em1)
	if !succ {
		t.Error("invalidate email")
	}
}

func TestValidatePhone(t *testing.T) {
	phone := "13000000000"
	succ := VerifyPhone(phone)
	if !succ {
		t.Error("invalidate email")
	}
}

func TestInvalidatePhone(t *testing.T) {
	phone := "12300000000"
	succ := VerifyPhone(phone)
	if !succ {
		t.Error("invalidate email")
	}
}
