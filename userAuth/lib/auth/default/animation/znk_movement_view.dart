
import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';
import 'dart:math' as math;


class _CircleCurve extends Curve {
  @override 
  double transform(double t) {
    return math.sin(t * math.pi * 2);
  }
}

class MovementWidget extends StatefulWidget {
  final Widget child;
  final double offset;
  final int itemCount;
  final double maxScale;
  MovementWidget({
    Key key,
    @required this.child,
    double offset = 5,
    int itemCount = 50,
    double scale = 1.0,
  }): offset = offset, 
    itemCount = itemCount, 
    maxScale = scale,
    assert(offset != 0), 
    assert(itemCount != 0),
    assert(scale >= 1.0),
    super(key: key);

  @override
  _MovementWidgetState createState() => _MovementWidgetState();
}

class _MovementWidgetState extends State<MovementWidget> with TickerProviderStateMixin {
  AnimationController _controller;
  Animation<dynamic> _movement;
  @override
  void initState() {
    super.initState();
    _controller = AnimationController(duration: Duration(seconds: 10), vsync: this);
    _initAnimations();
    _startAnimation();
  }

  void _initAnimations() {
    List<TweenSequenceItem> items = List<TweenSequenceItem>();
    int cnt = widget.itemCount > 0 ? widget.itemCount : 10;
    double r = widget.offset > 0 ? widget.offset : 5;
    r = r / 2.0;
    double step = r / cnt.toDouble();
    double tmpStep = -r;
    // 第一象限
    for (var i = 0; i < cnt; i++) {
      double x1 = tmpStep;
      double y1 = -_y(x1, r);
      tmpStep+=step;
      double x2 = tmpStep;
      double y2 = -_y(x2, r);
      TweenSequenceItem item = TweenSequenceItem(
        tween: Tween(
          begin: Offset(x1, y1),
          end: Offset(x2, y2)
        ),
        weight: 1,
      );
      items.add(item);
    }
    //第二象限
    tmpStep = 0;
    for (var i = 0; i < cnt; i++) {
      double x1 = tmpStep;
      double y1 = -_y(x1, r);
      tmpStep+=step;
      double x2 = tmpStep;
      double y2 = -_y(x2, r);
      TweenSequenceItem item = TweenSequenceItem(
        tween: Tween(
          begin: Offset(x1, y1),
          end: Offset(x2, y2)
        ),
        weight: 1,
      );
      items.add(item);
    }
    // 第三象限
    tmpStep = r;
    for (var i = 0; i < cnt; i++) {
      double x1 = tmpStep;
      double y1 = _y(x1, r);
      tmpStep-=step;
      double x2 = tmpStep;
      double y2 = _y(x2, r);
      TweenSequenceItem item = TweenSequenceItem(
        tween: Tween(
          begin: Offset(x1, y1),
          end: Offset(x2, y2)
        ),
        weight: 1,
      );
      items.add(item);
    }

    // 第四象限
    tmpStep = 0;
    for (var i = 0; i < cnt; i++) {
      double x1 = tmpStep;
      double y1 = _y(x1, r);
      tmpStep-=step;
      double x2 = tmpStep;
      double y2 = _y(x2, r);
      TweenSequenceItem item = TweenSequenceItem(
        tween: Tween(
          begin: Offset(x1, y1),
          end: Offset(x2, y2)
        ),
        weight: 1,
      );
      items.add(item);
    }
    

    _movement = TweenSequence(items).animate(
      CurvedAnimation(
        parent: _controller,
        curve: Interval(
          0.1, 
          0.5, 
          curve: Curves.linear
        ),
      ),
    )..addListener((){
      setState(() {
        
      });
    })..addStatusListener((status){
      print('current move status: $status');
      // if (status == AnimationStatus.completed) {
      //   _controller.forward();
      // }
    });
    // _controller.forward();
  }

  Future _startAnimation() async {
    try {
      await _controller?.repeat();
    } catch (e) {
      if (e is TickerCanceled) {
        print('ticker canceled');
      } else {
        print('animation failed $e');
      }
    }
  }
  /* 根据半径，x坐标，计算圆的y坐标 */
  double _y(double x, double r) {
    print('x $x');
    return math.sqrt(r * r - x * x);
  }

