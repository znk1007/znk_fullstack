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

//write 写入数据
func (ls *lgoSrv) write(req *userproto.UserLgoReq) {
	ls.pool.WriteHandler(func(j chan userpayload.Job) {
		ls.req = req
		j <- ls
	})
}

//read 读取数据
func (ls *lgoSrv) read(ctx context.Context) (*userproto.UserLgoRes, error) {
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case res := <-ls.resChan:
			return res.res, res.err
		}
	}
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
	tkstr := req.GetToken()
	if len(tkstr) == 0 {
		log.Info().Msg("token cannot be empty")
		ls.makeLogoutToken(acc, http.StatusBadRequest, errors.New("token cannot be empty"))
		return
	}
	if ls.doing[acc] {
		log.Info().Msg("please request later")
		ls.makeLogoutToken(acc, http.StatusBadRequest, errors.New("please request later"))
		return
	}
	err := ls.token.Parse(tkstr)
	if err != nil {
		log.Info().Msg(err.Error())
		ls.makeLogoutToken(acc, http.StatusBadRequest, err)
		return
	}
	res := ls.token.Result
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
