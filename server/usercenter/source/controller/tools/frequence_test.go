package usertools

import (
	"testing"
	"time"
)

var freq *Freq

func TestFreq(t *testing.T) {
	if freq == nil {
		freq = NewFreq(2)
	}
	now := time.Now().Unix()
	expired := freq.Expired("test", "userID", now)
	if expired {
		t.Log("expired")
	} else {
		t.Log("not expired")
	}
}

func BenchmarkFreq(t *testing.B) {
	if freq == nil {
		freq = NewFreq(1)
		t.Log("freq nil")
	}
	now := time.Now().Unix()
	expired := freq.Expired("test", "userID", now)
	t.Log("expired 1: ", expired)
	time.Sleep(time.Second * 2)
	now = time.Now().Unix()
	expired = freq.Expired("test", "userID", now)
	t.Log("expired 2: ", expired)
	time.Sleep(time.Second * 2)
	now = time.Now().Unix()
	expired = freq.Expired("test", "userID", now)
	t.Log("expired 3: ", expired)
}
