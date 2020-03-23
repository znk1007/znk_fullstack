package middleware

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	usertoken "github.com/znk_fullstack/server/usercenter/viewmodel/token"
	"google.golang.org/grpc/metadata"
)

//CheckToken 校验token
func CheckToken(ctx context.Context, checkTS bool) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Info().Msg("check token failed")
		return errors.New("check token failed")
	}
	var sign string
	if val, ok := md["sign"]; ok {
		if len(val) > 0 {
			sign = val[0]
		}
	}
	if len(sign) == 0 {
		log.Info().Msg("miss param `sign` or `sign` is empty")
		return errors.New("miss param `sign` or `sign` is empty")
	}
	tk, err := usertoken.ParseToken(sign)
	if err != nil {
		log.Info().Msg(err.Error())
		return err
	}
	key, ok := tk["appkey"]
	appkey := key.(string)
	if !ok {
		log.Info().Msg("miss param `appkey`")
		return errors.New("miss param `appkey`")
	}
	if len(appkey) == 0 {
		log.Info().Msg("appkey is empty")
		return errors.New("appkey is empty")
	}
	if appkey != usertoken.GetSecurityKeyString() {
		log.Info().Msg("appkey is wrong")
		return errors.New("appkey is wrong")
	}
	if checkTS {
		var ts interface{}
		ts, ok = tk["timestamp"]
		if !ok {
			log.Info().Msg("miss param `timestamp`")
			return errors.New("miss param `timestamp`")
		}
		var timestamp int64
		timestamp, err = strconv.ParseInt(ts.(string), 10, 64)
		now := time.Now().Unix()
		if now-timestamp > usertoken.ExpiredDuration().Microseconds() {
			log.Info().Msg("time expired")
			return errors.New("time expired")
		}
	}

	return nil
}
