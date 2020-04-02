package usergorm

import (
	"testing"

	userconf "github.com/znk_fullstack/server/usercenter/viewmodel/conf"
)

func TestConnectDB(t *testing.T) {
	ConnectMariaDB(userconf.Dev)
}
