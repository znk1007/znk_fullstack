import 'package:flutter/cupertino.dart';

class TabItem {
  /// 标题
  final String title;
  /// 常态图片
  final String icon;
  /// 高亮状态图片
  final String activeIcon;
  /// 下标
  final String index;
  /// 初始化
  TabItem({
    @required this.title,
    @required this.icon, 
    @required this.activeIcon, 
    @required this.index
  });
}



