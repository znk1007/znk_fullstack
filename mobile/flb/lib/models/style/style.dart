import 'package:flutter/material.dart';

class ThemeStyle extends ChangeNotifier {
  //分栏高度
  final double tabbarHeight = 48;
  //基础颜色
  Color primarySwatch = Colors.blue; 
  //显示密度
  VisualDensity visualDensity = VisualDensity.adaptivePlatformDensity;
}
