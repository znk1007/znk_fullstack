import 'package:flutter/material.dart';
import 'package:userAuth/auth/delegate/login_view_delegate.dart';
import 'package:userAuth/auth/resource/images/image_helper.dart';
import 'package:userAuth/auth/utils/tools/screen_helper.dart';

final double _bgImageHeight = ScreenHelper.setWidth(400);

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
            _BackgroundView(),
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
  /// 固定的背景组件
  Widget _fixedBackgroundWidget() {
    return Container(
      child: ImageHelper.load('auth_bg_image.png',
        fit: BoxFit.fitWidth, 
        height: _bgImageHeight,
        width: ScreenHelper.screenWidth,
      ),
    );
  }
  

}

class _BackgroundView extends StatefulWidget {
  _BackgroundView({Key key}) : super(key: key);

  @override
  __BackgroundViewState createState() => __BackgroundViewState();
}

class __BackgroundViewState extends State<_BackgroundView> with TickerProviderStateMixin{
  
  @override
  Widget build(BuildContext context) {
    return Container(
       child: ImageHelper.load('auth_bg_image.png',
        fit: BoxFit.fitWidth, 
        height: _bgImageHeight,
        width: ScreenHelper.screenWidth,
      ),
    );
  }
}