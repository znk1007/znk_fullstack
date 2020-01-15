import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:userAuth/auth/default/animation/znk_movement_view.dart';
import 'package:userAuth/auth/delegate/login_view_delegate.dart';
import 'package:userAuth/auth/resource/images/image_helper.dart';
import 'package:userAuth/auth/utils/tools/screen_helper.dart';
import 'dart:math' as math;

final double _bgImageHeight = ScreenHelper.setWidth(400);
final double _fieldHeight = ScreenHelper.setHeight(55);

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
            _AccPswInput(
              height: _fieldHeight  * 2,
            ),
          ],
        ),
        onTap: () {
          print('tap view');
          // _configCircle();
        },
      ),
    );
  }
}

class _AccPswInput extends StatefulWidget {
  final double height;
  _AccPswInput({
    Key key, 
    @required double height
  }) : height = height, 
        super(key: key);

  @override
  __AccPswInputState createState() => __AccPswInputState();
}

class __AccPswInputState extends State<_AccPswInput> {
  final TextEditingController _acController = TextEditingController();
  final TextEditingController _psdController = TextEditingController();
  final FocusNode _acFocusNode = FocusNode();
  final FocusNode _psdFocusNode = FocusNode();
  @override
  Widget build(BuildContext context) {
    return Container(
      height: widget.height,
       child: Row(
         children: <Widget>[
          //  TextField(
          //    controller: _acController,
          //    focusNode: _acFocusNode,
          //  ),
          //  TextField(
          //    controller: _psdController,
          //    focusNode: _psdFocusNode,
          //  ),
         ],
       ),
    );
  }
}

