import 'dart:ui';
import 'dart:ui';

import 'package:flutter/material.dart';

enum HUDType {
  text,
  progress,
  mixed,
}

enum HUDPosition {
  top,
  center,
  bottom
}

class HUD extends StatefulWidget {

  HUD({
    Key key,
    HUDType type,
    HUDPosition position,
    Color boxBackgroundColor,
    double borderRadius,
    bool isFullscreen,
    Color progressColor,
    String text,
    Color textColor,
    Size boxSize,
    int duration,
    Widget child,
  }) {
    
    this._hudType = type ?? HUDType.mixed;
    this._hudPosition = position ?? HUDPosition.center;
    this._boxBackgroundColor = boxBackgroundColor ?? Color.fromARGB(225, 0, 0, 0);
    this._borderRadius = borderRadius ?? 5.0;
    this._isFullscreen = isFullscreen ?? true;
    this._progressColor = progressColor ?? Colors.lightBlue;
    this._text = text ?? 'loading...';
    this._textColor = textColor ?? Colors.white;
    this._boxSize = boxSize ?? Size(0, 0);
    this._duration = duration ?? 0;
    this._child = child;
    if (this._hudType == HUDType.mixed) {
      this._textColor = this._progressColor;
      this._boxBackgroundColor = Colors.white;
    }
  }
  HUDType _hudType;

  HUDPosition _hudPosition;
  // 背景颜色
  Color _boxBackgroundColor;
  // 圆角
  double _borderRadius;
  // 盒子大小
  Size _boxSize;
  // 是否满屏
  bool _isFullscreen;
  // 进度条颜色
  Color _progressColor;
  // 显示文本
  String _text;
  // 文本颜色
  Color _textColor;
  // duration秒后隐藏，当duration>0有效
  int _duration;
  // 子组件
  Widget _child;
  // 是否已显示
  bool _isShowing = false;
  // 内容区域宽高
  Size _innerSize = Size(100, 100);
  // 状态
  _HUDState state;
  // 行数
  int _numberOfLines = 1;

  @override
  _HUDState createState() {
    state = _HUDState();
    return state;
  }
}

class _HUDState extends State<HUD> {

  // 遮罩视图
  Widget _fullscreenCorver() {
    if (widget._isFullscreen == true) {
      
      return Container(
        width: _windowSize.width,
        height: _windowSize.height,
        color: Color.fromRGBO(0, 0, 0, 0),
      );
    } else {
      return Container();
    }
  }
  // 内容区域视图
  Widget _contentView() {
    var containerMargin = EdgeInsets.all(0);
    if (widget._boxSize == Size.zero) {
      switch (widget._hudPosition) {
        case HUDPosition.center:
          {
            containerMargin = EdgeInsets.fromLTRB((_windowSize.width - widget._innerSize.width) / 2, 
                        _windowSize.height / 2 - widget._innerSize.height , 0, 0);
          }
          break;
        case HUDPosition.bottom:
          {
            containerMargin = EdgeInsets.fromLTRB((_windowSize.width - widget._innerSize.width) / 2, 
                        _windowSize.height - 2 * widget._innerSize.height, 0, 0);
          }
          break;
        case HUDPosition.top:
        {
          containerMargin = EdgeInsets.fromLTRB((_windowSize.width - widget._innerSize.width) / 2, 
                      0, 0, 0);
        }
          break;
        default:
      }
      return Container(
        width: widget._innerSize.width,
        height: widget._innerSize.height,
        margin: containerMargin,
        child: _contentChildrenView(),
      );
    } else {
      switch (widget._hudPosition) {
        case HUDPosition.center:
          {
            containerMargin = EdgeInsets.fromLTRB((_windowSize.width - widget._boxSize.width) / 2, 
                        _windowSize.height / 2 - widget._boxSize.height , 0, 0);
          }
          break;
        case HUDPosition.bottom:
          {
            containerMargin = EdgeInsets.fromLTRB((_windowSize.width - widget._boxSize.width) / 2, 
                        _windowSize.height - 2 * widget._boxSize.height, 0, 0);
          }
          break;
        case HUDPosition.top:
        {
          containerMargin = EdgeInsets.fromLTRB((_windowSize.width - widget._boxSize.width) / 2, 
                      0, 0, 0);
        }
          break;
        default:
      }
      return Container(
        width: widget._boxSize.width,
        height: widget._boxSize.height,
        margin: containerMargin,
        child: _contentChildrenView(),
        decoration: BoxDecoration(
          color: widget._boxBackgroundColor,
          borderRadius: BorderRadius.all(Radius.circular(widget._borderRadius)),
        ),
      );
    }
  }

