import 'package:flb/util/random/color.dart';
import 'package:flutter/material.dart';

class SearchStyle {
  //背景颜色
  Color backgroudColor = Colors.red;
  //倒角
  double cornerRadius = 6.0;
  //大小
  Size size = Size(200, 60);
  //边距
  EdgeInsets margin = EdgeInsets.zero;
  //默认样式
  static SearchStyle defaultStyle() => SearchStyle();
  //初始化
  SearchStyle(
      {this.backgroudColor = Colors.red,
      this.cornerRadius = 6.0,
      this.size = const Size(200, 60),
      this.margin = EdgeInsets.zero});
}

class SearchView extends StatefulWidget {
  //搜索样式
  SearchStyle _style;
  SearchView({Key key, SearchStyle style}) {
    _style = style ?? SearchStyle();
  }

  @override
  _SearchViewState createState() => _SearchViewState();
}

class _SearchViewState extends State<SearchView> {
  final TextEditingController _controller = TextEditingController();

  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    print('height: ${widget._style.size.height}');
    return Container(
      margin: widget._style.margin,
      decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(widget._style.cornerRadius),
          color: widget._style.backgroudColor),
      width: widget._style.size.width,
      height: widget._style.size.height,
      child: Container(
        width: widget._style.size.width - 10,
        height: widget._style.size.height * (2 / 3.0),
        margin: EdgeInsets.only(left: 3, top: 5, right: 3, bottom: 5),
        child: TextFormField(
          decoration: InputDecoration(
            prefixIcon: Icon(Icons.search, color: Colors.grey),
            border: InputBorder.none,
          ),
        ),
      ),
    );
  }
}
