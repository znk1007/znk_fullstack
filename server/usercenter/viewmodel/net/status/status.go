package netstatus

const (
	//RejectDevice 拒绝使用该设备登录
	RejectDevice = 1001
	//SessionInvalidate 登录session失效
	SessionInvalidate = 1002
	//UserInactive 用户被禁用
	UserInactive = 1003
	//UserLogout 用户退出登录
	UserLogout = 1004
	//UserNotRegisted 用户未注册
	UserNotRegisted = 1005
	//RequestFrequence 请求过于频繁
	RequestFrequence = 1006
	//PasswordNoMatch 密码不匹配
	PasswordNoMatch = 1007
)
