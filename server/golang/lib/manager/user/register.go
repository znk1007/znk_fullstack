package user

import (
	context "context"
	register "github.com/znk_fullstack/golang/protos/generated"

	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	// "github.com/znk_fullstack/golang/lib/utils/common"
	userdb "github.com/znk_fullstack/golang/lib/utils/database/user"
	objectid "github.com/znk_fullstack/golang/lib/utils/objectId"
	regexputils "github.com/znk_fullstack/golang/lib/utils/regexps"
	security "github.com/znk_fullstack/golang/lib/utils/security/server"
)

// RegisterService 注册服务
type RegisterService struct {
	token   security.Token
	ReqChan chan *register.RegistRequest
	ResChan chan *register.RegistResponse
}

// DoRequest 处理请求消息
func (rs *RegisterService) DoRequest() {
	for {
		select {
		case req := <-rs.ReqChan:
			for i, length := 0, len(rs.ReqChan); ; {
				acc := req.GetAccount()
				psw := req.GetPassword()
				device := req.GetDevice()
				if acc == "" || psw == "" || device == "" {
					rs.ResChan <- &register.RegistResponse{
						Account: acc,
						UserId:  "",
						Code:    1,
						Status:  0,
						Message: "parameters invalid, cannot be empty",
					}
				} else {
					msg := ""
					exists, err := userdb.IsUserExists(acc)
					if err != nil {
						msg = err.Error()
						rs.ResChan <- &register.RegistResponse{
							Account: acc,
							UserId:  "",
							Code:    1,
							Status:  0,
							Message: msg,
						}
					} else if exists {
						msg = "user already exists"
						rs.ResChan <- &register.RegistResponse{
							Account: acc,
							UserId:  "",
							Code:    1,
							Status:  0,
							Message: msg,
						}
					} else {
						psd := psw
						phone := ""
						email := ""
						if regexputils.VerifyEmail(acc) {
							email = acc
						}
						if regexputils.VerifyPhone(acc) {
							phone = acc
						}
						uID := objectid.NewObjectID().String()
						u := userdb.User{
							Password: psd,
							Active:   true,
							Device:   device,
							UserID:   uID,
							Account:  acc,
							Nickname: acc,
							Phone:    phone,
							Email:    email,
							Photo:    "",
						}
						state, err := u.Insert()
						msg = "Operation Success!"
						if err != nil || state != 1 {
							msg = err.Error()
						}
						rs.ResChan <- &register.RegistResponse{
							Account: req.GetAccount(),
							UserId:  uID,
							Code:    1,
							Status:  state,
							Message: msg,
						}
					}
				}
				if i++; i > length {
					break
				}
				req = <-rs.ReqChan
			}
		}
	}
}

// Regist 注册
func (rs *RegisterService) Regist(ctx context.Context, req *register.RegistRequest) (*register.RegistResponse, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, "RegisterService canceled")
	}
	if ok := rs.token.Check(ctx); !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata.FromIncomingContext err")
	}
	rs.ReqChan <- req
	res := <-rs.ResChan
	return res, nil

}