  @override
  void dispose() {
    _controller?.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {    
    // _configCircle();
    if (_movement == null) {
      return Container();
    }
    return ClipRect(
      child: Transform.scale(
        scale: widget.maxScale,
        child: Transform.translate(
          offset: _movement?.value,
          child: widget.child,
        ),
      ),
    );
  }
}

/*
/* 根据半径，x坐标，计算圆的y坐标 */
  double _y(double x, double r) {
    print('x $x');
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
*/
/*
目录

动画相关主要对象
缩放动画
非线性缩放动画
淡入淡出
非线性淡入淡出
平移动画
非线性平移动画
动画相关主要对象
Animation：可分为线性动画、非线性动画、步进函数动画或其它动画。通过 addListener 方法可以添加监听器，每当动画帧发生改变时均会调用，一般会配合 setState 方法用作UI重建。通过 addStatusListener 方法可添加状态改变监听器，如：动画开始、动画结束等
AnimationController：动画控制器，动画的开始、结束、停止、反向均由它控制，方法对应为：forward、stop、reverse
Curve：可使用此对象将动画设置为为匀速、加速或先加速后减速等。Curve 可以为线性或非线性
缩放动画
import 'package:flutter/material.dart';

/**
 * @des Animation Zoom
 * @author liyongli 20190516
 * */
class AnimationZoom extends StatefulWidget{

  @override
  State<StatefulWidget> createState() {
    return new _AnimationZoomState();
  }

}

/**
 * @des Animation Zoom State
 * @author liyongli 20190516
 * */
class _AnimationZoomState extends State<AnimationZoom> with TickerProviderStateMixin{

  // 放大动画
  Animation<double> animationEnlarge;
  // 放大动画控制器
  AnimationController enlargeAnimationController;
  // 缩小动画
  Animation<double> animationNarrow;
  // 缩小动画控制器
  AnimationController narrowAnimationController;
  
  @override
  void initState() {
    super.initState();
    
    // 定义动画持续时长
    enlargeAnimationController = new AnimationController(vsync: this,duration:Duration(seconds: 3) );
    narrowAnimationController = new AnimationController(vsync: this, duration: Duration(seconds: 3));

    // 定义缩放动画范围
    animationEnlarge = new Tween(begin: 10.0, end: 150.0).animate(enlargeAnimationController)
      ..addListener((){
        setState(() {});
      })
      ..addStatusListener((status){
        if(status == AnimationStatus.completed){
          narrowAnimationController.forward();
        }
      });

    animationNarrow = new Tween(begin: 150.0, end: 10.0).animate(narrowAnimationController)
      ..addListener((){
        setState(() {});
      });

    // 开启动画
    enlargeAnimationController.forward();

  }

  @override
  void dispose() {
    // 释放资源
    enlargeAnimationController.dispose();
    narrowAnimationController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(

        appBar: AppBar(
          title: Text("AnimationEnlarge"),
        ),

        body: Container(
          alignment: Alignment.center,
          child: Column(
            children: <Widget>[

              // 放大动画
              Padding(
                padding: EdgeInsets.all(10.0),
                child: Image.asset("images/image_widget_test.jpg",
                  width: animationEnlarge.value,
                  height: animationEnlarge.value,
                ),
              ),

              // 缩小动画
              Padding(
                padding: EdgeInsets.all(10.0),
                child: Image.asset("images/image_widget_test.jpg",
                  width: animationNarrow.value,
                  height: animationNarrow.value,
                ),
              ),

            ],
          ),
        )
      ),
    );
  }

}
非线性缩放动画
import 'package:flutter/material.dart';

/**
 * @des Animation Curved
 * @author liyongli 20190516
 * */
class AnimationCurved extends StatefulWidget{

  @override
  State<StatefulWidget> createState() {
    return new _AnimationCurvedState();
  }

}

/**
 * @des Animation Curved State
 * @author liyongli 20190516
 * */
class _AnimationCurvedState extends State<AnimationCurved> with TickerProviderStateMixin{

  // 放大动画
  Animation<double> animationCurved;
  // 放大动画控制器
  AnimationController animationCurvedController;

  @override
  void initState() {
    super.initState();

    // 定义动画持续时长
    animationCurvedController = new AnimationController(vsync: this,duration:Duration(seconds: 3) );

    // 定义具体曲线
    CurvedAnimation curve = new CurvedAnimation(parent: animationCurvedController, curve: Curves.elasticOut);
    
    // 定义缩放动画范围
    animationCurved = new Tween(begin: 100.0, end: 350.0).animate(curve)
      ..addListener((){
        setState(() {});
      });

    // 开启动画
    animationCurvedController.forward();

  }

