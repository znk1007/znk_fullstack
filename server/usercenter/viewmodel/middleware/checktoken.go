package usermiddleware

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	userjwt "github.com/znk_fullstack/server/usercenter/viewmodel/jwt"
	"google.golang.org/grpc/metadata"

	usercrypto "github.com/znk_fullstack/server/usercenter/viewmodel/crypto"
)

//CheckToken 校验token
func CheckToken(ctx context.Context, checkTS bool) (map[string]interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Info().Msg("check token failed")
		return nil, errors.New("check token failed")
	}
	var sign string
	if val, ok := md["sign"]; ok {
		if len(val) > 0 {
			sign = val[0]
		}
	}
	if len(sign) == 0 {
		log.Info().Msg("miss param `sign` or `sign` is empty")
		return nil, errors.New("miss param `sign` or `sign` is empty")
	}
	tk, err := userjwt.ParseToken(sign)
	if err != nil {
		log.Info().Msg(err.Error())
		return nil, err
	}
	key, ok := tk["appkey"]
	appkey := key.(string)
	if !ok {
		log.Info().Msg("miss param `appkey`")
		return nil, errors.New("miss param `appkey`")
	}
	if len(appkey) == 0 {
		log.Info().Msg("appkey is empty")
		return nil, errors.New("appkey is empty")
	}
	if appkey != usercrypto.GetSecurityKeyString() {
		log.Info().Msg("appkey is wrong")
		return nil, errors.New("appkey is wrong")
	}
	if checkTS {
		var ts interface{}
		ts, ok = tk["timestamp"]
		if !ok {
			log.Info().Msg("miss param `timestamp`")
			return nil, errors.New("miss param `timestamp`")
		}
		var timestamp int64
		timestamp, err = strconv.ParseInt(ts.(string), 10, 64)
		now := time.Now().Unix()
		if now-timestamp > ExpiredDuration().Microseconds() {
			log.Info().Msg("time expired")
			return nil, errors.New("time expired")
		}
	}
	return tk, nil
}

//ExpiredDuration 两分钟响应超时失效
func ExpiredDuration() time.Duration {
	return time.Duration(time.Minute * 2)
}
