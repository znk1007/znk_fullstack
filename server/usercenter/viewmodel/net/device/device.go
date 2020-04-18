package usernet

import (
	"context"

	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
)

type deviceSrv struct {
}

//UpdateDevice 更新设备信息
func (ds *deviceSrv) UpdateDevice(ctx context.Context, req *userproto.DvsUpdateReq) (res *userproto.DvsUpdateRes, err error) {
	return
}

//DeleteDevice 删除设备
func (ds *deviceSrv) DeleteDevice(ctx context.Context, req *userproto.DvsDeleteReq) (res *userproto.DvsDeleteRes, err error) {
	return
}
