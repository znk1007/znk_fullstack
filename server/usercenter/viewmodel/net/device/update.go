package devicenet

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	devicemodel "github.com/znk_fullstack/server/usercenter/model/device"
	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
	usermiddleware "github.com/znk_fullstack/server/usercenter/viewmodel/middleware"
	userpayload "github.com/znk_fullstack/server/usercenter/viewmodel/payload"
)

const (
	updateExpired = 60 * 5
	updateFreqExp = 60 * 10
)

type updateRes struct {
	res *userproto.DvsUpdateRes
	err error
}
type updateSrv struct {
	req     *userproto.DvsUpdateReq
	resChan chan updateRes
	doing   map[string]bool
	pool    *userpayload.WorkerPool
	token   *usermiddleware.Token
}

//newUpdateSrv new更新服务
func newUpdateSrv() *updateSrv {
	srv := &updateSrv{
		resChan: make(chan updateRes),
		doing:   make(map[string]bool),
		pool:    userpayload.NewWorkerPool(100),
		token:   usermiddleware.NewToken(updateExpired, updateFreqExp),
	}
	srv.pool.Run()
	return srv
}

//write 写入数据
func (us *updateSrv) write(req *userproto.DvsUpdateReq) {
	us.pool.WriteHandler(func(j chan userpayload.Job) {
		us.req = req
		j <- us
	})
}

//read 读取数据
func (us *updateSrv) read(ctx context.Context) (res *userproto.DvsUpdateRes, err error) {
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case res := <-us.resChan:
			return res.res, res.err
		}
	}
}

/*
用户ID：userID，
时间戳：timestamp，
设备ID：deviceID，
设备名：deviceName，
信任状态：state，
应用标识：appkey
*/

//handleUpdateDevice 处理设备更新
func (us *updateSrv) handleUpdateDevice() {
	req := us.req
	acc := req.GetAccount()
	//校验账号是否为空
	if len(acc) == 0 {
		log.Info().Msg("miss `account` or account cannot be empty")
		us.makeUpdateDeviceToken("", http.StatusBadRequest, errors.New("miss `account` or account cannot be empty"))
		return
	}
	//校验token是否为空
	tkstr := req.GetData()
	if len(tkstr) == 0 {
		msg := acc + "- `data` cannot be empty"
		log.Info().Msg(msg)
		us.makeUpdateDeviceToken("", http.StatusBadRequest, errors.New(msg))
		return
	}
	//是否正在请求
	if us.doing[acc] {
		msg := acc + "is update device, please try again later"
		log.Info().Msg(msg)
		us.makeUpdateDeviceToken(acc, http.StatusBadRequest, errors.New(msg))
		return
	}
	us.doing[acc] = true
	//校验token
	tk := us.token
	code, err := tk.Parse(acc, "update_device", tkstr)
	if err != nil {
		msg := acc + "- update device error: " + err.Error()
		log.Info().Msg(msg)
		us.makeUpdateDeviceToken(acc, code, err)
		return
	}
	//通用校验
	code, err = usermiddleware.CommonRequestVerify(acc, tk)
	if err != nil {
		msg := acc + "- update device error: " + err.Error()
		log.Info().Msg(msg)
		us.makeUpdateDeviceToken(acc, code, err)
		return
	}
	//请求数据
	res := tk.Result
	//校验userID
	userID, ok := res["userID"].(string)
	if !ok || len(userID) == 0 {
		msg := acc + "- `userID` cannot be empty"
		log.Info().Msg(msg)
		us.makeUpdateDeviceToken(acc, http.StatusBadRequest, errors.New("userID cannot be empty"))
		return
	}
	//校验state
	statestr, ok := res["state"].(string)
	if !ok || len(statestr) == 0 {
		msg := acc + "- `state` cannot be empty"
		log.Info().Msg(msg)
		us.makeUpdateDeviceToken(acc, http.StatusBadRequest, errors.New(msg))
		return
	}
	state, err := strconv.Atoi(statestr)
	if err != nil {
		msg := acc + "- update device error: " + err.Error()
		log.Info().Msg(msg)
		us.makeUpdateDeviceToken(acc, http.StatusBadRequest, err)
		return
	}
	//更新数据
	err = devicemodel.SetCurrentDevice(userID, tk.DeviceID, tk.DeviceName, tk.Platform, devicemodel.DeviceState(state), true)
	if err != nil {
		msg := acc + "- update device error: " + err.Error()
		log.Info().Msg(msg)
		us.makeUpdateDeviceToken(acc, http.StatusInternalServerError, errors.New(msg))
		return
	}
	us.makeUpdateDeviceToken(acc, http.StatusOK, nil)
	return
}

/*
状态码：code，
反馈消息：message，
时间戳：timestamp
*/
//makeUpdateDeviceToken 生成token
func (us *updateSrv) makeUpdateDeviceToken(acc string, code int, err error) {
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
	tk, err = us.token.Generate(resmap)
	res := updateRes{
		err: err,
		res: &userproto.DvsUpdateRes{
			Account: acc,
			Data:    tk,
		},
	}
	//删除正在操作状态
	delete(us.doing, acc)
	us.resChan <- res
}

func (us *updateSrv) Do() {
	go us.handleUpdateDevice()
}
