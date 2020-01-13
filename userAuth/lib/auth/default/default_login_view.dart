import 'package:flutter/material.dart';
import 'package:userAuth/auth/default/animation/znk_movement_view.dart';
import 'package:userAuth/auth/delegate/login_view_delegate.dart';
import 'package:userAuth/auth/resource/images/image_helper.dart';
import 'package:userAuth/auth/utils/tools/screen_helper.dart';
import 'dart:math' as math;

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
            MovementWidget(
              scale: 1.5,
              child: ImageHelper.load(
                'auth_bg_image.png',
                fit: BoxFit.fill,
                height: _bgImageHeight,
                width: ScreenHelper.screenWidth,
              ),
            ),
            Container(
              child: Text('测试二'),
            ),
          ],
        ),
        onTap: () {
          print('tap view');
          _configCircle();
        },
      ),
    );
  }

  /* 根据半径，x坐标，计算圆的y坐标 */
  double _y(double x, double r) {
    // print('x $x');
    return math.sqrt(r * r - x * x);
  }

  void _configCircle() {
    int cnt = 10;
    double r = 5 / 2.0;
    double step = r / cnt.toDouble();
    // 第一象限
    double tmpStep = -r;
    for (var i = 0; i < cnt; i++) {
      double x1 = tmpStep;
      double y1 = -_y(x1, r);
      print('x1 = $x1 and y1 = $y1');
      tmpStep+=step;
      double x2 = tmpStep;
      double y2 = -_y(x2, r);
      print('x2 = $x2 y2 == $y2');
    }
    //第二象限
    tmpStep = 0;
    for (var i = 0; i < cnt; i++) {
      double x1 = tmpStep;
      double y1 = -_y(x1, r);
      print('x1 = $x1 and y1 = $y1');
      tmpStep+=step;
      double x2 = tmpStep;
      double y2 = -_y(x2, r);
      print('x2 = $x2 y2 == $y2');
    }
    // 第三象限
    tmpStep = r;
    for (var i = 0; i < cnt; i++) {
      double x1 = tmpStep;
      double y1 = _y(x1, r);
      print('x1 = $x1 and y1 = $y1');
      tmpStep-=step;
      double x2 = tmpStep;
      double y2 = _y(x2, r);
      print('x2 = $x2 y2 == $y2');
    }

    // 第四象限
    tmpStep = 0;
    for (var i = 0; i < cnt; i++) {
      double x1 = tmpStep;
      double y1 = _y(x1, r);
      print('x1 = $x1 and y1 = $y1');
      tmpStep-=step;
      double x2 = tmpStep;
      double y2 = _y(x2, r);
      print('x2 = $x2 y2 == $y2');
    }
  }
}

