package usernet

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
	usermodel "github.com/znk_fullstack/server/usercenter/model/user"
	usermiddleware "github.com/znk_fullstack/server/usercenter/viewmodel/middleware"
	userpayload "github.com/znk_fullstack/server/usercenter/viewmodel/payload"
)

var rs *rgstSrv

const (
	registExpired = 60 * 2
)

//rgstRes 注册响应
type rgstRes struct {
	res *userproto.UserRgstRes
	err error
}

//rgstSrv 注册
type rgstSrv struct {
	req     *userproto.UserRgstReq
	resChan chan rgstRes
	doing   map[string]bool
	token   *usermiddleware.Token
	pool    userpayload.WorkerPool
}

//newRgstSrv 初始化注册服务
func newRgstSrv() *rgstSrv {
	srv := &rgstSrv{
		resChan: make(chan rgstRes),
		doing:   make(map[string]bool),
		token:   usermiddleware.NewToken(registExpired),
		pool:    userpayload.NewWorkerPool(100),
	}
	srv.pool.Run()
	return srv
}

//write 写入数据
func (s *rgstSrv) write(req *userproto.UserRgstReq) {
	s.pool.WriteHandler(func(jq chan userpayload.Job) {
		s.req = req
		jq <- s
	})
}

// 读取数据
func (s *rgstSrv) read(ctx context.Context) (*userproto.UserRgstRes, error) {
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case res := <-s.resChan:
			return res.res, res.err
		}
	}
}

func (s *rgstSrv) Do() {
	go s.handleRegist()
}

//handleRegist 处理注册
func (s *rgstSrv) handleRegist() {
	req := s.req
	acc := req.GetAccount()
	if len(acc) == 0 {
		log.Info().Msg("account cannot be empty")
		s.makeRegistToken("", "", http.StatusBadRequest, errors.New("account cannot be empty"))
		return
	}
	//判断是否有token
	tkstr := req.GetToken()
	if len(tkstr) == 0 {
		log.Info().Msg("token cannot be empty")
		s.makeRegistToken(acc, "", http.StatusBadRequest, errors.New("token cannot be empty"))
		return
	}

	//当前账号正在注册中
	if s.doing[acc] {
		log.Info().Msg("account is operating, please do it later")
		s.makeRegistToken(acc, "", http.StatusBadRequest, errors.New("account is operating, please do it later"))
		return
	}
	s.doing[acc] = true

	//解析校验token
	e := s.token.Parse(req.GetToken())
	if e != nil {
		log.Info().Msg(e.Error())
		s.makeRegistToken(acc, "", http.StatusBadRequest, e)
		return
	}

	if !s.token.Expired { //是否频繁请求
		log.Info().Msg("request too frequence")
		s.makeRegistToken(acc, "", http.StatusBadRequest, errors.New("please regist later on"))
		return
	}
	//redis 校验
	exs, oldTS, registed := usermodel.UserRegisted(acc)
	if e != nil {
		log.Info().Msg(e.Error())
		s.makeRegistToken(acc, "", http.StatusBadRequest, e)
		return
	}
	//如果存在redis中，曾调过注册方法
	if exs {
		//如果已注册
		if registed == 1 {
			s.makeRegistToken(acc, "", http.StatusBadRequest, errors.New("user has registed"))
			return
		}

		if oldTS < 0 {
			s.makeRegistToken(acc, "", http.StatusBadRequest, errors.New("miss param `timestamp`"))
			return
		}
		ts := time.Now().Unix()
		if ts-oldTS < int64(registExpired) {
			s.makeRegistToken(acc, "", http.StatusBadRequest, errors.New("please regist later on"))
			return
		}
	}
	res := s.token.Result
	pt, ok := res["photo"].(string)
	if !ok || len(pt) == 0 {
		pt = ""
	}

	psd, ok := res["password"].(string)
	if !ok || len(psd) == 0 {
		log.Info().Msg("password cannot be empty")
		s.makeRegistToken("", "", http.StatusBadRequest, errors.New("password cannot be empty"))
		return
	}
	userID := makeID()
	s.saveUser(acc, pt, userID, psd)
}

//saveUser 保存用户信息
func (s *rgstSrv) saveUser(acc string, photo string, userID string, password string) {

	e := usermodel.CreateUser(acc, photo, userID, password)
	if e != nil {
		log.Info().Msg(e.Error())
		s.makeRegistToken(acc, "", http.StatusBadRequest, e)
		return
	}
	s.makeRegistToken(acc, userID, http.StatusOK, errors.New("operation success"))
	return
}

/*
用户ID：userID，
时间戳：timestamp，
状态码：code，
反馈消息：message
*/

//makeRegistToken 注册token
func (s *rgstSrv) makeRegistToken(acc string, userID string, code int, err error) {
	var rgd int
	if code == http.StatusOK {
		rgd = 1
	} else {
		rgd = 0
		log.Info().Msg(err.Error())
	}
	//保存用户注册状态
	if len(acc) > 0 {
		ts := time.Now().Unix()
		usermodel.SetUserRegisted(acc, string(ts), rgd)
	}
	//生成响应数据
	resmap := map[string]interface{}{
		"code":    code,
		"message": err.Error(),
		"userID":  userID,
	}
	var tk string
	tk, err = s.token.Generate(resmap)
	res := rgstRes{
		res: &userproto.UserRgstRes{
			Account: acc,
			Token:   tk,
		},
		err: err,
	}
	//删除正在操作状态
	delete(s.doing, acc)
	s.resChan <- res
	return
}
