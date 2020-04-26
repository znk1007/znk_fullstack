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
	active   reqType = 3
)

const (
	updateInfoExpired = 60 * 5
	updateInfoFreqExp = 60 * 10
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
		token:       usermiddleware.NewToken(updateInfoExpired, updateInfoFreqExp),
	}
	srv.pool.Run()
	return srv
}

//handleUpdateInfo 处理更新信息请求
func (ui *updateInfoSrv) handleUpdateInfo() {
	switch ui.rt {
	case psw:
		ui.handleUpdatePsw()
	case phone:
		ui.handleUpdatePhone()
	case nickname:
		ui.handleUpdateNickname()
	}
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

//handleUpdatePsw 处理更新密码
func (ui *updateInfoSrv) handleUpdatePsw() {
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
	tkstr := req.GetData()
	if len(tkstr) == 0 {
		log.Info().Msg("`data` cannot be empty")
		ui.makeUpdatePswToken(acc, http.StatusBadRequest, errors.New("`data` cannot be empty"))
		return
	}
	//正在执行中
	if ui.doing[acc] {
		msg := acc + "is doing update password, please try again later"
		log.Info().Msg(msg)
		ui.makeUpdatePswToken(acc, http.StatusBadRequest, errors.New(msg))
		return
	}
	ui.doing[acc] = true
	//解析token
	tk := ui.token
	code, err := tk.Parse(acc, "update_password", tkstr)
	if err != nil {
		log.Info().Msg(err.Error())
		ui.makeUpdatePswToken(acc, code, err)
		return
	}
	//通用请求校验
	code, err = usermiddleware.CommonRequestVerify(acc, tk)
	if err != nil {
		log.Info().Msg(err.Error())
		ui.makeUpdatePswToken(acc, code, err)
		return
	}
	//更新密码
	newPsw, ok := tk.Result["newPsw"].(string)
	if !ok || len(newPsw) == 0 {
		log.Info().Msg("`newPsw` cannot be empty")
		ui.makeUpdatePswToken(acc, http.StatusBadRequest, errors.New("`newPsw` cannot be empty"))
		return
	}
	err = usermodel.SetUserPassword(acc, tk.UserID, newPsw)
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
	tk, err = ui.token.Generate(resmap)

	res := updatePswRes{
		res: &userproto.UserUpdatePswRes{
			Account: acc,
			Data:    tk,
		},
		err: err,
	}
	ui.pswRes <- res
	delete(ui.doing, acc)
}

//----------------------update phone-------------------

//writePhoneReq 写入更新手机号数据
func (ui *updateInfoSrv) writePhoneReq(req *userproto.UserUpdatePhoneReq) {
	ui.rt = phone
	ui.pool.WriteHandler(func(j chan userpayload.Job) {
		ui.phoneReq = req
		j <- ui
	})
}

//readPhoneRes 读取更新手机号响应数据
func (ui *updateInfoSrv) readPhoneRes(ctx context.Context) (*userproto.UserUpdatePhoneRes, error) {
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case res := <-ui.phoneRes:
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
手机号：phone
*/

//handleUpdatePhone 处理更新手机号请求
func (ui *updateInfoSrv) handleUpdatePhone() {
	req := ui.phoneReq
	//账号检测
	acc := req.GetAccount()
	if len(acc) == 0 {
		log.Info().Msg("`account` cannot be empty")
		ui.makeUpdatePhoneToken("", http.StatusBadRequest, errors.New("`account` cannot be empty"))
		return
	}
	//token校验
	tkstr := req.GetData()
	if len(tkstr) == 0 {
		log.Info().Msg("`data` cannot be empty")
		ui.makeUpdatePhoneToken(acc, http.StatusBadRequest, errors.New("`data` cannot be empty"))
		return
	}
	code, err := ui.token.Parse(req.GetAccount(), "update_phone", req.GetData())
	if err != nil {
		log.Info().Msg(err.Error())
		ui.makeUpdatePhoneToken(acc, code, err)
		return
	}
	tk := ui.token
	code, err = usermiddleware.CommonRequestVerify(acc, tk)
	if err != nil {
		log.Info().Msg(err.Error())
		ui.makeUpdatePhoneToken(acc, code, err)
		return
	}
	//正在执行中
	if ui.doing[acc] {
		log.Info().Msg("account is operating, please do it later")
		ui.makeUpdatePswToken(acc, http.StatusBadRequest, errors.New("account is operating, please do it later"))
		return
	}
	ui.doing[acc] = true
	//检验phone字段
	p, ok := tk.Result["phone"].(string)
	if !ok || len(p) == 0 {
		log.Info().Msg("`phone` cannot be empty")
		ui.makeUpdatePhoneToken(acc, http.StatusBadRequest, errors.New("`phone` cannot be empty"))
		return
	}
	//更新手机号
	err = usermodel.SetUserPhone(acc, tk.UserID, p)
	if err != nil {
		log.Info().Msg("internal server error")
		ui.makeUpdatePhoneToken(acc, http.StatusInternalServerError, errors.New("internal server error"))
		return
	}
	ui.makeUpdatePhoneToken(acc, http.StatusOK, nil)
}

//makeUpdatePhoneToken 生成更新手机号响应token
func (ui *updateInfoSrv) makeUpdatePhoneToken(acc string, code int, err error) {
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
	tk, err = ui.token.Generate(resmap)
	res := updatePhoneRes{
		res: &userproto.UserUpdatePhoneRes{
			Account: acc,
			Data:    tk,
		},
		err: err,
	}
	ui.phoneRes <- res
	delete(ui.doing, acc)
}

//------------------------------------update nickname--------------------------
//writeNicknameReq 写入更新昵称请求数据
func (ui *updateInfoSrv) writeNicknameReq(req *userproto.UserUpdateNicknameReq) {
	ui.pool.WriteHandler(func(j chan userpayload.Job) {
		ui.nicknameReq = req
		j <- ui
	})
}

//readNicknameRes 读取更新昵称响应数据
func (ui *updateInfoSrv) readNicknameRes(ctx context.Context) (*userproto.UserUpdateNicknameRes, error) {
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case res := <-ui.nicknameRes:
			return res.res, res.err
		}
	}
}

