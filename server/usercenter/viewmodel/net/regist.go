package usernet

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/znk_fullstack/server/usercenter/model"
	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
	userredis "github.com/znk_fullstack/server/usercenter/viewmodel/dao/redis"
	usermiddleware "github.com/znk_fullstack/server/usercenter/viewmodel/middleware"
	userpayload "github.com/znk_fullstack/server/usercenter/viewmodel/payload"
	"google.golang.org/grpc"
)

var testregist = false

var rs *registService
var check usermiddleware.CheckToken

const (
	expiredInterval = time.Minute * 2
)

func init() {
	rs = &registService{
		resChan: make(chan registResponse),
	}
	check = usermiddleware.Initialize(expiredInterval)
}

//registResponse 注册响应
type registResponse struct {
	res *userproto.RegistRes
	err error
}

//RegistService 注册
type registService struct {
	req     *userproto.RegistReq
	resChan chan registResponse
}

func (s registService) Do() {
	go s.handleRegist()
}

func (s registService) handleRegist() {
	req := s.req
	acc := req.GetAccount()
	if len(acc) == 0 {
		log.Info().Msg("account cannot be empty")
		s.makeToken("", http.StatusAccepted, "account cannot be empty")
		return
	}
	//redis 第一波墙，防止频繁操作数据库
	exs, oldTS, registed := model.AccRegisted(acc)
	if exs {
		if registed == 1 {
			s.makeToken(acc, http.StatusAccepted, "user has registed:")
			return
		}
		if oldTS < 0 {
			s.makeToken(acc, http.StatusAccepted, "param `timestamp` is error type")
			return
		}
		ts := time.Now().Unix()
		if ts-oldTS > int64(expiredInterval) {
			s.makeToken(acc, http.StatusAccepted, "please regist later on")
			return
		}
	}

	res, expired, e := check.Verify(req.GetToken())
	if e != nil {
		log.Info().Msg(e.Error())
		s.makeToken(acc, http.StatusAccepted, e.Error())
		return
	}

	if !expired {
		s.makeToken(acc, http.StatusAccepted, "please regist later on")
		return
	}

	psd, ok := res["password"]
	if !ok || psd == nil {
		log.Info().Msg("password cannot be empty")
		s.makeToken("", http.StatusBadRequest, "password cannot be empty")
		return
	}
	password := psd.(string)
	fmt.Println(password)
	succ := s.checkRegistToken(res)
	if !succ {
		log.Info().Msg(e.Error())
		s.makeToken(acc, http.StatusAccepted, e.Error())
		return
	}

	var rgd int
	userredis.HSet(acc, "ts", string(time.Now().Unix()), "registed", rgd)
}

/*
参数密码：password[CBCEncrypt]，
设备ID：deviceID，
平台：platform[web,iOS,Android]，
时间戳：timestamp，
应用标识：appkey
*/

func (s registService) checkRegistToken(reqMap map[string]interface{}) bool {
	var deviceID string
	var platform string
	dID, ok := reqMap["deviceID"]
	if !ok || dID == nil {
		log.Info().Msg("deviceID cannot be empty")
		s.makeToken("", http.StatusBadRequest, "deviceID cannot be empty")
		return false
	}
	deviceID = dID.(string)
	if len(deviceID) == 0 {
		log.Info().Msg("deviceID cannot be empty")
		s.makeToken("", http.StatusBadRequest, "deviceID cannot be empty")
		return false
	}

	plf, ok := reqMap["platform"]
	if !ok || plf == nil {
		log.Info().Msg("platform cannot be empty")
		s.makeToken("", http.StatusBadRequest, "platform cannot be empty")
		return false
	}
	platform = plf.(string)
	if len(platform) == 0 {
		log.Info().Msg("platform cannot be empty")
		s.makeToken("", http.StatusBadRequest, "platform cannot be empty")
		return false
	}
	return true
}

/*
用户ID：userID，
时间戳：timestamp，
状态码：code，
反馈消息：message
*/
func (s registService) makeToken(userID string, code int, msg string) {
	if testregist {
		return
	}
	ts := time.Now().Unix()
	resMap := map[string]interface{}{
		"timestamp": ts,
		"code":      code,
		"message":   msg,
	}
	if code == http.StatusOK {
		resMap["userID"] = userID
	} else {
		log.Info().Msg(msg)
	}
	tk, err := check.Generate(resMap)
	res := registResponse{
		res: &userproto.RegistRes{
			Account: "",
			Token:   tk,
		},
		err: err,
	}
	s.resChan <- res
	return
}

//registerRegistServer 注册到注册请求服务
func registerRegistServer(srv *grpc.Server) {
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
			return res.res, res.err
		}
	}
}
