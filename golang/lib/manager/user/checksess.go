package user

import (
	context "context"
	checkSess "znk/golang/protos/generated"

	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	userdb "znk/golang/lib/utils/database/user"
	security "znk/golang/lib/utils/security/server"
)

// CheckSessService 校验会话id
type CheckSessService struct {
	token   security.Token
	ReqChan chan *checkSess.CheckSessionRequest
	ResChan chan *checkSess.CheckSessionResponse
}

// DoRequest 处理请求
func (cs *CheckSessService) DoRequest() {
	for {
		select {
		case req := <-cs.ReqChan:
			for i, length := 0, len(cs.ReqChan); ; {
				uID := req.GetUserId()
				sess := req.GetSessionId()
				dvc := req.GetDevice()
				if uID == "" || sess == "" || dvc == "" {
					cs.ResChan <- &checkSess.CheckSessionResponse{
						Message: "params invalid",
						Code:    1,
						IsValid: false,
						Status:  0,
					}
				} else {
					sessionID := userdb.GetSessionID(uID)
					if sessionID == "" || sessionID != sess {
						cs.ResChan <- &checkSess.CheckSessionResponse{
							Message: "sessionId invalid",
							Code:    1,
							IsValid: false,
							Status:  0,
						}
					} else {
						cs.ResChan <- &checkSess.CheckSessionResponse{
							Message: "sessionId valid",
							Code:    1,
							IsValid: true,
							Status:  1,
						}
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

// CheckSession 校验会话ID
func (cs *CheckSessService) CheckSession(ctx context.Context, req *checkSess.CheckSessionRequest) (*checkSess.CheckSessionResponse, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, "RegisterService canceled")
	}
	if ok := cs.token.Check(ctx); !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata.FromIncomingContext err")
	}
	cs.ReqChan <- req
	res := <-cs.ResChan
	return res, nil
}
