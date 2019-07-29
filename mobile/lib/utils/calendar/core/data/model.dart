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
  bool isSelected = false;
  bool hasSchedule = false;// 是否有事项
  // 是否同一天
  bool isSameDay(CalendarModel other) {
    return this.year == other.year && this.month == other.month && this.day == other.day;
  }
  // 是否同一月
  bool isSameMonth(CalendarModel other) {
    return this.year == other.year && this.month == other.month;
  }

}

class CalendarManager {
  // 单例
  static CalendarManager get instance {
    if (_inner == null) {
      _inner = CalendarManager._();
    } 
    return _inner;
  }
  static CalendarManager _inner;
  CalendarManager._() {
    _models = [];
    _selectedModels = [];
  }

  List<CalendarModel> get calendarModels {
    return _models;
  }

  List<CalendarModel> get calendarSelectedModels {
    return _selectedModels;
  }

  List<CalendarModel> _selectedModels;

  // 模型
  List<CalendarModel> _models;
  // 当前页码
  int currentPage;

  int _firstYear = 1960;

  int _lastYear = 2060;

  bool _loadDay = false;

  // 预加载数据
  List<CalendarModel> preLoad({
    int startYear = 1960, 
    int endYear = 2060, 
    int currentYear = -1, 
    int currentMonth = -1,
    bool loadDay = false,
  }) {
    _loadDay = loadDay;
    _models = [];
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
        if (_loadDay) {
          int days = DateUtil.daysOfMonth(i, j);
          for (var k = 1; k <= days; k++) {
            CalendarModel model = CalendarModel()
              ..dateTime = DateTime(i, j, k);
              _models.add(model);
          }
        } else {
          CalendarModel model = CalendarModel()
            ..dateTime = DateTime(i, j);
          _models.add(model);
        }
      }
    }
    return _models;
  }

  void mapToView(int year, int month) {
    final numberOflines = DateUtil.numberOfLinesOfMonth(year, month);
    final totalNums = numberOflines * 7;
    for (var i = 0; i < totalNums; i++) {
      
    }
  }

}