package user

import (
	context "context"
	"strconv"
	login "znk/protos/generated"

	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	userdb "znk/lib/utils/database/user"
	crypto "znk/lib/utils/security"
	security "znk/lib/utils/security/server"
)

// LoginService 登录端
type LoginService struct {
	token   security.Token
	ReqChan chan *login.LoginRequest
	ResChan chan *login.LoginResponse
}

// DoRequest 处理请求
func (ls *LoginService) DoRequest() {
	for {
		select {
		case req := <-ls.ReqChan:
			for i, length := 0, len(ls.ReqChan); ; {
				acc := req.GetAccount()
				psw := req.GetPassword()
				uID := req.GetUserId()
				dvc := req.GetDevice()
				if acc == "" || psw == "" || uID == "" || dvc == "" {
					ls.ResChan <- &login.LoginResponse{
						Message: "parameters invalid, cannot be empty",
						Code:    1,
						User:    &login.User{},
						Status:  0,
					}
				} else {
					exists, err := userdb.IsUserExists(acc) //u.UserExists()
					if err != nil {
						ls.ResChan <- &login.LoginResponse{
							Message: err.Error(),
							Code:    1,
							User:    &login.User{},
							Status:  0,
						}
					} else if !exists {
						ls.ResChan <- &login.LoginResponse{
							Message: "user not found",
							Code:    1,
							User:    &login.User{},
							Status:  0,
						}
					} else {
						msg := "operation failed"
						commonU := &login.User{}
						dbUser := &userdb.User{}
						err := dbUser.GetActiveUser(uID)
						var stat int32
						if err != nil {
							msg = err.Error()
						} else {
							if uID != "" {
								password, orgErr := crypto.AESDecode(psw)
								savedPsw, savedErr := crypto.AESDecode(dbUser.Password)
								if orgErr == nil && savedErr == nil && password == savedPsw {
									sessionID := security.GenterateSessionID()
									err := userdb.UpdateSessionIDAndOnlineState(uID, sessionID, 60*60*24*7, true) //u.UpdateSessionIDAndOnlineState(sessionID, 60*60*24*7)
									if err == nil {
										msg = "operation success"
										stat = 1
										commonU = &login.User{
											SessionId: sessionID,
											Account:   dbUser.Account,
											Nickname:  dbUser.Nickname,
											Phone:     dbUser.Phone,
											Email:     dbUser.Email,
											Photo:     dbUser.Photo,
											UserId:    dbUser.UserID,
											CreatedAt: strconv.FormatInt(dbUser.CreatedAt.Unix(), 10),
										}
									}
								}
							}
						}
						ls.ResChan <- &login.LoginResponse{
							Message: msg,
							Code:    1,
							User:    commonU,
							Status:  stat,
						}
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

// Login 登录接口
func (ls *LoginService) Login(ctx context.Context, req *login.LoginRequest) (*login.LoginResponse, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, "LoginService canceled")
	}
	if ok := ls.token.Check(ctx); !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata.FromIncomingContext err")
	}
	ls.ReqChan <- req
	res := <-ls.ResChan
	return res, nil
}
