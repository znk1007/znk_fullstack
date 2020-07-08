import 'package:flb/util/random/random.dart';
import 'package:flutter/material.dart';
//获取占位回调
typedef SearchPlaceholderFunction = Function(String placeholder);
//获取输入内容回调
typedef SearchInputFunction = Function(String input);

class SearchStyle {
  //背景颜色
  Color backgroudColor = Colors.red;
  //倒角
  double cornerRadius = 6.0;
  //宽度
  double width = 200;
  //高度
  final double height = 45;
  //边距
  EdgeInsets margin = EdgeInsets.zero;
  //textColor 文字颜色
  Color textColor = Colors.grey[350];
  //是否可用、可输入
  bool enabled = true;
  //默认样式
  static SearchStyle defaultStyle() => SearchStyle();
  //初始化
  SearchStyle(
      {this.backgroudColor = Colors.red,
      this.cornerRadius = 6.0,
      this.width = 200,
      this.enabled = true,
      this.textColor = Colors.grey,
      this.margin = EdgeInsets.zero});
}

class SearchView extends StatefulWidget {
  //搜索样式
  SearchStyle _style;
  SearchView({Key key, SearchStyle style}) {
    _style = style ?? SearchStyle();
  }

  final _SearchViewState state = _SearchViewState();

  @override
  _SearchViewState createState() => state;
}

class _SearchViewState extends State<SearchView> {
  final TextEditingController _controller = TextEditingController();
  //占位文本
  String _curPlaceholder = '输入搜索内容';

  List<String> _placeholders = [];
  //更新占位文本
  void updatePlaceholer(List<String> texts) {
    setState(() {
    });
  }

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
    return Container(
      margin: widget._style.margin,
      decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(widget._style.cornerRadius),
          color: widget._style.backgroudColor),
      width: widget._style.width,
      height: widget._style.height,
      child: widget._style.enabled ? _searchFormField() : Container()
    );
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
          style: TextStyle(color:widget._style.textColor),
          decoration: InputDecoration(
            contentPadding: EdgeInsets.only(bottom:5),
            prefixIcon: Icon(Icons.search, color: Colors.grey),
            border: InputBorder.none,
          ),
        ),
      );
  }
}
