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
  // 日期最大行数 6或7
  int get maxCalendarRows;
}