//handleUpdateNickname 处理更新昵称
func (ui *updateInfoSrv) handleUpdateNickname() {
	req := ui.nicknameReq
	//账号检测
	acc := req.GetAccount()
	if len(acc) == 0 {
		log.Info().Msg("`account` cannot be empty")
		ui.makeUpdateNicknameToken("", http.StatusBadRequest, errors.New("`account` cannot be empty"))
		return
	}
	//token校验
	tkstr := req.GetData()
	if len(tkstr) == 0 {
		log.Info().Msg("`data` cannot be empty")
		ui.makeUpdateNicknameToken(acc, http.StatusBadRequest, errors.New("`data` cannot be empty"))
		return
	}
	//正在执行中
	if ui.doing[acc] {
		msg := acc + "is doing update password, please try again later"
		log.Info().Msg(msg)
		ui.makeUpdateNicknameToken(acc, http.StatusBadRequest, errors.New(msg))
		return
	}
	ui.doing[acc] = true
	//解析数据
	code, err := ui.token.Parse(req.GetAccount(), "update_phone", req.GetData())
	if err != nil {
		log.Info().Msg(err.Error())
		ui.makeUpdateNicknameToken(acc, code, err)
		return
	}
	tk := ui.token
	code, err = usermiddleware.CommonRequestVerify(acc, tk)
	if err != nil {
		log.Info().Msg(err.Error())
		ui.makeUpdateNicknameToken(acc, code, err)
		return
	}
	nk, ok := tk.Result["nickname"].(string)
	if !ok || len(nk) == 0 {
		log.Info().Msg("`nickname` cannot be empty")
		ui.makeUpdateNicknameToken(acc, http.StatusBadRequest, errors.New("`nickname` cannot be empty"))
		return
	}
	err = usermodel.SetUserNickname(acc, tk.UserID, nk)
	if err != nil {
		log.Info().Msg("internal server error")
		ui.makeUpdateNicknameToken(acc, http.StatusInternalServerError, err)
		return
	}
	ui.makeUpdateNicknameToken(acc, http.StatusOK, nil)
}

//makeUpdateNicknameToken 更新昵称响应token
func (ui *updateInfoSrv) makeUpdateNicknameToken(acc string, code int, err error) {
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
	tk, err = ui.token.Generate(resmap)
	res := updateNicknameRes{
		res: &userproto.UserUpdateNicknameRes{
			Account: acc,
			Data:    tk,
		},
		err: err,
	}
	ui.nicknameRes <- res
	delete(ui.doing, acc)
}

func (ui *updateInfoSrv) Do() {
	go ui.handleUpdateInfo()
}
