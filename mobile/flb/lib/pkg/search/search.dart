import 'package:flutter/material.dart';

//获取占位回调
typedef ZNKSearchPlaceholderFunction = Function(String placeholder);
//获取输入内容回调
typedef ZNKSearchInputFunction = Function(String input);

class ZNKSearchStyle {
  //背景颜色
  final Color backgroudColor;
  //倒角
  final double cornerRadius;
  //宽度
  final double width;
  //高度
  final double height;
  //边距
  final EdgeInsets margin;
  //textColor 文字颜色
  final Color textColor;
  //是否可用、可输入
  final bool enabled;
  //初始化
  ZNKSearchStyle(
      {this.backgroudColor = Colors.red,
      this.cornerRadius = 6.0,
      this.width = 200,
      this.height = 45,
      this.enabled = true,
      this.textColor = Colors.grey,
      this.margin = EdgeInsets.zero});
}

class ZNKSearchView extends StatefulWidget {
  //搜索样式
  ZNKSearchStyle _style;
  final Widget child;
  ZNKSearchView({Key key, ZNKSearchStyle style, this.child}) {
    _style = style ?? ZNKSearchStyle();
  }

  final _ZNKSearchViewState state = _ZNKSearchViewState();

  @override
  _ZNKSearchViewState createState() => state;
}

class _ZNKSearchViewState extends State<ZNKSearchView> {
  final TextEditingController _controller = TextEditingController();
  //占位文本
  String _curPlaceholder = '输入搜索内容';

  @override
  void initState() {
    super.initState();
    _controller.addListener(() {
      print('text: ${_controller.text}');
    });
  }

  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    Size searchSize = Size(20, 20);
    return Container(
        margin: widget._style.margin,
        decoration: BoxDecoration(
            borderRadius: BorderRadius.circular(widget._style.cornerRadius),
            color: widget._style.backgroudColor),
        width: widget._style.width,
        height: widget._style.height,
        child: widget._style.enabled
            ? _searchFormField()
            : (Stack(
                  children: [
                    Container(
                        margin: EdgeInsets.only(
                            left:
                                (widget._style.height - searchSize.width) / 2.0,
                            top: (widget._style.height - searchSize.height) /
                                2.0),
                        width: searchSize.width,
                        height: searchSize.height,
                        child: Icon(Icons.search)),
                    widget.child,
                  ],
                ) ??
                Container()));
  }

  //搜索输入框
  Widget _searchFormField() {
    return Container(
      width: widget._style.width - 10,
      height: widget._style.height * (2 / 3.0),
      margin: EdgeInsets.only(left: 3, top: 5, right: 15, bottom: 5),
      child: TextFormField(
        controller: _controller,
        enabled: widget._style.enabled,
        style: TextStyle(color: widget._style.textColor),
        decoration: InputDecoration(
          contentPadding: EdgeInsets.only(bottom: 5),
          prefixIcon: Icon(Icons.search, color: Colors.grey),
          border: InputBorder.none,
        ),
      ),
    );
  }
}
