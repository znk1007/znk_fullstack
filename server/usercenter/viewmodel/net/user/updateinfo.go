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

type reqType int

const (
	//更新密码
	psw      reqType = iota
	phone    reqType = 1
	nickname reqType = 2
)

const (
	updateInfoExpired = 60 * 2
)

//updatePswRes 更新密码响应
type updatePswRes struct {
	res *userproto.UserUpdatePswRes
	err error
}

//updatePhoneRes 更新手机响应
type updatePhoneRes struct {
	res *userproto.UserUpdatePhoneRes
	err error
}

//updateNicknameRes 更新昵称响应
type updateNicknameRes struct {
	res *userproto.UserUpdateNicknameRes
	err error
}

//updateInfoSrv 更新信息服务
type updateInfoSrv struct {
	phoneReq    *userproto.UserUpdatePhoneReq
	phoneRes    chan updatePhoneRes
	nicknameReq *userproto.UserUpdateNicknameReq
	nicknameRes chan updateNicknameRes
	pswReq      *userproto.UserUpdatePswReq
	pswRes      chan updatePswRes
	doing       map[string]bool
	token       *usermiddleware.Token
	pool        userpayload.WorkerPool
	rt          reqType
}

//newUpdateInfoSrv 初始化服务器
func newUpdateInfoSrv() *updateInfoSrv {
	srv := &updateInfoSrv{
		phoneRes:    make(chan updatePhoneRes),
		nicknameRes: make(chan updateNicknameRes),
		pswRes:      make(chan updatePswRes),
		doing:       make(map[string]bool),
		token:       usermiddleware.NewToken(300),
	}
	srv.pool.Run()
	return srv
}

//handleUpdateInfo 处理更新信息请求
func (ui *updateInfoSrv) handleUpdateInfo() {
	switch ui.rt {
	case psw:
		ui.handlUpdatePsw()
	}
}

//writePhoneReq 读入更新手机号数据
func (ui *updateInfoSrv) writePhoneReq(req *userproto.UserUpdatePhoneReq) {
	ui.rt = phone
	ui.pool.WriteHandler(func(j chan userpayload.Job) {
		ui.phoneReq = req
		j <- ui
	})
}

//------------------------------------update password---------------------------------------

//writePswReq 写入更新密码数据
func (ui *updateInfoSrv) writePswReq(req *userproto.UserUpdatePswReq) {
	ui.rt = psw
	ui.pool.WriteHandler(func(j chan userpayload.Job) {
		ui.pswReq = req
		j <- ui
	})
}

//readPswRes 读取更新密码响应数据
func (ui *updateInfoSrv) readPswRes(ctx context.Context) (res *userproto.UserUpdatePswRes, err error) {
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case res := <-ui.pswRes:
			return res.res, res.err
		}
	}
}

/*
用户ID：userID，
会话ID：sessionID，
时间戳：timestamp，
设备ID：deviceID，
设备名：deviceName，
平台类型：platform，
应用标识：appkey，
旧密码：password，
新密码：newPsw
*/

//handlUpdatePsw 处理更新密码
func (ui *updateInfoSrv) handlUpdatePsw() {
	req := ui.pswReq
	//账号
	acc := req.GetAccount()
	//校验账号是否为空
	if len(acc) == 0 {
		log.Info().Msg("miss `account` or account cannot be empty")
		ui.makeUpdatePswToken("", http.StatusBadRequest, errors.New("miss `account` or account cannot be empty"))
		return
	}
	//token
	tkstr := req.GetToken()
	if len(tkstr) == 0 {
		log.Info().Msg("token cannot be empty")
		ui.makeUpdatePswToken(acc, http.StatusBadRequest, errors.New("token cannot be empty"))
		return
	}
	//正在执行中
	if ui.doing[acc] {
		log.Info().Msg("account is operating, please do it later")
		ui.makeUpdatePswToken(acc, http.StatusBadRequest, errors.New("account is operating, please do it later"))
		return
	}
	//解析token
	tk := ui.token
	err := tk.Parse(tkstr)
	if err != nil {
		log.Info().Msg(err.Error())
		ui.makeUpdatePswToken(acc, http.StatusBadRequest, err)
		return
	}
	//通用请求校验
	code, err := usermiddleware.CommonRequestVerify(acc, tk)
	if err != nil {
		log.Info().Msg(err.Error())
		ui.makeUpdatePswToken(acc, code, err)
		return
	}
	//更新密码
	err = usermodel.SetUserPassword(acc, tk.UserID, tk.Password)
	if err != nil {
		log.Info().Msg(err.Error())
		ui.makeUpdatePswToken(acc, http.StatusBadRequest, err)
		return
	}
	ui.makeUpdatePswToken(acc, http.StatusOK, nil)
}

/*
状态码：code，
反馈消息：message，
时间戳：timestamp
*/
//makeUpdatePswToken 生成更新密码响应token
func (ui *updateInfoSrv) makeUpdatePswToken(acc string, code int, err error) {
	msg := ""
	if err != nil {
		msg = err.Error()
	}
	resmap := map[string]interface{}{
		"code":      code,
		"message":   msg,
		"timestamp": time.Now().String(),
	}
	var tk string
	tk, err = ui.token.Generate(resmap)

	res := updatePswRes{
		res: &userproto.UserUpdatePswRes{
			Account: acc,
			Token:   tk,
		},
		err: err,
	}
	ui.pswRes <- res
}

func (ui *updateInfoSrv) Do() {
	go ui.handleUpdateInfo()
}
