import 'package:flutter/rendering.dart';

enum WeekType {
  single,
  short,
  full,
}

enum MonthType {
  short,
  full,
}

abstract class CustomDateHelper {
  // 星期显示格式
  WeekType weekType;
  // 月份显示格式
  MonthType monthType;
  // 每月第一天星期天
  int get firstWeekday;
  // 缓存加载出来的日历数据
  bool get keepCache;
  // 星期分割线宽度
  int get weekdaySeparatorWidth;
  // 星期分割线颜色
  Color get weekdaySeparatorColor;
  // 非本月日期颜色
  Color get outerDayColor;
  // 本月日期颜色
  Color get innerDayColor;
  // 间隔加载页数
  int get diffLoadPage;

}