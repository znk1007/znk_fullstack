import 'package:flutter/material.dart';

class TabBarItem {
  /// 标题组件
  final Widget title;
  /// 常态视图组件
  final Widget icon;
  /// 选中状态视图组件
  final Widget activeIcon;
  /// 背景颜色
  final Color backgroundColor;
  /// 角标组件
  final Widget badge;
  /// 角标数
  final String badgeNum;
  /// 角标颜色
  final Color badgeColor;

  /// 初始化
  TabBarItem({
    @required this.icon,
    this.title,
    Widget activeIcon,
    this.backgroundColor,
    this.badge,
    this.badgeNum,
    Color badgeColor,
  })  : activeIcon = activeIcon ?? icon,
        badgeColor = badgeColor ?? Colors.red;
}
