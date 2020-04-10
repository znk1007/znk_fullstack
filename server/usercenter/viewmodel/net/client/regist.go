package client

import (
	_ "fmt"
	"time"

	userjwt "github.com/znk_fullstack/server/usercenter/viewmodel/jwt"
)

//JSON字符串，参数密码：password，设备ID：deviceID，平台：platform
type registClient struct {
	Acc       string
	timestamp string
	deviceID  string
	platform  string
	password  string
	appkey    string
}

func (rc *registClient) Token() map[string]interface{} {
	return map[string]interface{}{
		"deviceID":  rc.deviceID,
		"password":  rc.password,
		"platform":  rc.platform,
		"timestamp": rc.timestamp,
		"appkey":    rc.appkey,
	}
	userjwt.CreateUserJWT(time.Duration(time.Minute * 10))
	return nil
}
