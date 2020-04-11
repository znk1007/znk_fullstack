package usertools

import "regexp"

//VerifyEmail 正则校验邮箱
func VerifyEmail(email string) bool {
	patn := `[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(patn)
	return reg.MatchString(email)
}

//VerifyPhone 校验手机号
func VerifyPhone(phone string) bool {
	patn := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|(19[8,9]))\\d{8}$"
	reg := regexp.MustCompile(patn)
	return reg.MatchString(phone)
}

/*
用户名：  	/^[a-z0-9_-]{3,16}$/
密码：	    /^[a-z0-9_-]{6,18}$/
十六进制值：	/^#?([a-f0-9]{6}|[a-f0-9]{3})$/
电子邮箱	：  /^([a-z0-9_\.-]+)@([\da-z\.-]+)\.([a-z\.]{2,6})$/
		    /^[a-z\d]+(\.[a-z\d]+)*@([\da-z](-[\da-z])?)+(\.{1,2}[a-z]+)+$/
URL： 	    /^(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$/
IP 地址：	/((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?)/
            /^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$/
HTML 标签：	/^<([a-z]+)([^<]+)*(?:>(.*)<\/\1>|\s+\/>)$/

删除代码\\注释：      	(?<!http:|\S)//.*$
Unicode编码中的汉字范围：	/^[\u2E80-\u9FFF]+$/

*/
