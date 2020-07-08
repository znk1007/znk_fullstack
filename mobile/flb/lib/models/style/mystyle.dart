import 'package:flb/pkg/screen/screen.dart';
import 'package:flutter/material.dart';

class MyPageStyle extends ChangeNotifier {
  //用户信息展示宽度
  double profileBgWidth = Screen.screenWidth;
  //用户信息展示高度
  double profileBgHeight = Screen.setHeight(180).toDouble();
  //用户信息区域背景颜色
  Color profileBgColor = Colors.red[400];
  //头像直径
  double avatarL = Screen.setWidth(60).toDouble();
  //头像边距
  EdgeInsets avatarMargin =
      EdgeInsets.only(left: 50, top: Screen.safeTopArea + 20);
  //昵称高度
  double nicknameHeight = Screen.setHeight(40).toDouble();
  //收益模块高度
  double eqHeight = Screen.setWidth(50);
  //行高
  double rowHeight = Screen.setWidth(40);
}