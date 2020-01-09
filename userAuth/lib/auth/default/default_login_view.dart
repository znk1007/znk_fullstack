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
  AnimationController _controller;
  Animation<dynamic> _movement;
  @override
  void initState() {
    super.initState();
    initController();
    initAnimation();
    startAnimate();
    
  }
  /// 初始化动画控制器
  void initController() {
    _controller = AnimationController(duration: Duration(seconds: 1), vsync: this);
  }
  /// 初始化动画
  void initAnimation() {
    _movement = TweenSequence(configItems()).animate(
      CurvedAnimation(
        parent: _controller, 
        curve: Interval(
          0.1, 
          0.5,
          curve: Curves.linear
        ),
      ),
    )
    ..addListener(() {

    })
    ..addStatusListener((status) {

    });
  }
  /// 配置items
  List<TweenSequenceItem> configItems() {
    List<TweenSequenceItem> items = [];
    double idx = 0;
    double step = 0.125;
    while (idx <= 1) {
      print("idx == $idx");
      print("idx reverse == ${-1-idx}");
    //   TweenSequenceItem item = TweenSequenceItem(
    //   tween: EdgeInsetsTween(
    //     begin: EdgeInsets.only(left: idx, top: -1 - idx),
    //     end: EdgeInsets.only(left: idx + step, top: -0.125),
    //   ),
    //   weight: 1,
    // );
    // items.add(item);
      idx+=step;
    }
    TweenSequenceItem item = TweenSequenceItem(
      tween: EdgeInsetsTween(
        begin: EdgeInsets.only(left: 0, top: 0),
        end: EdgeInsets.only(left: 0.125, top: 0),
      ),
      weight: 1,
    );
    items.add(item);

    return items;
  }
  /// 开始动画
  Future startAnimate() async {
    try {
      await _controller.repeat();
    } catch (e) {
      if (e is TickerCanceled) {
        print('ticker canceled');
      } else {
        print('animation failed $e');
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    if (_movement == null) {
      return Container();
    }
    return Container(
      child: ImageHelper.load(
          'auth_bg_image.png',
          fit: BoxFit.fill,
          height: _bgImageHeight,
          width: ScreenHelper.screenWidth,
        ),
      alignment: Alignment.topCenter,
      // padding: _movement.value,
    );
  }
}