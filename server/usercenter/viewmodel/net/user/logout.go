package usernet

import (
	"errors"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
	usermiddleware "github.com/znk_fullstack/server/usercenter/viewmodel/middleware"
	userpayload "github.com/znk_fullstack/server/usercenter/viewmodel/payload"
)

type lgoRes struct {
	res *userproto.UserLgoRes
	err error
}

//lgoSrv 退出登录
type lgoSrv struct {
	req     *userproto.UserLgoReq
	resChan chan lgoRes
	doing   map[string]bool
	token   *usermiddleware.Token
	pool    userpayload.WorkerPool
}

//newLogSrv 初始化
func newLogSrv() *lgoSrv {
	srv := &lgoSrv{
		resChan: make(chan lgoRes),
		doing:   make(map[string]bool),
		pool:    userpayload.NewWorkerPool(100),
	}
	srv.pool.Run()
	return srv
}

//handleLogout 处理退出登录
func (ls *lgoSrv) handleLogout() {
	req := ls.req
	acc := req.GetAccount()
	if len(acc) == 0 {
		log.Info().Msg("account cannot be empty")
		ls.makeLogoutToken(acc, http.StatusBadRequest, errors.New("account cannot be empty"))
		return
	}
}

/*
状态码：code，
反馈消息：message，
时间戳：timestamp
*/
func (ls *lgoSrv) makeLogoutToken(acc string, code int, err error) {
	resmap := map[string]interface{}{
		"code":      code,
		"message":   err.Error(),
		"timestamp": time.Now().String(),
	}
	var tk string
	tk, err = ls.token.Generate(resmap)
	res := lgoRes{
		err: err,
		res: &userproto.UserLgoRes{
			Account: acc,
			Token:   tk,
		},
	}
	delete(ls.doing, acc)
	ls.resChan <- res
	return
}

func (ls *lgoSrv) Do() {
	go ls.handleLogout()
}
