import 'package:flutter/material.dart';

class ThemeStyle extends ChangeNotifier {
  //主题红色
  final Color redColor = Color(0xFFD81E06);
  //大视图背景颜色
  final Color backgroundColor = Color(0xFFF5F5F5);
  //文本深色
  final Color dartTextColor = Color(0xFF333333);
  //文本中色
  final Color middleTextColor = Color(0xFF666666);

  final Color lightTextColor = Color(0xFF999999);
  //分栏高度
  final double tabbarHeight = 48;
  //基础颜色
  Color primarySwatch = Colors.blue;
  //显示密度
  VisualDensity visualDensity = VisualDensity.adaptivePlatformDensity;
}
