package usernet

import (
	"context"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
	usermodel "github.com/znk_fullstack/server/usercenter/model/user"
	usermiddleware "github.com/znk_fullstack/server/usercenter/viewmodel/middleware"
	userpayload "github.com/znk_fullstack/server/usercenter/viewmodel/payload"
	"google.golang.org/grpc"
)

var rs *registService
var rvt usermiddleware.VerifyToken
var registPool userpayload.WorkerPool

const (
	registExpired = 60 * 2
)

func init() {
	rs = &registService{
		resChan: make(chan registResponse),
		doing:   make(map[string]bool),
	}
	rvt = usermiddleware.NewVerifyToken(registExpired)
	registPool = userpayload.CreateWorkerPool(100)
	registPool.Run()
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
	doing   map[string]bool
}

func (s *registService) Do() {
	go s.handleRegist()
}

//handleRegist 处理注册
func (s *registService) handleRegist() {
	req := s.req
	acc := req.GetAccount()
	if len(acc) == 0 {
		log.Info().Msg("account cannot be empty")
		s.makeRegistToken("", "", http.StatusBadRequest, "account cannot be empty")
		return
	}

	//当前账号正在注册中
	if s.doing[acc] {
		log.Info().Msg("account is registing")
		s.makeRegistToken(acc, "", http.StatusBadRequest, "account is registing")
		return
	}
	s.doing[acc] = true
	//解析校验token
	res, dID, plf, expired, e := rvt.Verify(req.GetToken())
	if !expired { //是否频繁请求
		s.makeRegistToken(acc, "", http.StatusBadRequest, "please regist later on")
		return
	}
	//redis 校验
	exs, oldTS, registed := usermodel.UserRegisted(acc)
	if e != nil {
		log.Info().Msg(e.Error())
		s.makeRegistToken(acc, "", http.StatusBadRequest, e.Error())
		return
	}
	//如果存在redis中，曾调过注册方法
	if exs {
		//如果已注册
		if registed == 1 {
			s.makeRegistToken(acc, "", http.StatusBadRequest, "user has registed:")
			return
		}

		if oldTS < 0 {
			s.makeRegistToken(acc, "", http.StatusBadRequest, "miss param `timestamp`")
			return
		}
		ts := time.Now().Unix()
		if ts-oldTS < int64(registExpired) {
			s.makeRegistToken(acc, "", http.StatusBadRequest, "please regist later on")
			return
		}
	}

	pt, ok := res["photo"].(string)
	if !ok || len(pt) == 0 {
		pt = ""
	}

	psd, ok := res["password"].(string)
	if !ok || len(psd) == 0 {
		log.Info().Msg("password cannot be empty")
		s.makeRegistToken("", "", http.StatusBadRequest, "password cannot be empty")
		return
	}

	userID := makeID()
	e = saveCurrentDevice(userID, dID, plf)
	if e != nil {
		log.Info().Msg(e.Error())
		s.makeRegistToken(acc, "", http.StatusBadRequest, e.Error())
		return
	}
	s.saveUser(acc, pt, userID, psd)
}

//saveUser 保存用户信息
func (s *registService) saveUser(acc string, photo string, userID string, password string) {

	e := usermodel.CreateUser(acc, photo, userID, password)
	if e != nil {
		log.Info().Msg(e.Error())
		s.makeRegistToken(acc, "", http.StatusBadRequest, e.Error())
		return
	}
	s.makeRegistToken(acc, userID, http.StatusOK, "operation success")
	return
}

/*
用户ID：userID，
时间戳：timestamp，
状态码：code，
反馈消息：message
*/

//makeRegistToken 注册token
func (s *registService) makeRegistToken(acc string, userID string, code int, msg string) {
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
		usermodel.SetUserRegisted(acc, string(ts), rgd)
	}
	//生成响应数据
	resmap := map[string]interface{}{
		"timestamp": ts,
		"code":      code,
		"message":   msg,
		"userID":    userID,
	}
	tk, err := rvt.Generate(resmap)
	res := registResponse{
		res: &userproto.RegistRes{
			Account: acc,
			Token:   tk,
		},
		err: err,
	}
	delete(s.doing, acc)
	s.resChan <- res
	return
}

//registerRegistServer 注册到注册请求服务
func registerRegistServer(srv *grpc.Server) {
	userproto.RegisterRegistSrvServer(srv, rs)
}

//UserReigst 注册
func (s *registService) UserReigst(ctx context.Context, req *userproto.RegistReq) (*userproto.RegistRes, error) {
	registPool.WriteHandler(func(jq chan userpayload.Job) {
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
