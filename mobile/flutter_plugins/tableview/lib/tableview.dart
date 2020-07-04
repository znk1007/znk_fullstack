library znktable;

import 'package:flutter/material.dart';

//ZNKTableStyle 表格样式
enum ZNKTableStyle {
  plain, //平铺
  grouped, //分组
}

//ZNKIndexPath 地址索引
class ZNKIndexPath {
  final int section; //段下标
  final int row; //行下标
  ZNKIndexPath({this.section, this.row});
}

class ZNKTable extends StatelessWidget {
  //表格样式
  final ZNKTableStyle style;
  //段数
  final int numberOfSection;
  //每段行数
  final int Function(int section) numberOfRowsInSection;
  //滚轴方向
  final Axis scrollDirection;
  //头部视图
  final List<Widget> Function(BuildContext context, bool innerBoxIsScrolled)
      headerSliverBuilder;
  //分割视图
  final Widget Function(BuildContext context, int index) separatorBuilder;

  //ScrollController 嵌套滚动控制器
  // ScrollController _nestedScrollCtl;

  ZNKTable(
      {Key key,
      this.style = ZNKTableStyle.plain,
      this.scrollDirection = Axis.vertical,
      this.numberOfSection = 1,
      @required this.numberOfRowsInSection,
      this.separatorBuilder,
      this.headerSliverBuilder})
      : super(key: key);

  @override
  Widget build(BuildContext context) {
    bool test = true;
    return Container(
      child: this.headerSliverBuilder != null
          ? NestedScrollView(
              headerSliverBuilder: this.headerSliverBuilder, body: null)
          : ListView.separated(
              physics: ScrollPhysics(),
              scrollDirection: this.scrollDirection,
              itemCount: 50,
              separatorBuilder: (BuildContext context, int index) {
                return this.separatorBuilder ?? Container();
              },
              itemBuilder: (BuildContext context, int index) {
                return Text('data $index');
              },
            ),
    );
  }

  List<Widget> _defaultHeaderSliverBuilder(
      BuildContext context, bool innerBoxIsScrolled) {
    return <Widget>[SliverAppBar()];
  }
}
