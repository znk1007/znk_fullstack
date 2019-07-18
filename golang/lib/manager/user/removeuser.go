package user

import (
	context "context"
	removeuser "github.com/znk1007/fullstack/protos/generated"

	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	utils "github.com/znk1007/fullstack/lib/utils/common"
	userdb "github.com/znk1007/fullstack/lib/utils/database/user"
	security "github.com/znk1007/fullstack/lib/utils/security/server"
)

// RemoveUserService 物理删除用户
type RemoveUserService struct {
	token   security.Token
	ReqChan chan *removeuser.RemoveUserRequest
	ResChan chan *removeuser.RemoveUserResponse
}

// DoRequest 处理请求
func (ru *RemoveUserService) DoRequest() {
	for {
		select {
		case req := <-ru.ReqChan:
			for i, length := 0, len(ru.ReqChan); ; {
				accs := req.GetAccounts()
				acc := req.GetAdminAccount()
				uID := req.GetAdminUserId()
				dvc := req.GetDevice()
				if uID == "" || len(accs) == 0 || dvc == "" {
					ru.ResChan <- &removeuser.RemoveUserResponse{
						Message: "parameters invalid, cannot be empty",
						Code:    1,
						Status:  0,
					}
				} else {
					accs = utils.DeleteString(accs, acc)
					exists, err := userdb.IsUserExists(acc)
					if err != nil {
						ru.ResChan <- &removeuser.RemoveUserResponse{
							Message: err.Error(),
							Code:    1,
							Status:  1,
						}
					} else if !exists {
						ru.ResChan <- &removeuser.RemoveUserResponse{
							Message: "user not exists",
							Code:    1,
							Status:  1,
						}
					} else {
						msg := "operation success"
						err := userdb.Remove(uID, accs)
						var stat int32
						if err != nil {
							stat = 0
							msg = err.Error()
						} else {
							stat = 1
						}
						ru.ResChan <- &removeuser.RemoveUserResponse{
							Message: msg,
							Code:    1,
							Status:  stat,
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

// Remove 物理删除用户接口
func (ru *RemoveUserService) Remove(ctx context.Context, req *removeuser.RemoveUserRequest) (*removeuser.RemoveUserResponse, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, "RemoveUserService canceled")
	}
	if ok := ru.token.Check(ctx); !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata.FromIncomingContext err")
	}
	ru.ReqChan <- req
	res := <-ru.ResChan
	return res, nil
}
