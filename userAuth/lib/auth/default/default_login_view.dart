import 'package:flutter/material.dart';
import 'package:userAuth/auth/delegate/login_view_delegate.dart';
import 'package:userAuth/auth/resource/images/image_helper.dart';

class DefaultLoginView extends StatelessWidget implements LoginViewDelegate {
  /// 账号
  String _account;
  /// 记住密码
  bool _keepPassword;
  /// 密码
  String _password;
  /// 验证码
  String _verifyCode;


  DefaultLoginView({
    Key key,
  }):super(key: key);
  
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
    return Container(
      child: InkWell(
        child: Stack(
          children: <Widget>[

          ],
        ),
      ),
    );
  }
  Widget _backgroundWidget() {
    return Container(
      child: ImageHelper.load('auth_bg_image.png'),
    );
  }
}