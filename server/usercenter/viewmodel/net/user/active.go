package usernet

import (
	"context"
	"errors"
	"net/http"
	"strconv"
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

/*
用户ID：userID，
会话ID：sessionID ，
时间戳：timestamp，
设备ID：deviceID，
设备名：deviceName，
平台类型：platform，
应用标识：appkey，
目标用户账号：targetAcc，
目标用户ID：targetID
*/

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
		msg := acc + "- `data` cannot be empty"
		log.Info().Msg(msg)
		as.makeActiveUserToken(acc, http.StatusBadRequest, errors.New(msg))
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
		msg := acc + "- active user error: " + err.Error()
		log.Info().Msg(msg)
		as.makeActiveUserToken(acc, code, err)
		return
	}
	code, err = usermiddleware.CommonRequestVerify(acc, tk)
	if err != nil {
		msg := acc + "- active user error: " + err.Error()
		log.Info().Msg(msg)
		as.makeActiveUserToken(acc, code, err)
		return
	}
	//查当前用户权限
	var user usermodel.UserDB
	user, err = usermodel.FindUser(acc, tk.UserID)
	if err != nil {
		msg := acc + "- active user error: " + err.Error()
		log.Info().Msg(msg)
		as.makeActiveUserToken(acc, netstatus.NoMatchUser, err)
		return
	}
	if user.Per > usermodel.Super {
		msg := acc + " has no permiss to active or inactive user"
		log.Info().Msg(msg)
		as.makeActiveUserToken(acc, netstatus.NoActivePermision, errors.New(msg))
		return
	}
	//目标账号
	res := tk.Result
	targetAcc, ok := res["targetAcc"].(string)
	if !ok || len(targetAcc) == 0 {
		msg := acc + "- `targetAcc` cannot be empty"
		log.Info().Msg(msg)
		as.makeActiveUserToken(acc, http.StatusBadRequest, errors.New(msg))
		return
	}
	targetID, ok := res["targetID"].(string)
	if !ok || len(targetID) == 0 {
		msg := acc + "- `targetID` cannot be empty"
		log.Info().Msg(msg)
		as.makeActiveUserToken(acc, http.StatusBadRequest, errors.New(msg))
		return
	}
	atstr, ok := res["active"].(string)
	if !ok || len(atstr) == 0 {
		msg := acc + "- `active` cannot be empty"
		log.Info().Msg(msg)
		as.makeActiveUserToken(acc, http.StatusBadRequest, errors.New(msg))
		return
	}
	var active int
	active, err = strconv.Atoi(atstr)
	if err != nil {
		msg := acc + "- `active` is error type"
		log.Info().Msg(msg)
		as.makeActiveUserToken(acc, http.StatusBadRequest, errors.New(msg))
		return
	}
	err = usermodel.SetUserActive(targetAcc, targetID, active)
	if err != nil {
		msg := acc + "internal server error: " + err.Error()
		log.Info().Msg(msg)
		as.makeActiveUserToken(acc, http.StatusInternalServerError, err)
		return
	}
}

//makeActiveUserToken 生成激活/禁用用户响应数据
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
