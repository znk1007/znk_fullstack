package security

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"time"

	"google.golang.org/grpc/metadata"
)

// Token 安全校验
type Token struct {
	appKey    string
	appSecret string
}

// Check 校验token
func (t *Token) Check(ctx context.Context) bool {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false
	}
	var (
		appKey    string
		appSecret string
	)
	if value, ok := md["app_key"]; ok {
		appKey = value[0]
	}
	if value, ok := md["app_secret"]; ok {
		appSecret = value[0]
	}
	if appKey != t.GetAppKey() || appSecret != t.GetAppSecret() {
		return false
	}
	return true
}

// GetAppKey 获取密钥
func (t *Token) GetAppKey() string {
	return "znk_project-item=20"
}

// GetAppSecret 获取安全码
func (t *Token) GetAppSecret() string {
	return "19911007"
}

// GenterateSessionID 生成会话id
func GenterateSessionID() string {
	curTime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(curTime, 10))
	return fmt.Sprintf("%x", h.Sum(nil))
}
