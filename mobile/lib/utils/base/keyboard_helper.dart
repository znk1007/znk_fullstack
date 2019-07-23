import 'package:flutter/material.dart';

class ScrollFocusNode extends FocusNode {
  final bool _useSystemKeyboard;
  final double _moveValue;
  ScrollFocusNode({bool useSystemKeyboard = true, double moveValue}):
  _useSystemKeyboard = useSystemKeyboard,
  _moveValue = moveValue;

  @override
  bool consumeKeyboardToken() {
    if (_useSystemKeyboard) {
      return super.consumeKeyboardToken();
    }
    return false;
  }
  // 移动值
  double get moveValue => _moveValue;
  // 是否使用系统键盘
  bool get userSystemKeyboard => _useSystemKeyboard;
}

abstract class KeyboardHelpWidget extends State<StatefulWidget> with WidgetsBindingObserver {
  // 滚动控制
  final ScrollController _controller = ScrollController();
  ScrollFocusNode _node;
  List<Widget> children();
  // 当前位置
  double _currentPosition = 0.0;

  @override
  void initState() { 
    super.initState();
    WidgetsBinding.instance.addObserver(this);
  }

  @override
  void dispose() { 
    _controller.dispose();
    WidgetsBinding.instance.removeObserver(this);
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SingleChildScrollView(
        controller: _controller,
        child: Column(
          children: children()..add(SizedBox(height: 400.0)),
        ),
      ),
    );
  }

  @override
  void didChangeMetrics() {
    if (_currentPosition != 0.0) {
      _node.unfocus();
      _reset();
    }
  }

  // 绑定输入焦点控件
  void bindInput(ScrollFocusNode node) {
    _node = node;
    _animateUp();
  }
  // 上移
  void _animateUp() {
    _controller.animateTo(
      _node.moveValue,
      duration: Duration(milliseconds: 250),
      curve: Curves.easeOut,
    ).then((val) {
      _currentPosition = _controller.offset;
    });
  }
  // 复位
  void _reset() {
    _controller.animateTo(
      0.0,
      duration: Duration(milliseconds: 250),
      curve: Curves.easeOut,
    ).then((val){
      _currentPosition = 0.0;
    });
  }

}