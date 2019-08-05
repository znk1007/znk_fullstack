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
  // 页数，3，5，7，。。。
  int get numberOfPage;
  // 缓存加载出来的日历数据
  bool get keepCache;

}