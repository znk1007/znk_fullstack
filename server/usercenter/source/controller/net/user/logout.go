package usernet

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
	usermiddleware "github.com/znk_fullstack/server/usercenter/source/controller/middleware"
	netstatus "github.com/znk_fullstack/server/usercenter/source/controller/net/status"
	userpayload "github.com/znk_fullstack/server/usercenter/source/controller/payload"
	userproto "github.com/znk_fullstack/server/usercenter/source/model/protos/generated"
	usermodel "github.com/znk_fullstack/server/usercenter/source/model/user"
)

const (
	logoutExpired = 60 * 5
	logoutExp     = 60 * 10
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
	pool    *userpayload.WorkerPool
}

//newLogSrv 初始化
func newLogSrv() *lgoSrv {
	srv := &lgoSrv{
		resChan: make(chan lgoRes),
		doing:   make(map[string]bool),
		pool:    userpayload.NewWorkerPool(100),
		token:   usermiddleware.NewToken(logoutExpired, logoutExp),
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
	tkstr := req.GetData()
	if len(tkstr) == 0 {
		log.Info().Msg("`data` cannot be empty")
		ls.makeLogoutToken(acc, http.StatusBadRequest, errors.New("`data` cannot be empty"))
		return
	}
	if ls.doing[acc] {
		msg := acc + "is doing logout, please try again later"
		log.Info().Msg(msg)
		ls.makeLogoutToken(acc, http.StatusBadRequest, errors.New(msg))
		return
	}
	tk := ls.token
	code, err := tk.Parse(acc, "logout", tkstr)
	if err != nil {
		msg := acc + " - internal server error: " + err.Error()
		log.Info().Msg(msg)
		ls.makeLogoutToken(acc, code, err)
		return
	}
	code, err = usermiddleware.CommonRequestVerify(acc, tk)
	if code == netstatus.UserLogout {
		ls.makeLogoutToken(acc, http.StatusOK, nil)
		return
	}
	if err != nil {
		msg := acc + " - internal server error: " + err.Error()
		log.Info().Msg(msg)
		ls.makeLogoutToken(acc, code, err)
		return
	}
	online, err := usermodel.UserOnline(acc, tk.UserID)
	if err != nil {
		msg := acc + " - internal server error: " + err.Error()
		log.Info().Msg(msg)
		ls.makeLogoutToken(acc, http.StatusInternalServerError, err)
		return
	}
	if online == 0 {
		ls.makeLogoutToken(acc, http.StatusOK, nil)
		return
	}
	err = usermodel.SetUserOnline(acc, tk.UserID, 0)
	if err != nil {
		msg := acc + " - internal server error: " + err.Error()
		log.Info().Msg(msg)
		ls.makeLogoutToken(acc, http.StatusInternalServerError, err)
		return
	}
	ls.makeLogoutToken(acc, http.StatusOK, nil)
}

/*
状态码：code，
反馈消息：message，
时间戳：timestamp
*/
func (ls *lgoSrv) makeLogoutToken(acc string, code int, err error) {
	msg := "operation success"
	if err != nil {
		msg = err.Error()
	}
	resmap := map[string]interface{}{
		"code":      code,
		"message":   msg,
		"timestamp": time.Now().Unix(),
	}
	var tk string
	tk, err = ls.token.Generate(resmap)
	res := lgoRes{
		err: err,
		res: &userproto.UserLgoRes{
			Account: acc,
			Data:    tk,
		},
	}
	ls.resChan <- res
	delete(ls.doing, acc)
	return
}

func (ls *lgoSrv) Do() {
	go ls.handleLogout()
}