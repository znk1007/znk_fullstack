import 'package:znk/utils/calendar/custom/custom_date_help.dart';

class DateUtil {

  // 是否是今天
  static bool isToday(DateTime time) {
    return DateTime.now().difference(time).inDays == 0;
  }
  // 国历星期
  static String cnWeek(DateTime time, WeekType type) {
    switch (time.weekday) {
      case DateTime.monday:
        return type == WeekType.single ? '一' : type == WeekType.short ? '周一' : '星期一';
        break;
      case DateTime.tuesday:
        return type == WeekType.single ? '二' : type == WeekType.short ? '周二' : '星期二';
        break;
      case DateTime.wednesday:
        return type == WeekType.single ? '三' : type == WeekType.short ? '周三' : '星期三';
        break;
      case DateTime.thursday:
        return type == WeekType.single ? '四' : type == WeekType.short ? '周四' : '星期四';
        break;
      case DateTime.friday:
        return type == WeekType.single ? '五' : type == WeekType.short ? '周五' : '星期五';
        break;
      case DateTime.saturday:
        return type == WeekType.single ? '六' : type == WeekType.short ? '周六' : '星期六';
        break;
      case DateTime.sunday:
        return type == WeekType.single ? '日' : type == WeekType.short ? '周日' : '星期日';
        break;
      default:
    }
    return '';
  }
  // 国历月份显示
  static String cnMonth(DateTime time, MonthType type) {
    switch (time.month) {
      case DateTime.january:
        return type == MonthType.short ? '一' : '一月';
        break;
      case DateTime.february:
        return type == MonthType.short ? '二' : '二月';
        break;
      case DateTime.march:
        return type == MonthType.short ? '三' : '三月';
        break;
      case DateTime.april:
        return type == MonthType.short ? '四' : '四月';
        break;
      case DateTime.may:
        return type == MonthType.short ? '五' : '五月';
        break;
      case DateTime.june:
        return type == MonthType.short ? '六' : '六月';
        break;
      case DateTime.july:
        return type == MonthType.short ? '七' : '七月';
        break;
      case DateTime.august:
        return type == MonthType.short ? '八' : '八月';
        break;
      case DateTime.september:
        return type == MonthType.short ? '九' : '九月';
        break;
      case DateTime.october:
        return type == MonthType.short ? '十' : '十月';
        break;
      case DateTime.november:
        return type == MonthType.short ? '十一' : '十一月';
        break;
      case DateTime.december:
        return type == MonthType.short ? '十二' : '十二月';
        break;
        
      default:
    }
    return '';
  }
  // 是否闰年
  static bool isLeapYear(int year) {
    return ((year % 4 == 0) && (year % 100) != 0) || (year % 400 == 0); 
  }
  
  // 每月天数
  static int daysOfMonth(int year, int month) {
    int count = 30;
    switch (month) {
      case 2:
      {
        bool isLeap = isLeapYear(year);
        count = isLeap ? 29 : 28;
      }
        break;
      case 4:
      case 6:
      case 9:
      case 11:
      {
        count = 30;
      }
        break;
      case 1:
      case 3:
      case 5:
      case 7:
      case 8:
      case 10:
      {
        count = 31;
      }
        break;
      default:
    }
    return count;
  }

}