package usermiddleware

import "testing"

func TestSession(t *testing.T) {
	sess, err := DefaultSess.SessionID("test")
	if err != nil {
		t.Fatal("get sessionID err: ", err.Error())
		return
	}
	t.Logf("sessionID: %v", sess)
	expired, err := DefaultSess.Parse(sess, "test")
	if err != nil {
		t.Fatal("parse sessionID err: ", err.Error())
		return
	}
	t.Log("expired: ", expired)
}
