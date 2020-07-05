import 'package:flutter/material.dart';

//ZNKIndexPath 地址索引
class ZNKIndexPath {
  final int section; //段下标
  final int row; //行下标
  ZNKIndexPath({this.section, this.row});
}

class ZNKTable extends StatelessWidget {
  final bool shrinkWrap;
  //段数
  final int numberOfSection;
  //每段行数
  final int Function(int section) numberOfRowsInSection;
  //每单元格视图
  final Widget Function(BuildContext context, ZNKIndexPath indexPath)
      cellForRowAtIndexPath;
  //滚轴方向
  final Axis scrollDirection;
  //头部视图
  final List<Widget> Function(BuildContext context, bool innerBoxIsScrolled)
      headerSliverBuilder;
  //分割视图
  final Widget Function(BuildContext context, ZNKIndexPath indexPath)
      separatorBuilder;
  //选择指定行
  final void Function(BuildContext context, ZNKIndexPath indexPath)
      didSelectRowAtIndexPath;
  //行高
  final double Function(BuildContext context, ZNKIndexPath indexPath)
      heightForRowAtIndexPath;
  //段头视图
  final Widget Function(BuildContext context, int section)
      viewForHeaderInSection;

  ZNKTable({
    Key key,
    this.shrinkWrap = true,
    this.scrollDirection = Axis.vertical,
    this.numberOfSection = 1,
    @required this.numberOfRowsInSection,
    @required this.cellForRowAtIndexPath,
    this.viewForHeaderInSection,
    this.didSelectRowAtIndexPath,
    this.heightForRowAtIndexPath,
    this.separatorBuilder,
    this.headerSliverBuilder,
  })  : assert(numberOfSection >= 1),
        super(key: key);

  @override
  Widget build(BuildContext context) {
    return this.headerSliverBuilder != null
          ? NestedScrollView(
              headerSliverBuilder: this.headerSliverBuilder,
              body: _separatedListView())
          : _separatedListView();
  }

  //单列表
  Widget _singleSeparatedListView(int section, bool shrinkWrap, bool scrollable) {
    return ListView.separated(
        physics: scrollable
            ? NeverScrollableScrollPhysics()
            : BouncingScrollPhysics(),
        shrinkWrap: shrinkWrap,
        scrollDirection: this.scrollDirection,
        itemBuilder: (BuildContext ctx, int index) {
          double rowHeight = this.heightForRowAtIndexPath != null
              ? this.heightForRowAtIndexPath(
                  ctx, ZNKIndexPath(section: section, row: index))
              : 44;
          return GestureDetector(
            child: Container(
              child: this.cellForRowAtIndexPath(
                  ctx, ZNKIndexPath(section: section, row: index)),
              height: rowHeight,
            ),
            onTap: () {
              if (this.didSelectRowAtIndexPath != null) {
                this.didSelectRowAtIndexPath(
                    ctx, ZNKIndexPath(section: section, row: index));
              }
            },
          );
        },
        separatorBuilder: (BuildContext context, int index) {
          return this.separatorBuilder != null
              ? this.separatorBuilder(
                  context, ZNKIndexPath(section: section, row: index))
              : Divider(height: 2);
        },
        itemCount: this.numberOfRowsInSection(section));
  }

  //嵌套表格
  Widget _separatedListView() {
    return this.numberOfSection > 1
        ? ListView.separated(
            physics: ScrollPhysics(),
            scrollDirection: this.scrollDirection,
            itemCount: this.numberOfSection,
            separatorBuilder: (BuildContext context, int index) {
              return Container();
            },
            itemBuilder: (BuildContext context, int section) {
              return Column(
                children: [
                  this.viewForHeaderInSection != null
                      ? this.viewForHeaderInSection(context, section)
                      : Container(),
                  _singleSeparatedListView(section, this.shrinkWrap, false),
                ],
              );
            },
          )
        : _singleSeparatedListView(0, this.shrinkWrap, true);
  }
}