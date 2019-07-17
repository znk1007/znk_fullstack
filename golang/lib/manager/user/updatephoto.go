package user

import (
	context "context"
	updatephoto "znk/protos/generated"

	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	userdb "znk/lib/utils/database/user"
	security "znk/lib/utils/security/server"
)

// UpdatePhotoService 修改头像
type UpdatePhotoService struct {
	token   security.Token
	ReqChan chan *updatephoto.UpdatePhotoRequest
	ResChan chan *updatephoto.UpdatePhotoResponse
}

// DoRequest 处理请求
func (up *UpdatePhotoService) DoRequest() {
	for {
		select {
		case req := <-up.ReqChan:
			for i, length := 0, len(up.ReqChan); ; {
				acc := req.GetAccount()
				sID := req.GetSessionId()
				uID := req.GetUserId()
				pt := req.GetPhoto()
				dvc := req.GetDevice()
				if sID == "" || uID == "" || pt == "" || dvc == "" {
					up.ResChan <- &updatephoto.UpdatePhotoResponse{
						Message: "parameters invalid, cannot be empty",
						Code:    1,
						Status:  0,
					}
				} else {

					exists, err := userdb.IsUserExists(acc)
					if err != nil {
						up.ResChan <- &updatephoto.UpdatePhotoResponse{
							Message: err.Error(),
							Code:    1,
							Status:  1,
						}
					} else if !exists {
						up.ResChan <- &updatephoto.UpdatePhotoResponse{
							Message: "user not exists",
							Code:    1,
							Status:  1,
						}
					} else {
						msg := "operation success"
						err := userdb.UpdatePhoto(uID, sID, pt)
						if err != nil {
							msg = err.Error()
						}
						up.ResChan <- &updatephoto.UpdatePhotoResponse{
							Message: msg,
							Code:    1,
							Status:  1,
						}
					}
				}

				if i++; i > length {
					break
				}
				req = <-up.ReqChan
			}
		}
	}
}

// Update 修改头像接口
func (up *UpdatePhotoService) Update(ctx context.Context, req *updatephoto.UpdatePhotoRequest) (*updatephoto.UpdatePhotoResponse, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, "LoginService canceled")
	}
	if ok := up.token.Check(ctx); !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata.FromIncomingContext err")
	}
	up.ReqChan <- req
	res := <-up.ResChan
	return res, nil
}
