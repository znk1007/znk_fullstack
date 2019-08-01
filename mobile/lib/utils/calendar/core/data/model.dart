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
  }

  int get totalPage {
    return _pages;
  }

  // 模型
  Map<String, CalendarModel> _modelsMap;
  // 当前页码
  int currentPage = 0;
  // 总页数
  int _pages = 0;

  int _firstYear = 1960;

  int _lastYear = 2060;

  DateTime _firstDate;

  DateTime _lastDate;

  String _key(int year, int month, int day) {
    return '$year' + '$month' + '$day';
  }

  // 预加载数据
  void preLoad({
    int startYear = 1960, 
    int endYear = 2060, 
  }) {
    _modelsMap = Map();
    int s = startYear < _firstYear ? startYear : _firstYear;
    _firstYear = s;
    int e = endYear > _lastYear ? endYear : _lastYear;
    _lastYear = e;
    for (var i = _firstYear; i <= _lastYear; i++) {
      for (var j = 1; j <= 12; j++) {
        int days = DateUtil.daysOfMonth(i, j);
        for (var k = 1; k <= days; k++) {
          CalendarModel model = CalendarModel()
            ..dateTime = DateTime(i, j, k);
            _modelsMap[_key(i, j, k)] = model;
        }
      }
    }
  }

  CalendarModel getModel(int year, int month, int day) {
    CalendarModel model = _modelsMap[_key(year, month, day)];
    return model;
  }

  List<String> _manageKeys(int year, int month) {
    int days = DateUtil.daysOfMonth(year, month);
    List<String> keys = [];
    for (var i = 0; i < days; i++) {
      keys.add(_key(year, month, i));
    }
    return keys;
  }

  // 切割模型
  List<CalendarModel> subModels(int year, int month) {
    List<CalendarModel> models = [];
    final keys = _manageKeys(year, month);
    for (var key in keys) {
      final model = _modelsMap[key];
      if (model != null) {
        models.add(_modelsMap[key]);
      }
    }
    return models;
  }
  // 加载数据
  List<CalendarModel> load(DateTime startTime, DateTime endTime) {
    int startYear = startTime.year;
    int startMonth = startTime.month;
    int startDay = startTime.day;
    int endYear = endTime.year;
    int endMonth = endTime.month;
    int endDay = endTime.day;
    List<CalendarModel> models = [];
    for (var i = startYear; i <= endYear; i++) {
      if (i == startYear && startYear != endYear) {
        for (var sm = startMonth; sm <= 12; sm++) {
          int days = DateUtil.daysOfMonth(i, sm);
          if (sm == startMonth) {
            for (var sd = startDay; sd <= days; sd++) {
              final model = _modelsMap[_key(i, sm, sd)];
              if (model != null) {
                models.add(model);
              }
            }
          } else {
            for (var sd = 1; sd <= days; sd++) {
              final model = _modelsMap[_key(i, sm, sd)];
              if (model != null) {
                models.add(model);
              }
            }
          }
        }
      } else if (i == endYear && startYear != endYear) {
        for (var em = 1; em <= endMonth; em++) {
          int days = DateUtil.daysOfMonth(i, em);
          if (em == endMonth) {
            for (var ed = 1; ed <= days; ed++) {
              final model = _modelsMap[_key(i, em, ed)];
              if (model != null) {
                models.add(model);
              }
            }
          } else {
            for (var ed = 1; ed <= days; ed++) {
              final model = _modelsMap[_key(i, em, ed)];
              if (model != null) {
                models.add(model);
              }
            }
          }
        }
      } else if (startYear == endYear && i == startYear && i == endYear) {
        for (var em = startMonth; em <= endMonth; em++) {
          int days = DateUtil.daysOfMonth(i, em);
          if (em == startMonth) {
            for (var ed = startDay; ed <= days; ed++) {
              final model = _modelsMap[_key(i, em, ed)];
              if (model != null) {
                models.add(model);
              }
            }
          } else if (em == endMonth) {
            for (var ed = 1; ed <= endDay; ed++) {
              final model = _modelsMap[_key(i, em, ed)];
              if (model != null) {
                models.add(model);
              }
            }
          } else {
            for (var ed = 1; ed <= days; ed++) {
              final model = _modelsMap[_key(i, em, ed)];
              if (model != null) {
                models.add(model);
              }
            }
          }
        }
      } else {
        for (var mm = 1; mm <= 12; mm++) {
          int days = DateUtil.daysOfMonth(i, mm);
          for (var md = 1; md <= days; md++) {
            final model = _modelsMap[_key(i, mm, md)];
            if (model != null) {
              models.add(model);
            }
          }
        }
      }
    }
    // print('number of model: ${models.length}');
    // for (var model in models) {
    //   print('time: ${model.dateTime}');
    // }
    return models;
  }
  // 年差
  int diffYear(int month, int diffMonth) {
    int tempYear = 0;
    if (diffMonth > 0) {
      int divideYear = diffMonth ~/ 12;
      int modeYear = diffMonth % 12;
      bool more = (modeYear + month >= 12);
      tempYear += divideYear;
      if (more) {
        tempYear += 1;
      }
    } else {
      diffMonth = -diffMonth;
      int divideYear = diffMonth ~/ 12;
      int modeYear = diffMonth % 12;
      bool more = (modeYear - month >= 0);
      tempYear -= divideYear;
      if (more) {
        tempYear -= 1;
      }
    }
    return tempYear;
  }
  // 月差
  int diffMonth(int month, int diffMonth) {
    int tempMonth = month;
    if (diffMonth > 0) {
      int modeYear = diffMonth % 12;
      bool more = (modeYear + month >= 12);
      if (more) {
        tempMonth = modeYear - month + 1;
      }
    } else {
      diffMonth = -diffMonth;
      int modeYear = diffMonth % 12;
      bool more = (modeYear - month >= 0);
      if (more) {
        tempMonth = 12 - (modeYear - month);
      }
    }
    return tempMonth;
  }

  void diffLoad(int year, int month, int diffMonth) {
    int tempYear = year;
    int tempMonth = month;
    if (diffMonth > 0) {
      int divideYear = diffMonth ~/ 12;
      int modeYear = diffMonth % 12;
      bool more = (modeYear + month >= 12);
      tempYear += divideYear;
      if (more) {
        tempMonth = modeYear - month + 1;
        tempYear += 1;
      }
    } else {
      diffMonth = -diffMonth;
      int divideYear = diffMonth ~/ 12;
      int modeYear = diffMonth % 12;
      bool more = (modeYear - month >= 0);
      tempYear -= divideYear;
      if (more) {
        tempMonth = 12 - (modeYear - month);
        tempYear -= 1;
      }
    }
    
    print('current year: $year, temp year: $tempYear, current month: $month, temp month: $tempMonth');
  }

  // 日期转视图
  List<CalendarModel> mapToView(int year, int month, {bool sundayFirst = true}) {
    final numberOflines = DateUtil.numberOfLinesOfMonth(year, month, true);
    int firstWeekday = DateUtil.firstWeekdayOfMonthForYearMonth(year, month);
    int curMonthDays = DateUtil.daysOfMonth(year, month);
    // last month
    // 1 sundayFirst == true ? -1 : -0
    // 2 sundayFirst == true ? -2 : -1
    // 3 sundayFirst == true ? -3 : -2
    // 4 sundayFirst == true ? -4 : -3
    // 5 sundayFirst == true ? -5 : -4
    // 6 sundayFirst == true ? -6 : -5
    // 7 sundayFirst == true ? -0 : -6
    // next month
    // total-current-last=next *
    // 1 sundayFirst == true ? +5 : +6
    // 2 sundayFirst == true ? +4 : +5
    // 3 sundayFirst == true ? +3 : +4
    // 4 sundayFirst == true ? +2 : +3
    // 5 sundayFirst == true ? +1 : +2
    // 6 sundayFirst == true ? +0 : +1
    // 7 sundayFirst == true ? +6 : +0
    int lastMonth = month - 1;
    int lastMonthDays = 0;
    int lastMonthDiff = -1;
    int fixLastMontDays = sundayFirst ? firstWeekday == 7 ? 0 : firstWeekday : firstWeekday - 1;
    int nextMonth = month+1;
    int tempCurMonthDays = curMonthDays + fixLastMontDays;
    int currentMonthIdx = 1;

    final totalNums = numberOflines * 7;
    String key = '';
    List<CalendarModel> models = [];
    for (var i = 0; i < totalNums; i++) {
      //last month      
      if (i < fixLastMontDays) {
        int lastYear = year;
        if (lastMonth <= 0) {
          lastMonth = 12;
          lastYear--;
        }
        if (lastMonthDiff == -1) {
          lastMonthDays = DateUtil.daysOfMonth(lastYear, lastMonth);
          lastMonthDiff = firstWeekday;
        }
        key = _key(lastYear, lastMonth, lastMonthDays);
        lastMonthDays--;
      } else if (i >= tempCurMonthDays) {
        // next month
        int nextYear = year;
        if (nextMonth > 12) {
          nextMonth = 1;
          nextYear++;
        }
        key = _key(nextYear, nextMonth, i-tempCurMonthDays+1);
        // print('\nnext: $key');
        
      } else {
        key = _key(year, month, currentMonthIdx++);
      }
      if (key.isNotEmpty) {
        models.add(_modelsMap[key]);
      }
    }
    return models;
  }

}