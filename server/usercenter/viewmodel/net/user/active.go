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
	netstatus "github.com/znk_fullstack/server/usercenter/viewmodel/net/status"
	userpayload "github.com/znk_fullstack/server/usercenter/viewmodel/payload"
)

//activeUserRes 禁用/激活用户响应
type activeUserRes struct {
	res *userproto.UserUpdateActiveRes
	err error
}

const (
	activeExpired = 60 * 5
	activeFreqExp = 60 * 10
)

//activeUserSrv 禁用/激活用户服务接口
type activeUserSrv struct {
	req   *userproto.UserUpdateActiveReq
	res   chan activeUserRes
	doing map[string]bool
	token *usermiddleware.Token
	pool  *userpayload.WorkerPool
}

//newActiveUserSrv 初始化服务器
func newActiveUserSrv() *activeUserSrv {
	srv := &activeUserSrv{
		res:   make(chan activeUserRes),
		doing: make(map[string]bool),
		token: usermiddleware.NewToken(activeExpired, activeFreqExp),
		pool:  userpayload.NewWorkerPool(100),
	}
	srv.pool.Run()
	return srv
}

//writeActiveReq 写入数据
func (as *activeUserSrv) writeActiveReq(req *userproto.UserUpdateActiveReq) {
	as.pool.WriteHandler(func(j chan userpayload.Job) {
		as.req = req
		j <- as
	})
}

//readActiveRes 读取数据
func (as *activeUserSrv) readActiveRes(ctx context.Context) (*userproto.UserUpdateActiveRes, error) {
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case res := <-as.res:
			return res.res, res.err
		}
	}
}

//handleActiveUser 处理请求
func (as *activeUserSrv) handleActiveUser() {
	req := as.req
	//账号
	acc := req.GetAccount()
	if len(acc) == 0 {
		log.Info().Msg("`account` cannot be empty")
		as.makeActiveUserToken(acc, http.StatusBadRequest, errors.New("`account` cannot be empty"))
		return
	}
	tkstr := req.GetData()
	if len(tkstr) == 0 {
		log.Info().Msg("`data` cannot be empty")
		as.makeActiveUserToken(acc, http.StatusBadRequest, errors.New("`data` cannot be empty"))
		return
	}
	//正在执行
	if as.doing[acc] {
		msg := acc + "is doing request active user, please try again later"
		log.Info().Msg(msg)
		as.makeActiveUserToken(acc, http.StatusBadRequest, errors.New(msg))
		return
	}
	as.doing[acc] = true
	//解析数据
	tk := as.token
	code, err := tk.Parse(acc, "active_user", tkstr)
	if err != nil {
		log.Info().Msg(err.Error())
		as.makeActiveUserToken(acc, code, err)
		return
	}
	code, err = usermiddleware.CommonRequestVerify(acc, tk)
	if err != nil {
		log.Info().Msg(err.Error())
		as.makeActiveUserToken(acc, code, err)
		return
	}
	//查当前用户权限
	var user usermodel.UserDB
	user, err = usermodel.FindUser(acc, tk.UserID)
	if err != nil {
		log.Info().Msg(err.Error())
		as.makeActiveUserToken(acc, netstatus.NoMatchUser, err)
		return
	}
	if user.Per > usermodel.Super {
		msg := acc + "has no permiss to active or inactive user"
		log.Info().Msg(msg)
		as.makeActiveUserToken(acc, netstatus.NoActivePermision, errors.New(msg))
		return
	}
}

func (as *activeUserSrv) makeActiveUserToken(acc string, code int, err error) {
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
	tk, err = as.token.Generate(resmap)
	res := activeUserRes{
		res: &userproto.UserUpdateActiveRes{
			Account: acc,
			Data:    tk,
		},
	}
	as.res <- res
	delete(as.doing, acc)
}

func (as *activeUserSrv) Do() {
	go as.handleActiveUser()
}
