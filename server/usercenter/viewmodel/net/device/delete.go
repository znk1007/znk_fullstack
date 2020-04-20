package devicenet

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
	devicemodel "github.com/znk_fullstack/server/usercenter/model/device"
	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
	usermiddleware "github.com/znk_fullstack/server/usercenter/viewmodel/middleware"
	userpayload "github.com/znk_fullstack/server/usercenter/viewmodel/payload"
)

const (
	deleteExpired = 60 * 5
)

//deleteRes 删除设备响应
type deleteRes struct {
	res *userproto.DvsDeleteRes
	err error
}

//deleteSrv 删除设备服务
type deleteSrv struct {
	req     *userproto.DvsDeleteReq
	resChan chan deleteRes
	doing   map[string]bool
	pool    userpayload.WorkerPool
	token   usermiddleware.Token
}

//newDeleteSrv 初始化删除服务
func newDeleteSrv() *deleteSrv {
	srv := &deleteSrv{
		resChan: make(chan deleteRes),
		doing:   make(map[string]bool),
		pool:    userpayload.NewWorkerPool(100),
		token:   usermiddleware.NewToken(deleteExpired),
	}
	srv.pool.Run()
	return srv
}

//write 写入数据
func (ds *deleteSrv) write(req *userproto.DvsDeleteReq) {
	ds.pool.WriteHandler(func(j chan userpayload.Job) {
		ds.req = req
		j <- ds
	})
}

//read 读取数据
func (ds *deleteSrv) read(ctx context.Context) (res *userproto.DvsDeleteRes, err error) {
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case res := <-ds.resChan:
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
应用标识：appkey
*/
//handlDeleteDevice 处理删除设备操作
func (ds *deleteSrv) handlDeleteDevice() {
	req := ds.req
	//账号
	acc := req.GetAccount()
	//校验账号是否为空
	if len(acc) == 0 {
		log.Info().Msg("miss `account` or account cannot be empty")
		ds.makeDeleteDeviceToken("", http.StatusBadRequest, errors.New("miss `account` or account cannot be empty"))
		return
	}
	//校验token是否为空
	tkstr := req.GetToken()
	if len(tkstr) == 0 {
		log.Info().Msg("token cannot be empty")
		ds.makeDeleteDeviceToken("", http.StatusBadRequest, errors.New("token cannot be empty"))
		return
	}
	//是否正在请求
	if ds.doing[acc] {
		log.Info().Msg("account is operating, please do it later")
		ds.makeDeleteDeviceToken(acc, http.StatusBadRequest, errors.New("account is operating, please do it later"))
		return
	}
	ds.doing[acc] = true
	//校验token
	tk := ds.token
	code, err := tk.VerifyByPswAndSess(tkstr)
	if err != nil {
		log.Info().Msg(err.Error())
		ds.makeDeleteDeviceToken(acc, code, err)
		return
	}
	//通用校验
	code, err = usermiddleware.CommonVerify(acc, tk)
	if err != nil {
		log.Info().Msg(err.Error())
		ds.makeDeleteDeviceToken(acc, code, err)
		return
	}
	res := tk.Result
	userID, ok := res["userID"].(string)
	if !ok || len(userID) == 0 {
		log.Info().Msg("userID cannot be empty")
		ds.makeDeleteDeviceToken(acc, http.StatusBadRequest, errors.New("userID cannot be empty"))
		return
	}
	//删除设备
	err = devicemodel.DelDevice(userID, tk.DeviceID)
	if err != nil {
		log.Info().Msg("internal server error")
		ds.makeDeleteDeviceToken(acc, http.StatusBadRequest, errors.New("internal server error"))
		return
	}
	ds.makeDeleteDeviceToken(acc, http.StatusOK, nil)
}

/*
状态码：code，
反馈消息：message，
时间戳：timestamp
*/
//makeDeleteDeviceToken 删除设备响应token
func (ds *deleteSrv) makeDeleteDeviceToken(acc string, code int, err error) {
	resmap := map[string]interface{}{
		"code":      code,
		"message":   err.Error(),
		"timestamp": time.Now().String(),
	}
	var tk string
	tk, err = ds.token.Generate(resmap)
	res := deleteRes{
		err: err,
		res: &userproto.DvsDeleteRes{
			Account: acc,
			Token:   tk,
		},
	}
	ds.resChan <- res
}

func (ds *deleteSrv) Do() {
	go ds.handlDeleteDevice()
}
