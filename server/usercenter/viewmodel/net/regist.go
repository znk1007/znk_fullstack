package usernet

import (
	"context"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/znk_fullstack/server/usercenter/model"
	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
	userredis "github.com/znk_fullstack/server/usercenter/viewmodel/dao/redis"
	userGenID "github.com/znk_fullstack/server/usercenter/viewmodel/generateId"
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
		s.makeToken("", "", http.StatusAccepted, "account cannot be empty")
		return
	}
	//redis 第一波墙，防止频繁操作数据库
	exs, oldTS, registed := model.AccRegisted(acc)
	//解析校验token
	res, dID, plf, expired, e := check.Verify(req.GetToken())
	if e != nil {
		log.Info().Msg(e.Error())
		s.makeToken(acc, "", http.StatusAccepted, e.Error())
		return
	}
	//如果存在redis中，曾调过注册方法
	if exs {
		//如果已注册
		if registed == 1 {
			s.makeToken(acc, "", http.StatusAccepted, "user has registed:")
			return
		}
		if !expired {
			s.makeToken(acc, "", http.StatusAccepted, "please regist later on")
			return
		}

		if oldTS < 0 {
			s.makeToken(acc, "", http.StatusAccepted, "miss param `timestamp`")
			return
		}
		ts := time.Now().Unix()
		if ts-oldTS < int64(expiredInterval) {
			s.makeToken(acc, "", http.StatusAccepted, "please regist later on")
			return
		}
	}

	psd, ok := res["password"].(string)
	if !ok || len(psd) == 0 {
		log.Info().Msg("password cannot be empty")
		s.makeToken("", "", http.StatusBadRequest, "password cannot be empty")
		return
	}

	succ, userID := s.makeDevice(dID, plf)
	if !succ {
		log.Info().Msg(e.Error())
		s.makeToken(acc, "", http.StatusAccepted, e.Error())
		return
	}
	user := &userproto.User{
		UserID:  userID,
		Account: acc,
	}
	model.CreateUser(user)
	var rgd int
	userredis.HSet(acc, "ts", string(time.Now().Unix()), "registed", rgd)
}

/*
设备ID：deviceID，
平台：platform[web,iOS,Android]，
用户ID：userID，
应用标识：appkey
*/

func (s registService) makeDevice(deviceID string, platform string) (succ bool, userID string) {
	var ok bool
	succ = false
	userID = ""
	if !ok || len(deviceID) == 0 {
		log.Info().Msg("deviceID cannot be empty")
		s.makeToken("", "", http.StatusBadRequest, "deviceID cannot be empty")
		return
	}
	if !ok || len(platform) == 0 {
		log.Info().Msg("platform cannot be empty")
		s.makeToken("", "", http.StatusBadRequest, "platform cannot be empty")
		return
	}

	userID = userGenID.GenerateID()
	if len(userID) == 0 {
		log.Info().Msg("userID cannot be empty")
		s.makeToken("", "", http.StatusBadRequest, "userID cannot be empty")
		return
	}
	dvs := &model.Device{
		DeviceID: deviceID,
		Platform: platform,
		Trust:    1,
		Online:   0,
		UserID:   userID,
	}
	model.CreateDevice(dvs)
	model.SetCurrentDeivce(userID, deviceID, 1, 0)
	return
}

/*
用户ID：userID，
时间戳：timestamp，
状态码：code，
反馈消息：message
*/
func (s registService) makeToken(acc string, userID string, code int, msg string) {
	if testregist {
		return
	}
	ts := time.Now().Unix()
	var rgd int
	if code == http.StatusOK {
		rgd = 1
	} else {
		rgd = 0
		log.Info().Msg(msg)
	}
	//保存用户注册状态
	if len(acc) > 0 {
		model.SetAccRegisted(acc, string(ts), rgd)
	}
	//生成响应数据
	resMap := map[string]interface{}{
		"timestamp": ts,
		"code":      code,
		"message":   msg,
		"userID":    userID,
	}
	tk, err := check.Generate(resMap)
	res := registResponse{
		res: &userproto.RegistRes{
			Account: acc,
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
