package usermiddleware

import "testing"

func TestSession(t *testing.T) {
	sess, err := SessionID()
	if err != nil {
		t.Fatal("get sessionID err: ", err.Error())
		return
	}
	t.Logf("sessionID: %v", sess)
	expired, err := ParseSessionID(sess)
	if err != nil {
		t.Fatal("parse sessionID err: ", err.Error())
		return
	}
	t.Log("expired: ", expired)
}
