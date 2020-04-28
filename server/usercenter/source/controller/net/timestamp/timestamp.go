package userts

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
	usermiddleware "github.com/znk_fullstack/server/usercenter/source/controller/middleware"
	userpayload "github.com/znk_fullstack/server/usercenter/source/controller/payload"
	userproto "github.com/znk_fullstack/server/usercenter/source/model/protos/generated"
)

const (
	timestampExpired = 60 * 5
	timestampFreqExp = 60 * 10
)

//tsRes 时间戳响应
type tsRes struct {
	res *userproto.TSRes
	err error
}

//tsSrv 时间戳服务
type tsSrv struct {
	req   *userproto.TSReq
	res   chan tsRes
	token *usermiddleware.Token
	pool  *userpayload.WorkerPool
}

//newTsSrv 初始化时间戳接口
func newTsSrv() *tsSrv {
	srv := &tsSrv{
		res:   make(chan tsRes),
		token: usermiddleware.NewToken(timestampExpired, timestampFreqExp),
		pool:  userpayload.NewWorkerPool(100),
	}
	srv.pool.Run()
	return srv
}

//writeTsReq 写入数据
func (ts *tsSrv) writeTsReq(req *userproto.TSReq) {
	ts.pool.WriteHandler(func(j chan userpayload.Job) {
		ts.req = req
		j <- ts
	})
}

//readTsRes 读取数据
func (ts *tsSrv) readTsRes(ctx context.Context) (*userproto.TSRes, error) {
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case res := <-ts.res:
			return res.res, res.err
		}
	}
}

//handleTimestamp 处理请求
func (ts *tsSrv) handleTimestamp() {
	req := ts.req
	//平台
	plf := req.GetPlatform()
	if len(plf) == 0 {
		msg := "`platform` cannot be empty"
		log.Info().Msg(msg)
		ts.makeTimestampToken(http.StatusBadRequest, errors.New(msg))
		return
	}
	dID := req.GetDeviceID()
	if len(dID) == 0 {
		msg := "`deviceID` cannot be empty"
		log.Info().Msg(msg)
		ts.makeTimestampToken(http.StatusBadRequest, errors.New(msg))
		return
	}
	dName := req.GetDeviceName()
	if len(dName) == 0 {
		msg := "`deviceName` cannot be empty"
		log.Info().Msg(msg)
		ts.makeTimestampToken(http.StatusBadRequest, errors.New(msg))
		return
	}
	ts.makeTimestampToken(http.StatusOK, nil)
}

func (ts *tsSrv) makeTimestampToken(code int32, err error) {
	msg := "operation success"
	if err != nil {
		msg = err.Error()
	}
	res := tsRes{
		err: err,
		res: &userproto.TSRes{
			Code:      code,
			Message:   msg,
			Timestamp: time.Now().String(),
		},
	}
	ts.res <- res
}

func (ts *tsSrv) Do() {
	go ts.handleTimestamp()
}
