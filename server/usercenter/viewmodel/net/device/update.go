package devicenet

import (
	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
	userpayload "github.com/znk_fullstack/server/usercenter/viewmodel/payload"
)

var updatePool userpayload.WorkerPool

type updateRes struct {
	res *userproto.DvsUpdateRes
	err error
}
type updateSrv struct {
	req     *userproto.DvsUpdateReq
	resChan chan updateRes
	doing   map[string]bool
}

//newUpdateSrv new更新服务
func newUpdateSrv() *updateSrv {
	updatePool = userpayload.CreateWorkerPool(100)
	return &updateSrv{
		resChan: make(chan updateRes),
		doing:   make(map[string]bool),
	}
}

//write 写入数据
func (us *updateSrv) write(req *userproto.DvsUpdateReq) {
	updatePool.WriteHandler(func(j chan userpayload.Job) {
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

//handleUpdateDevice 处理设备更新
func (us *updateSrv) handleUpdateDevice() {

}

func (us *updateSrv) Do() {
	go us.handleUpdateDevice()
}
