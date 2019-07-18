package user

import (
	context "context"
	activeuser "github.com/znk1007/fullstack/protos/generated"

	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	utils "github.com/znk1007/fullstack/lib/utils/common"
	userdb "github.com/znk1007/fullstack/lib/utils/database/user"
	security "github.com/znk1007/fullstack/lib/utils/security/server"
)

// ActiveUserService 禁用用户
type ActiveUserService struct {
	token   security.Token
	ReqChan chan *activeuser.ActiveUserRequest
	ResChan chan *activeuser.ActiveUserResponse
}

// DoRequest 处理请求
func (ru *ActiveUserService) DoRequest() {
	for {
		select {
		case req := <-ru.ReqChan:
			for i, length := 0, len(ru.ReqChan); ; {
				accs := req.GetAccounts()
				active := req.GetActive()
				acc := req.GetAdminAccount()
				uID := req.GetAdminUserId()
				dvc := req.GetDevice()
				if uID == "" || len(accs) == 0 || dvc == "" {
					ru.ResChan <- &activeuser.ActiveUserResponse{
						Message: "parameters invalid, cannot be empty",
						Code:    1,
						Status:  0,
					}
				} else {
					accs = utils.DeleteString(accs, acc)
					exists, err := userdb.IsUserExists(acc)
					if err != nil {
						ru.ResChan <- &activeuser.ActiveUserResponse{
							Message: err.Error(),
							Code:    1,
							Status:  1,
						}
					} else if !exists {
						ru.ResChan <- &activeuser.ActiveUserResponse{
							Message: "user not exists",
							Code:    1,
							Status:  1,
						}
					} else {
						msg := "operation success"
						err := userdb.ActiveUser(uID, accs, active)
						if err != nil {
							msg = err.Error()
						}
						ru.ResChan <- &activeuser.ActiveUserResponse{
							Message: msg,
							Code:    1,
							Status:  1,
						}
					}
				}

				if i++; i > length {
					break
				}
				req = <-ru.ReqChan
			}
		}
	}
}

// ActiveOr 是否禁用用户接口
func (ru *ActiveUserService) ActiveOr(ctx context.Context, req *activeuser.ActiveUserRequest) (*activeuser.ActiveUserResponse, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, "ActiveUserService canceled")
	}
	if ok := ru.token.Check(ctx); !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata.FromIncomingContext err")
	}
	ru.ReqChan <- req
	res := <-ru.ResChan
	return res, nil
}
