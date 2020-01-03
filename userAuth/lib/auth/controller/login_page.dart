import 'package:flutter/material.dart';

import 'delegates/login_view_delegate.dart';



class DefaultLoginPage extends StatelessWidget implements LoginViewDelegate {
  /// 账号
  String _account;
  /// 记住密码
  bool _keepPassword;
  /// 密码
  String _password;
  /// 验证码
  String _verifyCode;
  
  @override
  String get account => _account;

  @override
  bool get keepPassword => _keepPassword;

  @override
  String get password => _password;

  @override
  String get verifyCode => _verifyCode;

  @override
  Widget build(BuildContext context) {
    // TODO: implement build
    throw UnimplementedError();
  }
  
}
class LoginPage<T extends LoginViewDelegate> extends StatefulWidget {
  T _value;
  LoginPage({
    Key key,
    T value,
  }) {
    _value = value;
  }

  @override
  _LoginPageState createState() => _LoginPageState();
}
class _LoginPageState extends State<LoginPage> {

  @override
  Widget build(BuildContext context) {
    widget._value.account;
    var child = (widget._value != null) ? widget._value : DefaultLoginPage();
    return Container(
       child: child,
    );
  }
}