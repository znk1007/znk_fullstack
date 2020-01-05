import 'package:flutter/material.dart';
import 'package:userAuth/auth/delegate/login_view_delegate.dart';
import 'package:userAuth/auth/resource/images/image_helper.dart';
import 'package:userAuth/auth/utils/tools/screen_helper.dart';

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
    return Ink(
      child: InkWell(
        child: Stack(
          overflow: Overflow.visible,
          children: <Widget>[
            _fixedBackgroundWidget(),
            Container(
              child: Text('测试二'),
            ),
          ],
        ),
        onTap: () {
          print('tap view');
        },
      ),
    );
  }
  Widget _fixedBackgroundWidget() {
    return Container(
      child: ImageHelper.load('auth_bg_image.png'),
    );
  }
}