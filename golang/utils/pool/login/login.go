package login

/*登录事务接口*/
type Job interface {
	/// 账号
	Account() string
	/// 验证码
	Code() string
	/// 密码
	Password() string
}

/*登录事务*/
type Login struct {
	Account string
	VerifyCode string
	Password string
}

func (l Login)Login(account string, password string, verifyCode string) error {
	return nil
}
func (l Login)GetLoginVerifyCode(phone string) error {
	return nil
}