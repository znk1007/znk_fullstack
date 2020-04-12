package usernet

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
	usermiddleware "github.com/znk_fullstack/server/usercenter/viewmodel/middleware"
	userpayload "github.com/znk_fullstack/server/usercenter/viewmodel/payload"
)

var ls *loginService
var lvt usermiddleware.VerifyToken
var loginPool userpayload.WorkerPool

const (
	loginExpired = 60 * 5
)

type loginResponse struct {
	res *userproto.LoginRes
	err error
}

//loginService 登录服务
type loginService struct {
	req     *userproto.LoginReq
	resChan chan loginResponse
	doing   map[string]bool
}

func init() {
	ls = &loginService{
		resChan: make(chan loginResponse),
		doing:   make(map[string]bool),
	}
	lvt = usermiddleware.NewVerifyToken(loginExpired)
	loginPool = userpayload.CreateWorkerPool(100)
	loginPool.Run()
}

//handleLogin 处理登录请求
func (l *loginService) handleLogin() {
	acc := l.req.GetAccount()
	if len(acc) == 0 {
		log.Info().Msg("account cannot be empty")
		l.makeLoginToken("", http.StatusBadRequest, errors.New("account cannot be empty"), nil)
		return
	}
	// res, dID, plf, exp, e := lvt.Verify(l.req.GetToken())
	// if e != nil {
	// 	log.Info().Msg(e.Error())
	// 	l.makeLoginToken(acc, http.StatusBadRequest, e, nil)
	// 	return
	// }

}

/*
状态码：code，
反馈消息：message，
时间戳：timestamp，
用户信息：user
*/

//makeLoginToken 登录token
func (l *loginService) makeLoginToken(acc string, code int, err error, user *userproto.User) {
	ts := time.Now().Unix()
	resmap := map[string]interface{}{
		"code":      code,
		"message":   err.Error(),
		"timestamp": string(ts),
		"user":      user,
	}
	tk, err := lvt.Generate(resmap)
	l.resChan <- loginResponse{
		err: err,
		res: &userproto.LoginRes{
			Account: acc,
			Token:   tk,
		},
	}
}

//Do 执行任务
func (l *loginService) Do() {
	go l.handleLogin()
}

//UserLogin 登录接口
func (l *loginService) UserLogin(ctx context.Context, req *userproto.LoginReq) (*userproto.LoginRes, error) {
	loginPool.WriteHandler(func(j chan userpayload.Job) {
		l.req = req
		j <- l
	})
	for {
		select {
		case res := <-l.resChan:
			return res.res, res.err
		}
	}
}
