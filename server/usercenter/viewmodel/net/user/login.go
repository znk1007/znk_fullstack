package usernet

import (
	"errors"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
	devicemodel "github.com/znk_fullstack/server/usercenter/model/device"
	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
	usermodel "github.com/znk_fullstack/server/usercenter/model/user"
	usermiddleware "github.com/znk_fullstack/server/usercenter/viewmodel/middleware"
	userpayload "github.com/znk_fullstack/server/usercenter/viewmodel/payload"
)

var lvt usermiddleware.VerifyToken
var loginPool userpayload.WorkerPool

const (
	loginExpired = 60 * 5
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
}

//newRgstSrv 初始化注册服务
func newLgnSrv() *lgnSrv {
	lvt = usermiddleware.NewVerifyToken(loginExpired)
	loginPool = userpayload.CreateWorkerPool(100)
	loginPool.Run()
	return &lgnSrv{
		resChan: make(chan lgnRes),
		doing:   make(map[string]bool),
	}
}

//write 写入数据
func (l *lgnSrv) write(req *userproto.UserLgnReq) {
	loginPool.WriteHandler(func(jq chan userpayload.Job) {
		l.req = req
		jq <- l
	})
}

// 读取数据
func (l *lgnSrv) read() (*userproto.UserLgnRes, error) {
	for {
		select {
		case res := <-l.resChan:
			return res.res, res.err
		}
	}
}

//handleLogin 处理登录请求
func (l *lgnSrv) handleLogin() {
	acc := l.req.GetAccount()
	if len(acc) == 0 {
		log.Info().Msg("account cannot be empty")
		l.makeLoginToken("", http.StatusBadRequest, errors.New("account cannot be empty"), nil)
		return
	}
	//正在处理登陆操作
	if l.doing[acc] {
		return
	}
	l.doing[acc] = true

	//校验token
	e := lvt.Verify(l.req.GetToken())
	if e != nil {
		log.Info().Msg(e.Error())
		l.makeLoginToken(acc, http.StatusBadRequest, e, nil)
		return
	}
	//超时检测
	if !lvt.Expired {
		log.Info().Msg("login request too frequence")
		l.makeLoginToken(acc, http.StatusBadRequest, errors.New("please try again later"), nil)
		return
	}
	//是否已注册
	exs, ts, rgd := usermodel.UserRegisted(acc)
	if !exs || rgd == 0 {
		log.Info().Msg("account not registed")
		l.makeLoginToken(acc, http.StatusBadRequest, errors.New("account not registed"), nil)
		return
	}
	//请求频繁度检测
	now := time.Now().Unix()
	if now-ts < loginExpired {
		log.Info().Msg("login request too frequence")
		l.makeLoginToken(acc, http.StatusBadRequest, errors.New("please try again later"), nil)
		return
	}
	res := lvt.Result
	//用户ID检测
	userID, ok := res["userID"].(string)
	if !ok || len(userID) == 0 {
		log.Info().Msg("userID cannot be empty")
		l.makeLoginToken(acc, http.StatusBadRequest, errors.New("userID cannot be empty"), nil)
		return
	}

	//查相关设备
	dvcexs := devicemodel.DeviceExists(userID)
	if !dvcexs {
		err := devicemodel.SetCurrentDevice(userID, lvt.DeviceID, lvt.DeviceName, lvt.Platform, 1)
		if err != nil {
			log.Info().Msg(err.Error())
			l.makeLoginToken(acc, http.StatusInternalServerError, err, nil)
			return
		}
	} else {
		device, err := devicemodel.CurrentDevice(userID)
		if err != nil {
			log.Info().Msg(err.Error())
			l.makeLoginToken(acc, http.StatusInternalServerError, err, nil)
			return
		}
		if device.State == devicemodel.Reject {
			log.Info().Msg("device has been reject use")
			l.makeLoginToken(acc, 1001, errors.New("device has been reject use"), nil)
			return
		}
	}

	//查用户数据
	user, err := usermodel.FindUser(acc, userID)
	if err != nil || user == nil {
		log.Info().Msg("user not exists")
		l.makeLoginToken(acc, http.StatusBadRequest, err, nil)
		return
	}

	// phone, email, nickname, photo, err := usermodel.FindUser(acc, userID)
	// if err != nil {
	// 	log.Info().Msg(err.Error())

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
func (l *lgnSrv) makeLoginToken(acc string, code int, err error, user *userproto.User) {
	ts := time.Now().Unix()
	sess, e := usermiddleware.SessionID(user.UserID)
	if e != nil || len(sess) == 0 {
		err = errors.New("internal server error")
	}
	resmap := map[string]interface{}{
		"code":      code,
		"message":   err.Error(),
		"timestamp": string(ts),
		"user":      user,
		"sessionID": sess,
	}
	tk, err := lvt.Generate(resmap)
	res := lgnRes{
		err: err,
		res: &userproto.UserLgnRes{
			Account: acc,
			Token:   tk,
		},
	}
	//删除正在操作状态
	delete(l.doing, acc)
	l.resChan <- res
}

//Do 执行任务
func (l *lgnSrv) Do() {
	go l.handleLogin()
}
