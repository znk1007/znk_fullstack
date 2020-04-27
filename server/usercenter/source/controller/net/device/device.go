package devicenet

import (
	"context"

	userproto "github.com/znk_fullstack/server/usercenter/source/model/protos/generated"
)

var dsrv *deviceSrv

//deviceSrv 设备相关服务
type deviceSrv struct {
	uSrv *updateSrv
	dSrv *deleteSrv
}

func init() {
	dsrv = &deviceSrv{
		uSrv: newUpdateSrv(),
		dSrv: newDeleteSrv(),
	}
}

//UpdateDevice 更新设备信息
func (ds *deviceSrv) UpdateDevice(ctx context.Context, req *userproto.DvsUpdateReq) (res *userproto.DvsUpdateRes, err error) {
	ds.uSrv.write(req)
	return ds.uSrv.read(ctx)
}

//DeleteDevice 删除设备
func (ds *deviceSrv) DeleteDevice(ctx context.Context, req *userproto.DvsDeleteReq) (res *userproto.DvsDeleteRes, err error) {
	ds.dSrv.write(req)
	return ds.dSrv.read(ctx)
}
