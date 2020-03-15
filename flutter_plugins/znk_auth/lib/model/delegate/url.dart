abstract class UrlDelegate {
  /* 请求地址域名 */
  String get baseUrl;
  /* 获取验证码 */
  String get getVerifyCode;
  /* 校验验证码 */
  String get verifyCode;
  /* 获取图形验证码 */
  String get getVerifyGraphicalCode;
  /* 验证图形验证码 */
  String get verifyGraphicalCode;
  /* 登录 */
  String get login;
  /* 忘记密码 */
  String get forgetPassword;
  /* 注册 */
  String get regist;
}