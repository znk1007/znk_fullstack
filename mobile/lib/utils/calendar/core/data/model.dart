import 'package:znk/utils/calendar/core/data/util.dart';

class CalendarModel {
  DateTime dateTime;
  int get year {
    return dateTime.year;
  }
  int get month {
    return dateTime.month;
  }
  int get day {
    return dateTime.day;
  }

  bool get isToday {
    return dateTime.difference(DateTime.now().add(Duration(days: 56))).inDays == 0; 
  }
  bool isSelected;
  bool hasSchedule;// 是否有事项
}

class ModelManager {
  // 单例
  static ModelManager get instance {
    if (_inner == null) {
      _inner = ModelManager._();
    } 
    return _inner;
  }
  static ModelManager _inner;
  ModelManager._();
  // 模型
  List<CalendarModel> models;
  // 当前页码
  int currentPage;

  int _firstYear = 1960;

  int _lastYear = 2060;
  // 预加载数据
  Future preLoad({int startYear = 1960, int endYear = 2060, int currentYear = -1, int currentMonth = -1}) async {
    int s = startYear < _firstYear ? startYear : _firstYear;
    _firstYear = s;
    int currentMonthIdx = 0;
    int e = endYear > _lastYear ? endYear : _lastYear;
    _lastYear = e;
    for (var i = _firstYear; i < _lastYear; i++) {
      for (var j = 1; j <= 12; j++) {
        if (currentYear == -1 || currentMonth == -1) {
          DateTime now = DateTime.now();
          currentYear = now.year;
          currentMonth = now.month;
        }
        if (i == currentYear && j == currentMonth) {
          currentPage = currentMonthIdx;
        }
        currentMonthIdx++;
        int days = DateUtil.daysOfMonth(i, j);
        for (var k = 0; k < days; k++) {
          
        }
      }
    }
  }

}