  Widget _contentChildrenView() {
    switch (widget._hudType) {
      case HUDType.text:
        {
          return Container(
            alignment: Alignment.center,
            child: Text(
              widget._text,
              maxLines: widget._numberOfLines,
              overflow: TextOverflow.ellipsis,
              style: TextStyle(
                color: widget._textColor,
              ),
            ),
          );
        }
        break;
      case HUDType.progress:
        {
          return _progressChild();
        }
        break;
      case HUDType.mixed:
        {
          return Container(
            alignment: Alignment.center,
            child: Column(
              children: <Widget>[
                _progressChild(),
                Container(
                  margin: EdgeInsets.fromLTRB(0, 15, 0, 0),
                  child: Text(
                    widget._text,
                    maxLines: widget._numberOfLines,
                    style: TextStyle(
                      color: widget._textColor,
                    ),
                  ),
                ),
              ],
            ),
          );
        }
        break;
      default:
        return Container();
    }
  }

  Size _windowSize = Size(0, 0);
  bool _isVisible = false;
  @override
  void initState() {
    super.initState();
    _windowSize = MediaQueryData.fromWindow(window).size;
  }
  
  // 进度条
  Widget _progressChild() {
    return CircularProgressIndicator(
      strokeWidth: 6,
      valueColor: AlwaysStoppedAnimation(widget._progressColor),
    );
  }
  // 显示文本
  Future showText({String content}) async {
    await show(type: HUDType.text, content: content);
  }
  // 显示进度条
  Future showProgress() async {
    await show(type: HUDType.progress);
  }
  // 显示混合
  Future showMixed({String content}) async {
    await show(type: HUDType.mixed, content: content);
  } 

  // 显示
  Future show({@required HUDType type, String content}) async {
    if (widget._isShowing == true && type == widget._hudType && widget._text == content) {
      return;
    }
    this.hide();
    widget._isShowing = true;
    widget._text = content ?? 'loading';
    widget._hudType = type;
    final txt = Runes(content);
    widget._numberOfLines = txt.length > 10 ? 2 : 1;
    switch (type) {
      case HUDType.text:
      {
        widget._boxSize = Size(180, 40);
        widget._boxBackgroundColor = Color.fromARGB(225, 0, 0, 0);
        widget._duration = 3;
        widget._textColor = Colors.white;
      }
        break;
      case HUDType.progress:
      {
        widget._boxSize = Size(50, 50);
        widget._boxBackgroundColor = Color.fromRGBO(0, 0, 0, 0);
      }
        break;
      case HUDType.mixed:
      {
        widget._boxSize = Size(100, 100);
        widget._textColor = widget._progressColor;
        widget._boxBackgroundColor = Color.fromARGB(0, 0, 0, 0);
      }
        break;
      default:
    }
    
    setState(() {
      _isVisible = true;
    });
    if (widget._duration > 0) {
      await Future.delayed(Duration(seconds: widget._duration));
      this.hide();
    }
  }
  // 隐藏
  void hide() {
    widget._isShowing = false;
    setState(() {
      _isVisible = false;
    });
  }

  @override
  Widget build(BuildContext context) {
    if (_isVisible == true) {
      return Stack(
        children: <Widget>[
          _fullscreenCorver(),
          _contentView(),
        ],
      );
    } else {
      return Container();
    }
    
  }
}



