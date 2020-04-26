package usernet

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
	usermiddleware "github.com/znk_fullstack/server/usercenter/viewmodel/middleware"
	netstatus "github.com/znk_fullstack/server/usercenter/viewmodel/net/status"
	userpayload "github.com/znk_fullstack/server/usercenter/viewmodel/payload"
)

const (
	statusExpired = 60 * 5
	statusFreqExp = 60 * 10
)

type userStatusRes struct {
	res *userproto.UserStatusRes
	err error
}

//userStatusSrv 用户状态
type userStatusSrv struct {
	req   *userproto.UserStatusReq
	res   chan userStatusRes
	doing map[string]bool
	token *usermiddleware.Token
	pool  userpayload.WorkerPool
}

//newStatusSrv 初始化服务器
func newStatusSrv() *userStatusSrv {
	srv := &userStatusSrv{
		res:   make(chan userStatusRes),
		doing: make(map[string]bool),
		token: usermiddleware.NewToken(statusExpired, statusFreqExp),
	}
	srv.pool.Run()
	return srv
}

//write 写入数据
func (ss *userStatusSrv) write(req *userproto.UserStatusReq) {
	ss.pool.WriteHandler(func(j chan userpayload.Job) {
		ss.req = req
		j <- ss
	})
}

//read 读取数据
func (ss *userStatusSrv) read(ctx context.Context) (*userproto.UserStatusRes, error) {
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case res := <-ss.res:
			return res.res, nil
		}
	}
}

//handleUserStatus 处理请求
func (ss *userStatusSrv) handleUserStatus() {
	req := ss.req
	acc := req.GetAccount()
	if len(acc) == 0 {
		log.Info().Msg("account cannot be empty")
		ss.makeStatusToken(acc, http.StatusBadRequest, 0, 0, errors.New("account cannot be empty"))
		return
	}
	tkstr := req.GetData()
	if len(tkstr) == 0 {
		log.Info().Msg("`data` cannot be empty")
		ss.makeStatusToken(acc, http.StatusBadRequest, 0, 0, errors.New("`data` cannot be empty"))
		return
	}
	tk := ss.token
	code, err := tk.Parse(acc, "user_status", tkstr)
	if err != nil {
		log.Info().Msg(err.Error())
		ss.makeStatusToken(acc, code, 0, 0, err)
		return
	}
	code, err = usermiddleware.CommonRequestVerify(acc, tk)
	//用户是否存在
	if code == netstatus.UserInactive || code == netstatus.UserNotRegisted {
		log.Info().Msg(err.Error())
		ss.makeStatusToken(acc, code, 0, 0, err)
		return
	}
	if code == netstatus.SessionInvalidate || code == netstatus.UserLogout {
		log.Info().Msg(err.Error())
		ss.makeStatusToken(acc, code, 0, 1, err)
		return
	}
	ss.makeStatusToken(acc, http.StatusOK, 1, 1, nil)
}

/*
参数-
状态码：code，
反馈消息：message，
时间戳：timestamp，
会话ID状态 status，
是否被禁用 active
*/
//makeStatusToken 生成响应token
func (ss *userStatusSrv) makeStatusToken(acc string, code, online, active int, err error) {
	msg := "operation success"
	if err != nil {
		msg = err.Error()
	}
	resmap := map[string]interface{}{
		"code":      code,
		"message":   msg,
		"timestamp": time.Now().Unix(),
		"online":    online,
		"active":    active,
	}
	tk, err := ss.token.Generate(resmap)
	res := userStatusRes{
		res: &userproto.UserStatusRes{
			Account: acc,
			Data:    tk,
		},
	}
	ss.res <- res
	delete(ss.doing, acc)
}

func (ss *userStatusSrv) Do() {
	go ss.handleUserStatus()
}
