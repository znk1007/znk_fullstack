package devicenet

import (
	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
	userpayload "github.com/znk_fullstack/server/usercenter/viewmodel/payload"
)

var deletePool userpayload.WorkerPool

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
}

//newDeleteSrv 初始化删除服务
func newDeleteSrv() *deleteSrv {
	deletePool = userpayload.CreateWorkerPool(100)
	return &deleteSrv{
		resChan: make(chan deleteRes),
		doing:   make(map[string]bool),
	}
}

//write 写入数据
func (ds *deleteSrv) write(req *userproto.DvsDeleteReq) {
	deletePool.WriteHandler(func(j chan userpayload.Job) {
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

func (ds *deleteSrv) Do() {

}
