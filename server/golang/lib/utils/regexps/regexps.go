package regexputils

import "regexp"

// VerifyEmail 校验邮箱格式
func VerifyEmail(email string) bool {
	// pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w)*`
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// VerifyPhone 校验手机号
func VerifyPhone(phone string) bool {
	pattern := `^[1][3,5,8,9][0-9]{9}$|^[9][2,8][0-9]{9}$|^14[5|6|7|8|9][0-9]{8}$|^16[1|2|4|5|6|7][0-9]{8}$|^17[0-8][0-9]{8}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(phone)
}
