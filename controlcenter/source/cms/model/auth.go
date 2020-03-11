package model

import (
	_ "github.com/dgrijalva/jwt-go"
)

//AuthInfo 验证信息
type AuthInfo struct {
	Account  string `form:"account" json:"account" xml:"account" binding:"required"`
	Token    string `form:"token" json:"token" xml:"token" binding:"required"`
	Platform string `form:"platform" json:"platform" xml:"platform" binding:"required"`
}

//RegistInfo 注册信息
type RegistInfo struct {
	Account  string `form:"account" json:"account" xml:"account" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
	Platform string `form:"platform" json:"platform" xml:"platform" binding:"required"`
}

//LoginInfo 登录信息
type LoginInfo struct {
	Account  string `form:"account" json:"account" xml:"account" binding:"required"`
	UserID   string `form:"userId" json:"userId" xml:"userId" binding:"required"`
	Token    string `form:"token" json:"token" xml:"token" binding:"required"`
	Platform string `form:"platform" json:"platform" xml:"platform" binding:"required"`
}