  @override
  void dispose() {
    // 释放资源
    animationCurvedController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(

        appBar: AppBar(
          title: Text("AnimationEnlarge"),
        ),

        body: Container(
          alignment: Alignment.center,
          child: Column(
            children: <Widget>[
              Padding(
                padding: EdgeInsets.all(10.0),
                child: Image.asset("images/image_widget_test.jpg",
                  width: animationCurved.value,
                  height: animationCurved.value,
                ),
              ),
            ],
          ),
        )
      ),
    );
  }

}
淡入淡出
import 'package:flutter/material.dart';

/**
 * @des Animation Fade
 * @author liyongli 20190517
 * */
class AnimationFade extends StatefulWidget{

  @override
  State<StatefulWidget> createState() {
    return new _AnimationFadeState();
  }

}

/**
 * @des Animation Fade State
 * @author liyongli 20190517
 * */
class _AnimationFadeState extends State<AnimationFade> with TickerProviderStateMixin{

  // 初始 animationType 为 1.0 为可见状态，为 0.0 时不可见
  double animationType = 1.0;
  // 动画持续时长
  int animationSeconds = 2;
  
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(

        appBar: AppBar(
          title: Text("AnimationFade"),
        ),

        body: Container(
          alignment: Alignment.center,
          child: Column(
            children: <Widget>[

              new AnimatedOpacity(
                  opacity: animationType,
                  duration: new Duration(seconds: animationSeconds),
                  child:new Container(
                    child:new Text("轻轻地来轻轻地走轻轻地来轻轻地走轻轻地来轻轻地走轻轻地来轻轻地走轻轻地来轻轻地走轻轻地来轻轻地走轻轻地来轻轻地走轻轻地来轻轻地走轻轻地来轻轻地走轻轻地来轻轻地走") ,)
              ),

              new RaisedButton(
                child:new Container(
                  child: new Text("淡入 and 淡出"),
                ) ,
                onPressed: _changeAnimationType,//添加点击事件
              ),
            ],
          ),
        )
      ),
    );
  }

  // 修改文字显示状态（赋值倒置）
  _changeAnimationType() {
    setState(
       () => animationType = animationType == 0 ? 1.0 : 0.0
    );
  }

}
非线性淡入淡出
import 'package:flutter/material.dart';

/**
 * @des Animation Fade
 * @author liyongli 20190517
 * */
class AnimationFade extends StatefulWidget{

  @override
  State<StatefulWidget> createState() {
    return new _AnimationFadeState();
  }

}

/**
 * @des Animation Fade State
 * @author liyongli 20190517
 * */
class _AnimationFadeState extends State<AnimationFade> with TickerProviderStateMixin{

  // 初始animationType为1.0为可见状态，为0.0时不可见
  double animationType = 1.0;
  // 动画持续时长
  int animationSeconds = 2;
  
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(

        appBar: AppBar(
          title: Text("AnimationFade"),
        ),

        body: Container(
          alignment: Alignment.center,
          child: Column(
            children: <Widget>[

              new AnimatedOpacity(
                  opacity: animationType,
                  curve: Curves.elasticInOut, // 这里是设置非线性动画的关键
                  duration: new Duration(seconds: animationSeconds),
                  child:new Container(
                    child:new Text("轻轻地来轻轻地走轻轻地来轻轻地走轻轻地来轻轻地走轻轻地来轻轻地走轻轻地来轻轻地走轻轻地来轻轻地走轻轻地来轻轻地走轻轻地来轻轻地走轻轻地来轻轻地走轻轻地来轻轻地走") ,)
              ),
              new RaisedButton(
                child:new Container(
                  child: new Text("淡入 and 淡出"),
                ) ,
                onPressed: _changeAnimationType,//添加点击事件
              ),
            ],
          ),
        )
      ),
    );
  }

  // 修改文字显示状态
  _changeAnimationType() {
    setState(
       () => animationType = animationType == 0 ? 1.0 : 0.0
    );
  }

}
平移动画
import 'package:flutter/material.dart';

/**
 * @des Animation XY
 * @author liyongli 20190517
 * */
class AnimationXY extends StatefulWidget{

