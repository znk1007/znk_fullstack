package usermiddleware

import (
	"testing"
)

func TestAccessCtrlSucc(t *testing.T) {
	cnt := 900000
	fa := NewFreqAccess(1, cnt/2)
	for idx := 0; idx < cnt; idx++ {
		freq := fa.AccessCtrl("test", func() bool {
			return true
		})
		if freq {
			t.Fatal("access too frequence")
		}
	}
}
