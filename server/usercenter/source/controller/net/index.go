package indexnet

import (
	"net"

	"github.com/rs/zerolog/log"
	userconf "github.com/znk_fullstack/server/usercenter/source/controller/conf"
	usernet "github.com/znk_fullstack/server/usercenter/source/controller/net/user"
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
	usernet.RegisterRegistServer(s)
}