  @override
  State<StatefulWidget> createState() {
    return new _AnimationXYState();
  }

}

/**
 * @des Animation XY State
 * @author liyongli 20190517
 * */
class _AnimationXYState extends State<AnimationXY> with TickerProviderStateMixin{

  // 左右移动动画
  Animation<EdgeInsets> animationX;
  // 左右移动动画控制器
  AnimationController xAnimationController;
  // 上下移动动画
  Animation<EdgeInsets> animationY;
  // 上下移动动画控制器
  AnimationController yAnimationController;

  @override
  void initState() {
    super.initState();

    // 定义动画持续时长
    xAnimationController = new AnimationController(vsync: this,duration:Duration(seconds: 3) );
    yAnimationController = new AnimationController(vsync: this, duration: Duration(seconds: 3));

    // 定义平移动画范围
    animationX = new EdgeInsetsTween(begin: EdgeInsets.only(left: 0.0), end: EdgeInsets.only(left: 100.0)).animate(xAnimationController)
      ..addListener((){
        setState(() {});
      });

    animationY = new EdgeInsetsTween(begin: EdgeInsets.only(top: 0.0), end: EdgeInsets.only(top: 100.0)).animate(yAnimationController)
      ..addListener((){
        setState(() {});
      });

    // 开启动画
    xAnimationController.forward();
    yAnimationController.forward();

  }

  @override
  void dispose() {
    // 释放资源
    xAnimationController.dispose();
    yAnimationController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(

          appBar: AppBar(
            title: Text("animationX"),
          ),

          body: Container(
            alignment: Alignment.center,
            child: Column(
              children: <Widget>[

                // 左右移动
                Padding(
                  padding: animationX.value,
                  child: Image.asset("images/image_widget_test.jpg",
                  ),
                ),

                // 上下移动
                Padding(
                  padding: animationY.value,
                  child: Image.asset("images/image_widget_test.jpg",
                  ),
                ),

              ],
            ),
          )
      ),
    );
  }

}
非线性平移动画
import 'package:flutter/material.dart';

/**
 * @des Animation XY
 * @author liyongli 20190517
 * */
class AnimationXY extends StatefulWidget{

  @override
  State<StatefulWidget> createState() {
    return new _AnimationXYState();
  }

}

/**
 * @des Animation XY State
 * @author liyongli 20190517
 * */
class _AnimationXYState extends State<AnimationXY> with TickerProviderStateMixin{

  // 左右移动动画
  Animation<EdgeInsets> animationX;
  // 左右移动动画控制器
  AnimationController xAnimationController;
  // 上下移动动画
  Animation<EdgeInsets> animationY;
  // 上下移动动画控制器
  AnimationController yAnimationController;

  @override
  void initState() {
    super.initState();

    // 定义动画持续时长（使用 CurvedAnimation 设置非线性动画）
    xAnimationController = new AnimationController(vsync: this,duration:Duration(seconds: 3) );
    yAnimationController = new AnimationController(vsync: this, duration: Duration(seconds: 3));

    // 定义平移动画范围
    animationX = new EdgeInsetsTween(begin: EdgeInsets.only(left: 0.0), end: EdgeInsets.only(left: 100.0)).animate(CurvedAnimation(parent: xAnimationController, curve: Interval(0.1, 0.6)))
      ..addListener((){
        setState(() {});
      });

    animationY = new EdgeInsetsTween(begin: EdgeInsets.only(top: 0.0), end: EdgeInsets.only(top: 100.0)).animate(CurvedAnimation(parent: yAnimationController, curve: Interval(0.1, 0.6)))
      ..addListener((){
        setState(() {});
      });

    // 开启动画
    xAnimationController.forward();
    yAnimationController.forward();

  }

  @override
  void dispose() {
    // 释放资源
    xAnimationController.dispose();
    yAnimationController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(

          appBar: AppBar(
            title: Text("animationX"),
          ),

          body: Container(
            alignment: Alignment.center,
            child: Column(
              children: <Widget>[

                // 左右移动
                Padding(
                  padding: animationX.value,
                  child: Image.asset("images/image_widget_test.jpg",
                  ),
                ),

                // 上下移动
                Padding(
                  padding: animationY.value,
                  child: Image.asset("images/image_widget_test.jpg",
                  ),
                ),

              ],
            ),
          )
      ),
    );
  }

}

*/