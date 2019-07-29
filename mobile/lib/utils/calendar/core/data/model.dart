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
    _modelsMap = Map();
    _selectedModels = Map();
  }

  Map<String, CalendarModel> get calendarModelsMap {
    return _modelsMap;
  }

  Map<String, CalendarModel> get calendarSelectedModelsMap {
    return _selectedModels;
  }

  Map<String, CalendarModel> _selectedModels;

  // 模型
  Map<String, CalendarModel> _modelsMap;
  // 当前页码
  int currentPage;

  int _firstYear = 1960;

  int _lastYear = 2060;

  String _key(int year, int month, int day) {
    return '$year' + '$month' + '$day';
  }

  // 预加载数据
  void preLoad({
    int startYear = 1960, 
    int endYear = 2060, 
    int currentYear = -1, 
    int currentMonth = -1,
  }) {
    _modelsMap = Map();
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
        
        int days = DateUtil.daysOfMonth(i, j);
        for (var k = 1; k <= days; k++) {
          CalendarModel model = CalendarModel()
            ..dateTime = DateTime(i, j, k);
            _modelsMap[_key(i, j, k)] = model;
        }
        currentMonthIdx++;
      }
    }
  }

  CalendarModel getModel(int year, int month, int day) {
    CalendarModel model = _modelsMap[_key(year, month, day)];
    return model;
  }

  List<CalendarModel> _getModels(int year, int month, int startDay, int offset) {
    CalendarModel temp = getModel(year, month, startDay);
    if (temp == null) {
      return null;
    }
    List<CalendarModel> temps = [];
    if (offset == 0) {
      return [temp];
    } else if (offset > 0) {
      int curMonthDays = DateUtil.daysOfMonth(year, month);
      for (var i = 1; i <= offset; i++) {
        int tempDay = startDay + i;
        if (tempDay > curMonthDays) {
          tempDay -= curMonthDays;
          month += 1;
        }
        print('> 0 temp day: $tempDay');
        CalendarModel model = _modelsMap[_key(year, month, tempDay)];
        if (model != null) {
          temps.add(model);
        }
      }
    } else {
      int curMonthDays = DateUtil.daysOfMonth(year, month-1);
      offset = -offset;
      for (var i = offset; i > 0; i--) {
        int tempDay = startDay - i;
        if (tempDay <= 0) {
          tempDay = curMonthDays + tempDay;
          month -= 1;
        }
        CalendarModel model = _modelsMap[_key(year, month, tempDay)];
        if (model != null) {
          temps.add(model);
        }
      }
    }
    return temps;
  }

  // 切割模型
  List<CalendarModel> subModels(int year, int month) {
    List<CalendarModel> models = [];
    for (var v in _modelsMap.values) {
      if (v.year == year && v.month == month) {
        models.add(v);
      }
    }
    return models;
  }

  void mapToView(int year, int month) {
    final numberOflines = DateUtil.numberOfLinesOfMonth(year, month);
    int firstWeekday = DateUtil.firstWeekdayOfMonth_1(year, month);
    // 1 -1
    // 2 -2
    // 3 -3
    // 4 -4
    // 5 -5
    // 6 -6
    // 7 -0
    
    final totalNums = numberOflines * 7;
    for (var i = 0; i < totalNums; i++) {
      //上个月最后几天 >= 0

    }
  }

}