package usernet

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
	userproto "github.com/znk_fullstack/server/usercenter/model/protos/generated"
	userredis "github.com/znk_fullstack/server/usercenter/viewmodel/dao/redis"
	usermiddleware "github.com/znk_fullstack/server/usercenter/viewmodel/middleware"
	userpayload "github.com/znk_fullstack/server/usercenter/viewmodel/payload"
	"google.golang.org/grpc"
)

var rs *registService
var check usermiddleware.CheckToken

func init() {
	rs = &registService{
		resChan: make(chan registResponse),
	}
	check = usermiddleware.Create(1000 * 60 * 5)
}

//registResponse 注册响应
type registResponse struct {
	res *userproto.RegistRes
	err error
}

//RegistService 注册
type registService struct {
	req     *userproto.RegistReq
	resChan chan registResponse
}

func (s registService) Do() {
	go s.handleRegist()
}

func (s registService) handleRegist() {
	req := s.req
	acc := req.GetAccount()
	fmt.Println(acc)
	genRes := func(acc string, tk string, e error) {
		res := registResponse{
			res: &userproto.RegistRes{
				Account: "",
				Token:   tk,
			},
			err: e,
		}
		s.resChan <- res
	}
	//第一层redis，防止频繁操作数据库
	exists := userredis.Exists(acc)
	if exists {
		tk, e := s.generateRegistToken(acc, http.StatusBadRequest, acc+"has registed")
		genRes(acc, tk, e)
		return
	}
	if len(acc) == 0 {
		log.Info().Msg("account cannot be empty")
		tk, e := s.generateRegistToken("", http.StatusBadRequest, "account cannot be empty")
		if e != nil {
			genRes("", tk, e)
			return
		}
		genRes("", tk, errors.New("account cannot be empty"))
		return
	}
	check.UJWT.Parse(s.req.GetToken())
	tkMap, _, e := check.UJWT.Result()
	if e != nil {
		log.Info().Msg(e.Error())
		tk, e := s.generateRegistToken(acc, http.StatusBadRequest, e.Error())
		genRes(acc, tk, e)
		return
	}

	_, e = s.checkRegistToken(tkMap)
	if e != nil {
		log.Info().Msg(e.Error())
		tk, e := s.generateRegistToken(acc, http.StatusBadRequest, e.Error())
		genRes(acc, tk, e)
		return
	}

	tsstr := acc + "_" + string(time.Now().Unix())
	userredis.HSet(acc, tsstr)
}

/*
参数密码：password[CBCEncrypt]，
设备ID：deviceID，
平台：platform[web,iOS,Android]，
时间戳：timestamp，
应用标识：appkey
*/

func (s registService) checkRegistToken(reqMap map[string]interface{}) (tk string, err error) {
	var deviceID string
	var password string
	var platform string
	dID, ok := reqMap["deviceID"]
	deviceID = dID.(string)
	if !ok || len(deviceID) == 0 {
		log.Info().Msg("deviceID cannot be empty")
		tk, err = s.generateRegistToken("", http.StatusBadRequest, "deviceID cannot be empty")
		return
	}
	plf, ok := reqMap["platform"]
	platform = plf.(string)
	if !ok || len(platform) == 0 {
		log.Info().Msg("platform cannot be empty")
		tk, err = s.generateRegistToken("", http.StatusBadRequest, "platform cannot be empty")
		return
	}
	psd, ok := reqMap["password"]
	password = psd.(string)
	if !ok || len(password) == 0 {
		log.Info().Msg("password cannot be empty")
		tk, err = s.generateRegistToken("", http.StatusBadRequest, "password cannot be empty")
		return
	}
	// _, e := usercrypto.CBCEncrypt(password)
	// if e != nil {
	// 	log.Info().Msg("encrypt password err: " + e.Error())
	// 	tk, err = s.generateRegistToken("", http.StatusInternalServerError, "interval server error")
	// 	return
	// }

	return
}

/*
用户ID：userID，
时间戳：timestamp，
状态码：code，
反馈消息：message
*/
func (s registService) generateRegistToken(userID string, code int, msg string) (tk string, err error) {
	ts := time.Now().Unix()
	resMap := map[string]interface{}{
		"timestamp": ts,
		"code":      code,
		"message":   msg,
	}
	if code == http.StatusOK {
		resMap["userID"] = userID
	}
	tk, err = check.UJWT.Token(resMap)
	return
}

//RegisterRegistServer 注册到注册请求服务
func RegisterRegistServer(srv *grpc.Server) {
	userproto.RegisterRegistSrvServer(srv, rs)
}

//UserReigst 注册
func (s registService) UserReigst(ctx context.Context, req *userproto.RegistReq) (*userproto.RegistRes, error) {
	userpayload.Pool.WriteHandler(func(jq chan userpayload.Job) {
		s.req = req
		jq <- s
	})
	for {
		select {
		case res := <-s.resChan:
			return res.res, res.err
		}
	}
}
