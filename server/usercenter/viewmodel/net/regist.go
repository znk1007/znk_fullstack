package usernet

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
	usercrypto "github.com/znk_fullstack/server/usercenter/viewmodel/crypto"
	userjwt "github.com/znk_fullstack/server/usercenter/viewmodel/jwt"
	userpayload "github.com/znk_fullstack/server/usercenter/viewmodel/payload"
	"google.golang.org/grpc"
)

var rs *registService

func init() {
	rs = &registService{
		resChan: make(chan *userproto.RegistRes),
	}
}

type registState struct {
	succ bool
	msg  string
}

//RegistService 注册
type registService struct {
	req     *userproto.RegistReq
	resChan chan *userproto.RegistRes
}

func (s registService) Do() {
	go s.handleRegist()
}

func (s registService) handleRegist() {
	req := s.req
	acc := req.GetAccount()
	fmt.Println(acc)

}

/*
参数密码：password[CBCEncrypt]，
设备ID：deviceID，
平台：platform[web,iOS,Android]，
时间戳：timestamp，
应用标识：appkey
*/

func checkRegistToken(reqMap map[string]interface{}) (tk string, err error) {
	var deviceID string
	var password string
	var platform string
	dID, ok := reqMap["deviceID"]
	deviceID = dID.(string)
	if !ok || len(deviceID) == 0 {
		log.Info().Msg("deviceID cannot be empty")
		tk, err = generateRegistToken("", http.StatusBadRequest, "deviceID cannot be empty")
		return
	}
	plf, ok := reqMap["platform"]
	platform = plf.(string)
	if !ok || len(platform) == 0 {
		log.Info().Msg("platform cannot be empty")
		tk, err = generateRegistToken("", http.StatusBadRequest, "platform cannot be empty")
		return
	}
	psd, ok := reqMap["password"]
	password = psd.(string)
	if !ok || len(password) == 0 {
		log.Info().Msg("password cannot be empty")
		tk, err = generateRegistToken("", http.StatusBadRequest, "password cannot be empty")
		return
	}
	_, e := usercrypto.CBCEncrypt(password)
	if e != nil {
		log.Info().Msg("encrypt password err: " + e.Error())
		tk, err = generateRegistToken("", http.StatusInternalServerError, "interval server error")
		return
	}

	return
}

/*
用户ID：userID，
时间戳：timestamp，
状态码：code，
反馈消息：message
*/
func generateRegistToken(userID string, code int, msg string) (tk string, err error) {
	ts := time.Now().Unix()
	resMap := map[string]interface{}{
		"timestamp": ts,
		"code":      code,
		"message":   msg,
	}
	if code == http.StatusOK {
		resMap["userID"] = userID
	}
	tk, err = userjwt.CreateToken(time.Duration(time.Minute*1), resMap)
	return
}

//RegisterRegistServer 注册到注册请求服务
func RegisterRegistServer(srv *grpc.Server) {
	userproto.RegisterRegistSrvServer(srv, rs)
}

//UserReigst 注册
func (s registService) UserReigst(ctx context.Context, req *userproto.RegistReq) (*userproto.RegistRes, error) {
	userpayload.Pool.WriteHandler(func(jq chan userpayload.Job) {
		s.req = req
		jq <- s
	})
	for {
		select {
		case res := <-s.resChan:
			return res, nil
		}
	}
}
