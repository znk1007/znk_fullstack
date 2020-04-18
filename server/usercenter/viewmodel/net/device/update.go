package devicenet

import (
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"
	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
	usermiddleware "github.com/znk_fullstack/server/usercenter/viewmodel/middleware"
	userpayload "github.com/znk_fullstack/server/usercenter/viewmodel/payload"
)

const (
	updateExpired = 60 * 5
)

type updateRes struct {
	res *userproto.DvsUpdateRes
	err error
}
type updateSrv struct {
	req     *userproto.DvsUpdateReq
	resChan chan updateRes
	doing   map[string]bool
	pool    userpayload.WorkerPool
	token   usermiddleware.Token
}

//newUpdateSrv new更新服务
func newUpdateSrv() *updateSrv {
	srv := &updateSrv{
		resChan: make(chan updateRes),
		doing:   make(map[string]bool),
		pool:    userpayload.NewWorkerPool(100),
		token:   usermiddleware.NewToken(updateExpired),
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
func (us *updateSrv) read() (res *userproto.DvsUpdateRes, err error) {
	for {
		select {
		case res := <-us.resChan:
			return res.res, res.err
		}
	}
}

/*
用户ID：userID，
密码：password，
时间戳：timestamp，
设备ID：deviceID，
设备名：deviceName，
信任状态：state，
应用标识：appkey
*/

//handleUpdateDevice 处理设备更新
func (us *updateSrv) handleUpdateDevice() {
	acc := us.req.GetAccount()
	if len(acc) == 0 {
		log.Info().Msg("miss `account` or account cannot be empty")
		us.makeUpdateDeviceToken("", http.StatusBadRequest, errors.New("miss `account` or account cannot be empty"))
		return
	}
	if us.doing[acc] {
		log.Info().Msg("account is operating, please do it later")
		us.makeUpdateDeviceToken(acc, http.StatusBadRequest, errors.New("account is operating, please do it later"))
		return
	}
	us.doing[acc] = true
}

/*
状态码：code，
反馈消息：message，
时间戳：timestamp
*/
//makeUpdateDeviceToken 生成token
func (us *updateSrv) makeUpdateDeviceToken(acc string, code int, err error) {
	resmap := map[string]interface{}{
		"code":    code,
		"message": err.Error(),
	}
	var tk string
	tk, err = us.token.Generate(resmap)
	res := updateRes{
		err: err,
		res: &userproto.DvsUpdateRes{
			Account: acc,
			Token:   tk,
		},
	}
	//删除正在操作状态
	delete(us.doing, acc)
	us.resChan <- res
}

func (us *updateSrv) Do() {
	go us.handleUpdateDevice()
}
