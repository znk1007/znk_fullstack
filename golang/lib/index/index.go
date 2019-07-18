package index

import (
	"log"
	"net"

	protos "github.com/znk1007/fullstack/protos/generated"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"

	"github.com/znk1007/fullstack/lib/manager/user"
	"github.com/znk1007/fullstack/lib/utils/middleware"
	security "github.com/znk1007/fullstack/lib/utils/security/server"
	// _ "github.com/znk1007/fullstack/lib/utils/socket"
)

const (
	//PORT 端口号
	PORT = "9001"
)

// StartServer 启动服务
func StartServer() {
	go startMongoServer()
	startGRPC()

}

func startMongoServer() {

}

// 启动grpc
func startGRPC() {
	sc := security.ServerConfig{
		CAFile:   "lib/utils/security/keys/ca/ca.pem",
		CertFile: "lib/utils/security/keys/server/server.pem",
		KeyFile:  "lib/utils/security/keys/server/server.key",
	}

	cdt, err := sc.GenerateServerTLSCredentials()
	if err != nil {
		log.Fatalf("GenerateServerTLSCredentials err: %v", err)
		return
	}
	opts := []grpc.ServerOption{
		grpc.Creds(cdt),
		grpc_middleware.WithUnaryServerChain(
			middleware.RecoveryInterceptor,
			middleware.LoggingInterceptor,
		),
	}
	server := grpc.NewServer(opts...)
	listen, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
		return
	}
	apis(server)
	log.Println("listening on port: ", PORT)
	server.Serve(listen)
}

// apis 接口
func apis(server *grpc.Server) {
	rigisterAPI(server)
	loginAPI(server)
	logoutAPI(server)
	checkUserIDAPI(server)
	checkSessAPI(server)
	updateNicknameAPI(server)
	updatePhotoAPI(server)
	updateOnlineAPI(server)
	activeUserAPI(server)
	removeUserAPI(server)

}

// rigistServer 注册接口
func rigisterAPI(server *grpc.Server) {
	if server == nil {
		log.Fatalln("Rigist API err: Server not started")
		return
	}
	api := &user.RegisterService{
		ResChan: make(chan *protos.RegistResponse),
		ReqChan: make(chan *protos.RegistRequest),
	}
	go api.DoRequest()
	protos.RegisterRegisterServer(server, api)
}

// loginAPI 登录接口
func loginAPI(server *grpc.Server) {
	if server == nil {
		log.Fatalln("Login API err: Server not started")
		return
	}
	api := &user.LoginService{
		ReqChan: make(chan *protos.LoginRequest),
		ResChan: make(chan *protos.LoginResponse),
	}
	go api.DoRequest()
	protos.RegisterLoginServer(server, api)
}

// loginAPI 登录接口
func logoutAPI(server *grpc.Server) {
	if server == nil {
		log.Fatalln("Login API err: Server not started")
		return
	}
	api := &user.LogoutService{
		ReqChan: make(chan *protos.LogoutRequest),
		ResChan: make(chan *protos.LogoutResponse),
	}
	go api.DoRequest()
	protos.RegisterLogoutServiceServer(server, api)
}

// checkUserIDAPI 校验用户ID
func checkUserIDAPI(server *grpc.Server) {
	if server == nil {
		log.Fatalln("CheckUserId API err: Server not started")
		return
	}
	api := &user.CheckUserIDService{
		ReqChan: make(chan *protos.CheckUserIdRequest),
		ResChan: make(chan *protos.CheckUserIdResponse),
	}
	go api.DoRequest()
	protos.RegisterCheckUserIdServer(server, api)
}

// checkSessAPI 会话id接口
func checkSessAPI(server *grpc.Server) {
	if server == nil {
		log.Fatalln("CheckUserId API err: Server not started")
		return
	}
	api := &user.CheckSessService{
		ReqChan: make(chan *protos.CheckSessionRequest),
		ResChan: make(chan *protos.CheckSessionResponse),
	}
	go api.DoRequest()
	protos.RegisterCheckSessionServer(server, api)
}

// updateNickname 修改昵称
func updateNicknameAPI(server *grpc.Server) {
	if server == nil {
		log.Fatalln("CheckUserId API err: Server not started")
		return
	}
	api := &user.UpdateNicknameService{
		ReqChan: make(chan *protos.UpdateNicknameRequest),
		ResChan: make(chan *protos.UpdateNicknameResponse),
	}
	go api.DoRequest()
	protos.RegisterUpdateNicknameServer(server, api)
}

// updatePhoto 修改头像
func updatePhotoAPI(server *grpc.Server) {
	if server == nil {
		log.Fatalln("CheckUserId API err: Server not started")
		return
	}
	api := &user.UpdatePhotoService{
		ReqChan: make(chan *protos.UpdatePhotoRequest),
		ResChan: make(chan *protos.UpdatePhotoResponse),
	}
	go api.DoRequest()
	protos.RegisterUpdatePhotoServer(server, api)
}

// updateOnlineAPI 修改在线状态
func updateOnlineAPI(server *grpc.Server) {
	if server == nil {
		log.Fatalln("CheckUserId API err: Server not started")
		return
	}
	api := &user.UpdateOnlineService{
		ReqChan: make(chan *protos.UpdateOnlineRequest),
		ResChan: make(chan *protos.UpdateOnlineResponse),
	}
	go api.DoRequest()
	protos.RegisterUpdateOnlineServer(server, api)
}

// 是否禁用用户接口
func activeUserAPI(server *grpc.Server) {
	if server == nil {
		log.Fatalln("CheckUserId API err: Server not started")
		return
	}
	api := &user.ActiveUserService{
		ReqChan: make(chan *protos.ActiveUserRequest),
		ResChan: make(chan *protos.ActiveUserResponse),
	}
	go api.DoRequest()
	protos.RegisterActiveUserServer(server, api)
}

// removeUser 物理删除用户
func removeUserAPI(server *grpc.Server) {
	if server == nil {
		log.Fatalln("CheckUserId API err: Server not started")
		return
	}
	api := &user.RemoveUserService{
		ReqChan: make(chan *protos.RemoveUserRequest),
		ResChan: make(chan *protos.RemoveUserResponse),
	}
	go api.DoRequest()
	protos.RegisterRemoveUserServer(server, api)
}
