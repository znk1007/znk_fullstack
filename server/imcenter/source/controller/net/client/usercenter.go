package imnetclient

import (
	"context"
	"errors"
	"time"

	"github.com/rs/zerolog/log"
	imconf "github.com/znk_fullstack/server/imcenter/source/controller/conf"
	imjwt "github.com/znk_fullstack/server/imcenter/source/controller/jwt"
	userproto "github.com/znk_fullstack/server/imcenter/source/model/protos/generated"
	"google.golang.org/grpc"
)

const (
	appkey = "fullstack-manznk"
)

var ij *imjwt.IMJWT

func init() {
	ij = imjwt.NewUserJWT(60 * 10)
}

//RegistUserCenter 注册用户中心
func RegistUserCenter(acc, pwd string) (err error) {
	if len(acc) == 0 {
		err = errors.New("`account` cannot be empty")
		return
	}
	if len(pwd) == 0 {
		err = errors.New("`password` cannot be empty")
		return
	}
	creds, e := imconf.TLSCredentials()
	if e != nil {
		err = e
		log.Info().Msg(e.Error())
		return
	}
	ucEnv := imconf.GetUcEnv()
	addr := ucEnv.Host + ":" + ucEnv.Port
	conn, e := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(creds),
		grpc.WithBlock(),
	)
	if e != nil {
		err = e
		log.Info().Msg(e.Error())
		return
	}
	defer conn.Close()
	datamap := map[string]interface{}{
		"password":   pwd,
		"deviceID":   "imcenter_device",
		"platform":   "imcenter_platform",
		"deviceName": "imcenter",
		"appkey":     appkey,
	}
	var data string
	data, e = ij.Token(datamap)
	if e != nil {
		err = e
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rc := userproto.NewUserSrvClient(conn)
	var res *userproto.UserRgstRes
	res, e = rc.Regist(ctx, &userproto.UserRgstReq{
		Account: acc,
		Data:    data,
	})
	if e != nil {
		err = e
		log.Info().Msg(e.Error())
		return
	}
	res.Data
	return
}

//LoginUserCenter 登录用户中心
func LoginUserCenter(acc, pwd string) (err error) {
	if len(acc) == 0 {
		err = errors.New("`account` cannot be empty")
		return
	}
	if len(pwd) == 0 {
		err = errors.New("`password` cannot be empty")
		return
	}
	creds, e := imconf.TLSCredentials()
	if e != nil {
		err = e
		log.Info().Msg(e.Error())
		return
	}
	ucEnv := imconf.GetUcEnv()
	addr := ucEnv.Host + ":" + ucEnv.Port
	conn, e := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(creds),
		grpc.WithBlock(),
	)
	if e != nil {
		err = e
		log.Info().Msg(e.Error())
		return
	}
	defer conn.Close()
	datamap := map[string]interface{}{
		"password":   pwd,
		"deviceID":   "imcenter_device",
		"platform":   "imcenter_platform",
		"deviceName": "imcenter",
		"appkey":     appkey,
	}
	var data string
	data, e = ij.Token(datamap)
	if e != nil {
		log.Info().Msg(e.Error())
		err = e
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	lc := userproto.NewUserSrvClient(conn)
	res, e := lc.Login(ctx, &userproto.UserLgnReq{
		Account: acc,
		Data:    data,
	})
	if e != nil {
		err = e
		return
	}
	return
}
