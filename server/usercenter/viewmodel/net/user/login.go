package usernet

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
	devicemodel "github.com/znk_fullstack/server/usercenter/model/device"
	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
	usermodel "github.com/znk_fullstack/server/usercenter/model/user"
	usermiddleware "github.com/znk_fullstack/server/usercenter/viewmodel/middleware"
	netstatus "github.com/znk_fullstack/server/usercenter/viewmodel/net/status"
	userpayload "github.com/znk_fullstack/server/usercenter/viewmodel/payload"
)

const (
	loginExpired = 60 * 5
	loginFreqExp = 60 * 10
)

type lgnRes struct {
	res *userproto.UserLgnRes
	err error
}

//lgnSrv 登录服务
type lgnSrv struct {
	req     *userproto.UserLgnReq
	resChan chan lgnRes
	doing   map[string]bool
	token   *usermiddleware.Token
	pool    *userpayload.WorkerPool
}

//newRgstSrv 初始化注册服务
func newLgnSrv() *lgnSrv {
	srv := &lgnSrv{
		resChan: make(chan lgnRes),
		doing:   make(map[string]bool),
		token:   usermiddleware.NewToken(loginExpired, loginFreqExp),
		pool:    userpayload.NewWorkerPool(100),
	}
	srv.pool.Run()
	return srv
}

//write 写入数据
func (l *lgnSrv) write(req *userproto.UserLgnReq) {
	l.pool.WriteHandler(func(jq chan userpayload.Job) {
		l.req = req
		jq <- l
	})
}

// 读取数据
func (l *lgnSrv) read(ctx context.Context) (*userproto.UserLgnRes, error) {
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case res := <-l.resChan:
			return res.res, res.err
		}
	}
}

//handleLogin 处理登录请求
func (l *lgnSrv) handleLogin() {
	req := l.req
	acc := req.GetAccount()
	if len(acc) == 0 {
		log.Info().Msg("account cannot be empty")
		l.makeLoginToken("", "", http.StatusBadRequest, errors.New("account cannot be empty"), nil)
		return
	}
	//判断是否有token
	tkstr := req.GetData()
	if len(tkstr) == 0 {
		log.Info().Msg("`data` cannot be empty")
		l.makeLoginToken(acc, "", http.StatusBadRequest, errors.New("`data` cannot be empty"), nil)
		return
	}
	//正在处理登陆操作
	if l.doing[acc] {
		msg := acc + "is doing login, please try again later"
		log.Info().Msg(msg)
		l.makeLoginToken(acc, "", http.StatusBadRequest, errors.New(msg), nil)
		return
	}
	l.doing[acc] = true

	//校验token
	tk := l.token
	code, err := tk.Parse(acc, "login", tkstr)
	if err != nil {
		log.Info().Msg(err.Error())
		l.makeLoginToken(acc, "", code, err, nil)
		return
	}
	//登录条件校验
	code, err = usermiddleware.BaseVerify(acc, tk)
	if err != nil {
		log.Info().Msg(err.Error())
		l.makeLoginToken(acc, "", code, err, nil)
		return
	}
	//是否已注册
	exs, ts, rgd := usermodel.UserRegisted(acc)
	if !exs || rgd == 0 {
		log.Info().Msg("account not registed")
		l.makeLoginToken(acc, "", http.StatusBadRequest, errors.New("account not registed"), nil)
		return
	}
	//请求频繁度检测
	now := time.Now().Unix()
	if now-ts < loginExpired {
		log.Info().Msg("login request too frequence")
		l.makeLoginToken(acc, "", http.StatusBadRequest, errors.New("please try again later"), nil)
		return
	}

	res := tk.Result
	//用户ID检测
	userID, ok := res["userID"].(string)
	if !ok || len(userID) == 0 {
		log.Info().Msg("userID cannot be empty")
		l.makeLoginToken(acc, "", http.StatusBadRequest, errors.New("userID cannot be empty"), nil)
		return
	}

	//查相关设备
	dvcexs := devicemodel.DeviceExists(userID)
	if !dvcexs {
		err := devicemodel.SetCurrentDevice(userID, tk.DeviceID, tk.DeviceName, tk.Platform, 1, false)
		if err != nil {
			log.Info().Msg(err.Error())
			l.makeLoginToken(acc, "", http.StatusInternalServerError, err, nil)
			return
		}
	} else {
		device, err := devicemodel.CurrentDevice(userID)
		if err != nil {
			log.Info().Msg(err.Error())
			l.makeLoginToken(acc, "", http.StatusInternalServerError, err, nil)
			return
		}
		if device.State == devicemodel.Reject {
			log.Info().Msg("device has been reject use")
			l.makeLoginToken(acc, "", netstatus.RejectDevice, errors.New("device has been reject use"), nil)
			return
		}
	}

	//查用户数据
	user, err := usermodel.FindUser(acc, userID)
	if err != nil || user == nil {
		log.Info().Msg("user not exists")
		l.makeLoginToken(acc, "", http.StatusBadRequest, err, nil)
		return
	}
	usermodel.UserOnline(acc, userID)
	l.makeLoginToken(acc, tk.DeviceID, http.StatusOK, nil, user)
	return
}

/*
状态码：code，
反馈消息：message，
时间戳：timestamp，
用户信息：user
*/

//makeLoginToken 登录token
func (l *lgnSrv) makeLoginToken(acc, deviceID string, code int, err error, user *userproto.User) {
	var sess string
	//无错，用户数据不为空才生成session
	if err == nil && user != nil {
		sess, err = usermiddleware.DefaultSess.SessionID(user.UserID, deviceID)
		if err != nil || len(sess) == 0 {
			err = errors.New("internal server error")
		}
	}
	resmap := map[string]interface{}{
		"code":      code,
		"message":   err.Error(),
		"user":      user,
		"sessionID": sess,
	}
	var tk string
	tk, err = l.token.Generate(resmap)
	res := lgnRes{
		err: err,
		res: &userproto.UserLgnRes{
			Account: acc,
			Data:    tk,
		},
	}
	//删除正在操作状态
	delete(l.doing, acc)
	l.resChan <- res
	return
}

//Do 执行任务
func (l *lgnSrv) Do() {
	go l.handleLogin()
}
