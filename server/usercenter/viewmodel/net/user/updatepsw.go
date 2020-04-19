package usernet

import (
	"context"

	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
	usermiddleware "github.com/znk_fullstack/server/usercenter/viewmodel/middleware"
	userpayload "github.com/znk_fullstack/server/usercenter/viewmodel/payload"
)

const (
	updatePswExpired = 60 * 2
)

//updatePswRes 更新密码响应
type updatePswRes struct {
	res *userproto.UserUpdatePswRes
	err error
}

//updatePswSrv 更新密码服务
type updatePswSrv struct {
	req     *userproto.UserUpdatePswReq
	resChan chan updatePswRes
	doing   map[string]bool
	token   usermiddleware.Token
	pool    userpayload.WorkerPool
}

//newUpdatePswSrv 初始化更新密码服务
func newUpdatePswSrv() *updatePswSrv {
	srv := &updatePswSrv{
		resChan: make(chan updatePswRes),
		doing:   make(map[string]bool),
		token:   usermiddleware.NewToken(updatePswExpired),
		pool:    userpayload.NewWorkerPool(100),
	}
	srv.pool.Run()
	return srv
}

//write 写入数据
func (up *updatePswSrv) write(req *userproto.UserUpdatePswReq) {
	up.pool.WriteHandler(func(j chan userpayload.Job) {
		up.req = req
		j <- up
	})
}

//read 读取数据
func (up *updatePswSrv) read(ctx context.Context) (res *userproto.UserUpdatePswRes, err error) {
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case res := <-up.resChan:
			return res.res, res.err
		}
	}
}

//handlUpdatePsw 处理更新密码
func (up *updatePswSrv) handlUpdatePsw() {

}

func (up *updatePswSrv) Do() {
	go up.handlUpdatePsw()
}
