package devicenet

import (
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
func (ds *deleteSrv) read() (res *userproto.DvsDeleteRes, err error) {
	for {
		select {
		case res := <-ds.resChan:
			return res.res, res.err
		}
	}
}

func (ds *deleteSrv) handlDeleteDevice() {

}

func (ds *deleteSrv) Do() {
	go ds.handlDeleteDevice()
}
