package user

import (
	context "context"
	updatenickname "znk/protos/generated"

	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	userdb "znk/lib/utils/database/user"
	security "znk/lib/utils/security/server"
)

// UpdateNicknameService 修改昵称
type UpdateNicknameService struct {
	token   security.Token
	ReqChan chan *updatenickname.UpdateNicknameRequest
	ResChan chan *updatenickname.UpdateNicknameResponse
}

// DoRequest 处理请求
func (un *UpdateNicknameService) DoRequest() {
	for {
		select {
		case req := <-un.ReqChan:
			for i, length := 0, len(un.ReqChan); ; {
				acc := req.GetAccount()
				sID := req.GetSessionId()
				uID := req.GetUserId()
				nn := req.GetNickname()
				dvc := req.GetDevice()
				if sID == "" || uID == "" || nn == "" || dvc == "" {
					un.ResChan <- &updatenickname.UpdateNicknameResponse{
						Message: "parameters invalid, cannot be empty",
						Code:    1,
						Status:  0,
					}
				} else {

					exists, err := userdb.IsUserExists(acc)
					if err != nil {
						un.ResChan <- &updatenickname.UpdateNicknameResponse{
							Message: err.Error(),
							Code:    1,
							Status:  0,
						}
					} else if !exists {
						un.ResChan <- &updatenickname.UpdateNicknameResponse{
							Message: "user not exists",
							Code:    1,
							Status:  0,
						}
					} else {
						msg := "operation success"
						err := userdb.UpdateNickname(uID, sID, nn) //u.UpdateNickname(sID)
						if err != nil {
							msg = err.Error()
						}
						un.ResChan <- &updatenickname.UpdateNicknameResponse{
							Message: msg,
							Code:    1,
							Status:  1,
						}
					}
				}

				if i++; i > length {
					break
				}
				req = <-un.ReqChan
			}
		}
	}
}

// Update 更新昵称接口
func (un *UpdateNicknameService) Update(ctx context.Context, req *updatenickname.UpdateNicknameRequest) (*updatenickname.UpdateNicknameResponse, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, "LoginService canceled")
	}
	if ok := un.token.Check(ctx); !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata.FromIncomingContext err")
	}
	un.ReqChan <- req
	res := <-un.ResChan
	return res, nil
}
