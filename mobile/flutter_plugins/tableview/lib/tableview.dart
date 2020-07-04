library znktable;

import 'package:flutter/material.dart';

//ZNKIndexPath 地址索引
class ZNKIndexPath {
  final int section; //段下标
  final int row; //行下标
  ZNKIndexPath({this.section, this.row});
}

class ZNKTable extends StatelessWidget {
  //段数
  final int numberOfSection;
  //每段行数
  final int Function(int section) numberOfRowsInSection;
  //每单元格视图
  final Widget Function(ZNKIndexPath indexPath) cellForRowAtIndexPath;
  //滚轴方向
  final Axis scrollDirection;
  //头部视图
  final List<Widget> Function(BuildContext context, bool innerBoxIsScrolled)
      headerSliverBuilder;
  //分割视图
  final Widget Function(BuildContext context, int index) separatorBuilder;
  //选择指定行
  final void Function(BuildContext context, ZNKIndexPath indexPath) didSelectRowAtIndexPath;

  //ScrollController 嵌套滚动控制器
  // ScrollController _nestedScrollCtl;

  ZNKTable(
      {Key key,
      this.scrollDirection = Axis.vertical,
      this.numberOfSection = 1,
      @required this.numberOfRowsInSection,
      @required this.cellForRowAtIndexPath,
      this.didSelectRowAtIndexPath,
      this.separatorBuilder,
      this.headerSliverBuilder})
      : assert(numberOfSection >= 1),
        super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      child: this.headerSliverBuilder != null
          ? NestedScrollView(
              headerSliverBuilder: this.headerSliverBuilder,
              body: _separatedListView())
          : _separatedListView(),
    );
  }

  //分割线表格
  Widget _separatedListView() {
    return this.numberOfSection > 1
        ? ListView.separated(
            physics: ScrollPhysics(),
            scrollDirection: this.scrollDirection,
            itemCount: this.numberOfSection,
            separatorBuilder: (BuildContext context, int index) {
              return this.separatorBuilder ?? Container();
            },
            itemBuilder: (BuildContext context, int index) {
              return Text('data $index');
            },
          )
        : ListView.separated(
            itemBuilder: (BuildContext ctx, int index) {
              return GestureDetector(child: this
                  .cellForRowAtIndexPath(ZNKIndexPath(section: 0, row: index)),onTap: () {
                    if (this.didSelectRowAtIndexPath != null) {
                      this.didSelectRowAtIndexPath(ctx, ZNKIndexPath(section: 0, row: index));
                    }
                  },);
            },
            separatorBuilder: (BuildContext context, int index) {
              return this.separatorBuilder ?? Container();
            },
            itemCount: this.numberOfRowsInSection(0));
  }
}
