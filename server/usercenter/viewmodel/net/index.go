package usernet

import (
	"net"

	"github.com/rs/zerolog/log"
	"github.com/znk_fullstack/server/usercenter/model"
	userconf "github.com/znk_fullstack/server/usercenter/viewmodel/conf"
	userGenID "github.com/znk_fullstack/server/usercenter/viewmodel/generateId"
	"google.golang.org/grpc"
)

//RunRPC 运行rpc服务
func RunRPC() {
	rpcConf := userconf.RPCSrvConf()
	lis, err := net.Listen("tcp", rpcConf.Host+":"+rpcConf.Port)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	s := grpc.NewServer()
	connectSrv(s)
	if err := s.Serve(lis); err != nil {
		log.Fatal().Msg(err.Error())
	}
}

//connectSrv 连接服务
func connectSrv(s *grpc.Server) {
	registerRegistServer(s)
}

//makeID 生成唯一ID
func makeID() string {
	return userGenID.GenerateID()
}

/*
设备ID：deviceID，
平台：platform[web,iOS,Android]，
用户ID：userID，
应用标识：appkey，
*/

//saveCurrentDevice 保存当前设备信息
func saveCurrentDevice(userID string, deviceID string, platform string) (err error) {
	dvs := &model.Device{
		DeviceID: deviceID,
		Platform: platform,
		Trust:    1,
		Online:   0,
		UserID:   userID,
	}
	_, err = model.CreateDevice(dvs)
	if err != nil {
		log.Info().Msg(err.Error())
		return
	}
	err = model.SetCurrentDeivce(userID, deviceID, 1, 0)
	return
}
