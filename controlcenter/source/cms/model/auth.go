package model
//AuthInfo 验证消息
type AuthInfo struct {
	Account   string `json:"account"`
	Password  string `json:"password"`
	SessionID string `json:"sessionId"`
	Platform  string `json:"platform"`
}