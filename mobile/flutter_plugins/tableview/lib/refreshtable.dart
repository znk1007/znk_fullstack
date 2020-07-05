library znkrefreshtable;
//ZNKLoadingState 加载状态
import 'package:flutter/material.dart';
import 'package:tableview/tableview.dart';

enum ZNKLoadingState {
  idle, //闲置
  loading, //加载中
  finished, //加载完成
}

//ZNKPullRefresh 下拉刷新
class ZNKPullRefresh {
  //下拉刷新临界点前标题
  final String beforeTitle;
  //下拉刷新临界点前标题
  final String afterTitle;
  //下拉刷新警示视图
  final Widget indicator;
  //上次更新日期
  DateTime time;
  //初始化
  ZNKPullRefresh(
      {this.beforeTitle = '下拉刷新...',
      this.afterTitle = '松开刷新...',
      this.indicator = const CircularProgressIndicator()});
}

//ZNKPushRefresh 上拉刷新数据配置
class ZNKPushRefresh {
  //上拉刷新临界点前标题
  String beforeTitle;
  //上拉刷新临界点后标题
  String afterTitle;
  //上拉刷新警示视图
  Widget indicator;
  //上次更新日期
  DateTime time;
  //初始化
  ZNKPushRefresh(
      {this.beforeTitle = '上拉刷新...',
      this.afterTitle = '松开刷新...',
      this.indicator = const CircularProgressIndicator()});
}

class ZNKRefreshView extends StatefulWidget {
  // ScrollController 嵌套滚动控制器
  ScrollController _scrollCtl;
  //刷新是否有效
  bool _refreshValid = false;
  //最大偏移值
  final double _maxOffset = 100;
  //下拉刷新
  final ZNKPullRefresh pullRefresh;
  //下拉刷新回调
  final Function(double offset, double total, ZNKLoadingState state)
      pullRefreshHandler;
  //上拉刷新
  final ZNKPushRefresh pushRefresh;
  //上拉刷新回调
  final Function(double offset, double total, ZNKLoadingState state)
      pushRefreshHandler;
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

  ZNKRefreshView({
    Key key,
    this.pullRefresh,
    this.pullRefreshHandler,
    this.pushRefresh,
    this.pushRefreshHandler,
    this.scrollDirection = Axis.vertical,
    this.numberOfSection = 1,
    @required this.numberOfRowsInSection,
    @required this.cellForRowAtIndexPath,
    this.viewForHeaderInSection,
    this.didSelectRowAtIndexPath,
    this.heightForRowAtIndexPath,
    this.separatorBuilder,
    this.headerSliverBuilder,
  }) : super(key: key);

  @override
  _ZNKRefreshViewState createState() => _ZNKRefreshViewState();
}

class _ZNKRefreshViewState extends State<ZNKRefreshView> {
  @override
  Widget build(BuildContext context) {
    Size size = MediaQuery.of(context).size;
    return Container(
      child: Stack(children: [
        Container(
          width: size.width,
          height: 2,
          color: Colors.brown,
        ),
        ZNKTable(
            numberOfRowsInSection: widget.numberOfRowsInSection,
            cellForRowAtIndexPath: widget.cellForRowAtIndexPath,
            scrollDirection: widget.scrollDirection,
            numberOfSection: widget.numberOfSection,
            viewForHeaderInSection: widget.viewForHeaderInSection,
            didSelectRowAtIndexPath: widget.didSelectRowAtIndexPath,
            separatorBuilder: widget.separatorBuilder,
            headerSliverBuilder: widget.headerSliverBuilder,
            heightForRowAtIndexPath: widget.heightForRowAtIndexPath,
          ),
        Container(
          width: size.width,
          height: 2,
          color: Colors.cyan,
        ),
      ]),
    );
  }
}
