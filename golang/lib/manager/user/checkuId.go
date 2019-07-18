package user

import (
	context "context"
	checkId "github.com/znk_fullstack/golang/protos/generated"

	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	userdb "github.com/znk_fullstack/golang/lib/utils/database/user"
	security "github.com/znk_fullstack/golang/lib/utils/security/server"
)

// CheckUserIDService 校验UserID
type CheckUserIDService struct {
	token   security.Token
	ReqChan chan *checkId.CheckUserIdRequest
	ResChan chan *checkId.CheckUserIdResponse
}

// DoRequest 处理请求信息
func (cs *CheckUserIDService) DoRequest() {
	for {
		select {
		case req := <-cs.ReqChan:
			for i, length := 0, len(cs.ReqChan); ; {
				acc := req.GetAccount()
				dvc := req.GetDevice()
				if acc == "" || dvc == "" {
					cs.ResChan <- &checkId.CheckUserIdResponse{
						Code:    1,
						Message: "parameters invalid, cannot be empty",
						UserId:  "",
						Status:  0,
					}
				} else {
					msg := "operation failed"
					userID, err := userdb.GetUserID(acc)
					if err == nil && userID != "" {
						msg = "opeeration success"
					} else {
						if userID == "" {
							msg = "get userId failed"
						} else {
							msg = err.Error()
						}
					}

					cs.ResChan <- &checkId.CheckUserIdResponse{
						Code:    1,
						Message: msg,
						Status:  1,
						UserId:  userID,
					}
				}

				if i++; i > length {
					break
				}
				req = <-cs.ReqChan
			}
		}
	}
}

// Check 接口
func (cs *CheckUserIDService) Check(ctx context.Context, req *checkId.CheckUserIdRequest) (*checkId.CheckUserIdResponse, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, "LoginService canceled")
	}
	if ok := cs.token.Check(ctx); !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata.FromIncomingContext err")
	}
	cs.ReqChan <- req
	res := <-cs.ResChan
	return res, nil
}
