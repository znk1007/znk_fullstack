package model

//Platform 平台
type Platform string

const (
	//Web 网页
	Web Platform = "web"
	//Mobile 移动端
	Mobile Platform = "mobile"
)

//AuthInfo 验证信息
type AuthInfo struct {
	Token    string   `form:"token" json:"token" xml:"token" binding:"required"`
	Platform Platform `form:"platform" json:"platform" xml:"platform" binding:"required"`
}

//RegistInfo 注册信息
type RegistInfo struct {
	Account  string   `form:"account" json:"account" xml:"account" binding:"required"`
	Password string   `form:"password" json:"password" xml:"password" binding:"required"`
	Platform Platform `form:"platform" json:"platform" xml:"platform" binding:"required"`
}

//LoginInfo 登录信息
type LoginInfo struct {
	Account  string   `form:"account" json:"account" xml:"account" binding:"required"`
	UserID   string   `form:"userId" json:"userId" xml:"userId" binding:"required"`
	Token    string   `form:"token" json:"token" xml:"token" binding:"required"`
	Platform Platform `form:"platform" json:"platform" xml:"platform" binding:"required"`
}
