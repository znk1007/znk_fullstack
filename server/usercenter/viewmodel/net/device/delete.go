package devicenet

import (
	"context"
	"time"

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
