import 'package:flutter/material.dart';
/// 获取验证码回调
typedef GetVerifyCodeAction = Function(String phone);
/// 登录回调
typedef LoginAction = Function(String account, String password);

abstract class LoginViewDelegate extends Widget {
  /// 账号
  String get account;
  /// 密码
  String get password;
  /// 验证密码
  String get verifyCode;
  /// 记住密码
  bool get keepPassword;
  /// 获取验证码
  // void getVerifyCode(GetVerifyCodeAction action);
  // /// 登录
  // void loginAction(LoginAction action);
  
}