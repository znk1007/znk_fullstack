package pool


/*注册事务*/
type RegisterJob interface {
	Register(account string) error
	GetRegisterVerifyCode(phone string) error
}

type UpdateJob interface {
	UpdateNickname(userId string, sessionId string, nickname string) error
}


