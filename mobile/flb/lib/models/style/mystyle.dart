import 'package:flb/pkg/screen/screen.dart';
import 'package:flutter/material.dart';

class MyPageStyle extends ChangeNotifier {
  //用户信息展示宽度
  double profileBgWidth = ZNKScreen.screenWidth;
  //用户信息展示高度
  double profileBgHeight = ZNKScreen.setHeight(180).toDouble();
  //用户信息区域背景颜色
  Color profileBgColor = Colors.red[400];
  //头像直径
  double avatarL = ZNKScreen.setWidth(60).toDouble();
  //头像边距
  EdgeInsets avatarMargin =
      EdgeInsets.only(left: 50, top: ZNKScreen.safeTopArea + 20);
  //昵称高度
  double nicknameHeight = ZNKScreen.setHeight(40).toDouble();
  //收益模块高度
  double eqHeight = ZNKScreen.setWidth(50);
  //行高
  double rowHeight = ZNKScreen.setWidth(40);
}
