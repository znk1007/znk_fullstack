package middleware

import (
	"context"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	usertoken "github.com/znk_fullstack/server/usercenter/viewmodel/token"
	"google.golang.org/grpc/metadata"
)

//CheckToken 校验token
func CheckToken(ctx context.Context, checkTS bool) bool {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Info().Msg("check token failed")
		return false
	}
	var sign string
	if val, ok := md["sign"]; ok {
		if len(val) > 0 {
			sign = val[0]
		}
	}
	if len(sign) == 0 {
		log.Info().Msg("miss param `sign` or `sign` is empty")
		return false
	}
	tk, err := usertoken.ParseToken(sign)
	if err != nil {
		log.Info().Msg(err.Error())
		return false
	}
	key, ok := tk["appkey"]
	appkey := key.(string)
	if !ok {
		log.Info().Msg("miss param `appkey`")
		return false
	}
	if len(appkey) == 0 {
		log.Info().Msg("appkey is empty")
		return false
	}
	if appkey != usertoken.GetSecurityKeyString() {
		log.Info().Msg("appkey is wrong")
		return false
	}
	if checkTS {
		var ts interface{}
		ts, ok = tk["timestamp"]
		if !ok {
			log.Info().Msg("miss param `timestamp`")
			return false
		}
		var timestamp int64
		timestamp, err = strconv.ParseInt(ts.(string), 10, 64)
		now := time.Now().Unix()
		if timestamp {

		}
	}

	return true
}
