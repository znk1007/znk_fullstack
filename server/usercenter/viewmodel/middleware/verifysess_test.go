package usermiddleware

import "testing"

func TestSession(t *testing.T) {
	sess, err := SessionID("test")
	if err != nil {
		t.Fatal("get sessionID err: ", err.Error())
		return
	}
	t.Logf("sessionID: %v", sess)
	expired, err := ParseSessionID(sess, "test")
	if err != nil {
		t.Fatal("parse sessionID err: ", err.Error())
		return
	}
	t.Log("expired: ", expired)
}
