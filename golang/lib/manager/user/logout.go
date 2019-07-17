package user

import (
	context "context"
	logout "znk/protos/generated"

	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	userdb "znk/lib/utils/database/user"
	security "znk/lib/utils/security/server"
)

// LogoutService 退出登录
type LogoutService struct {
	token   security.Token
	ReqChan chan *logout.LogoutRequest
	ResChan chan *logout.LogoutResponse
}

// DoRequest 处理请求
func (ls *LogoutService) DoRequest() {
	for {
		select {
		case req := <-ls.ReqChan:
			for i, length := 0, len(ls.ReqChan); ; {
				uID := req.GetUserId()
				sessID := req.GetSessionId()
				if uID == "" || sessID == "" {
					ls.ResChan <- &logout.LogoutResponse{
						Message: "parameters invalid, cannot be empty",
						Code:    1,
						Status:  0,
					}
				} else {
					msg := "operation failed"
					var stat int32
					var code int32
					err := userdb.UpdateSessionIDAndOnlineState(uID, "default", 1, false)
					if err == nil {
						stat = 1
						code = 1
						msg = "operation success"
					}
					ls.ResChan <- &logout.LogoutResponse{
						Message: msg,
						Code:    code,
						Status:  stat,
					}
				}
				if i++; i > length {
					break
				}
				req = <-ls.ReqChan
			}
		}
	}
}

// Logout 退出登录请求
func (ls *LogoutService) Logout(ctx context.Context, req *logout.LogoutRequest) (*logout.LogoutResponse, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, "LogoutService canceled")
	}
	if ok := ls.token.Check(ctx); !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata.FromIncomingContext err")
	}
	ls.ReqChan <- req
	res := <-ls.ResChan
	return res, nil
}
