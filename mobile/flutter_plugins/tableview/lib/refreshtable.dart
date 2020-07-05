//ZNKLoadingState 加载状态
import 'package:flutter/material.dart';

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
