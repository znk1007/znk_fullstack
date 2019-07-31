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
  // 第一天是否星期天
  bool get sundayFirst;
  // 页数，3，5，7，。。。
  int get numberOfPage;

}