import 'package:znk/utils/calendar/core/data/util.dart';

class CalendarPageModel {
  final int year;
  final int month;
  final List<CalendarModel> models;
  CalendarPageModel({this.year, this.month, this.models});
}

class CalendarModel {
  // 日期
  DateTime dateTime;
  // 行
  int row;
  // 列
  int column;
  // 是否已选中
  bool isSelected = false;
  // 是否有日程
  bool hasSchedule = false;// 是否有事项
  // 日程id
  String scheduleId;
  // 年份
  int get year {
    return dateTime.year;
  }
  // 月份
  int get month {
    return dateTime.month;
  }
  // 天
  int get day {
    return dateTime.day;
  }
  // 是否今天
  bool get isToday {
    return dateTime.difference(DateTime.now().add(Duration(days: 56))).inDays == 0; 
  }

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
  CalendarManager._();

 
  // 模型
  Map<String, CalendarModel> _modelsMap = Map();

  int _firstYear = 1960;

  int _lastYear = 2060;

  // 分页网格已加载模型
  List<CalendarPageModel> _pageModels = [];
  // 键值
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

  CalendarModel getModel(
    int year, 
    int month, 
    int day
  ) {
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
  List<CalendarModel> load(
    DateTime startTime, 
    DateTime endTime
  ) {
    int startYear = startTime.year;
    int startMonth = startTime.month;
    int startDay = startTime.day == 0 ? 1 : startTime.day;
    int endYear = endTime.year;
    int endMonth = endTime.month;
    int endDay = endTime.day == 0 ? DateUtil.daysOfMonth(endYear, endMonth) : endTime.day;
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
  int diffYear(
    int month, 
    int diffMonth
  ) {
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
  int diffMonth(
    int month, 
    int diffMonth
  ) {
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
  // 差加载
  List<CalendarModel> diffLoad(
    int year, 
    int month, 
    int diffMonth
  ) {
    int tempYear = year;
    int tempMonth = month;
    DateTime startTime;
    DateTime endTime;
    if (diffMonth > 0) {
      int divideYear = diffMonth ~/ 12;
      int modeYear = diffMonth % 12;
      bool more = (modeYear + month >= 12);
      tempYear += divideYear;
      if (more) {
        tempMonth = modeYear - month + 1;
        tempYear += 1;
      }
      startTime = DateTime(year, month, 0);
      endTime = DateTime(tempYear, tempMonth, 0);
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
      startTime = DateTime.utc(tempYear, tempMonth, 1);
      endTime = DateTime.utc(year, month, DateUtil.daysOfMonth(year, month));
    }
    return load(startTime, endTime);
  }
  // 星期差
  int diffWeekday(
    int firstWeekday, 
    int currentWeekday, 
    {
      bool backward = true
    }
  ) {
    int diff = currentWeekday - firstWeekday;
    if (backward) {
      if (diff < 0) {
        diff += 7;
      }
    } else {
      if (diff < 0) {
        diff = -diff - 1 ;
      } else {
        diff = 6 - diff;
      }
    }
    return diff;
  }
  // 测试星期差
  void testDiffWeekday(
    {
      bool backward = true
    }
  ) {
    for (var i = 1; i <= 7; i++) {
      print(' ');
      for (var j = 1; j <= 7; j++) {
        int diff = diffWeekday(i, j, backward: backward);
        print('first weekday = $i : the weekday = $j : diffWeekday = $diff');
      }
    print('-------------------------');
    }
  }

  // pre month
  // * firstWeekday = 7
  // * 1 -1
  // * 2 -2
  // * 3 -3
  // * 4 -4
  // * 5 -5
  // * 6 -6
  // * 7 -0
  // * firstWeekday = 1
  // * 1 -0
  // * 2 -1
  // * 3 -2
  // * 4 -3
  // * 5 -4
  // * 6 -5
  // * 7 -6
  // * firstWeekday = 2
  // * 1 -6
  // * 2 -0
  // * 3 -1
  // * 4 -2
  // * 5 -3
  // * 6 -4
  // * 7 -5
  // * firstWeekday = 3
  // * 1 -0
  // * 2 -1
  // * 3 -2
  // * 4 -3
  // * 5 -4
  // * 6 -5
  // * 7 -6
  // * firstWeekday = 4
  // * 1 -0
  // * 2 -1
  // * 3 -2
  // * 4 -3
  // * 5 -4
  // * 6 -5
  // * 7 -6
  // * firstWeekday = 5
  // * 1 -0
  // * 2 -1
  // * 3 -2
  // * 4 -3
  // * 5 -4
  // * 6 -5
  // * 7 -6
  // * firstWeekday = 6
  // * 1 -0
  // * 2 -1
  // * 3 -2
  // * 4 -3
  // * 5 -4
  // * 6 -5
  // * 7 -6
  // int diff = currentWeekday - firstWeekday;
  // if (diff < 0) {
  //   diff += 7;
  // }
  // diff;
  // next month
  // if (diff < 0) {
  //   diff = -diff - 1 ;
  // } else {
  //   diff = 6 - diff;
  // }
  // 日期和日历对应，可设置首日星期，适用二维图如GridView
  List<CalendarModel> mapToGridView(
    int year, 
    int month, 
    {
      int firstWeekday = 7, 
      bool fixedLines = true
    }
  ) {
    final numberOfLines = DateUtil.numberOfLinesOfMonth(year, month, fixedLines);
    int firstDayWeekday = DateUtil.firstWeekdayOfMonthForYearMonth(year, month);
    int curMonthDays = DateUtil.daysOfMonth(year, month);
    int preDays = diffWeekday(firstWeekday, firstDayWeekday);
    int tempCurMonthDays = curMonthDays + preDays;
    int preYear = year;
    int preMonth = month - 1;
    if (preMonth <= 0) {
      preMonth = 12;
      preYear--;
    }
    int nextYear = year;
    int nextMonth = month + 1;
    if (nextMonth > 12) {
      nextMonth = 12;
      nextYear++;
    }
    int preMonthDays = DateUtil.daysOfMonth(preYear, preMonth) - preDays;
    int currentMonthDayIdx = 1;
    List<CalendarModel> models = [];
    for (var i = 0; i < numberOfLines; i++) {
      for (var j = 0; j < 7; j++) {
        int idx = (i * 7) + j;
        if (idx < preDays) {
          int day = ++preMonthDays;
          String key = _key(preYear, preMonth, day);
          CalendarModel model = _modelsMap[key];
          if (model != null) {
            model.column = j;
            model.row = i;
            models.add(model);
          } else {
            model = CalendarModel()
              ..dateTime = DateTime(preYear, preMonth, day);
            _modelsMap[key] = model;
            models.add(model);
            // print('pre month day null');
          }
        } else if (idx >= tempCurMonthDays) {
          int day = idx-tempCurMonthDays+1;
          String key = _key(nextYear, nextMonth, day);
          CalendarModel model = _modelsMap[key];
          if (model != null) {
            model.column = j;
            model.row = i;
            models.add(model);
          } else {
            model = CalendarModel()
              ..dateTime = DateTime(preYear, preMonth, preMonthDays);
            _modelsMap[key] = model;
            models.add(model);
            // print('next month day null');
          }
        } else {
          int day = currentMonthDayIdx++;
          String key = _key(year, month, day);
          CalendarModel model = _modelsMap[key];
          if (model != null) {
            model.column = j;
            model.row = i;
            models.add(model);
          } else {
            model = CalendarModel()
              ..dateTime = DateTime(year, month, day);
            _modelsMap[key] = model;
            models.add(model);
            // print('current month day null');
          }
        }
      }
    }
    return models;
  }
  // 日期和日历对应，可设置首日星期，适用二维图如GridView+pageView
  List<CalendarPageModel> mapToGridViews(
    int year, 
    int month, 
    {
      int pages = 5, 
      int firstWeekday = 7, 
      bool fixedLines = true,
      bool keep = true,
    }
  ) {
    pages = pages < 3 ? 3 : pages;
    int middlePage = pages ~/ 2;
    int tempMonth = month;
    int tempYear = year;
    int len = _pageModels.length;
    if (keep == false && len > pages) {
      int idx = _pageModels.indexWhere((page) => page.year == year && page.month == month);
      if (idx != -1) {
        int removeIdx = len - 1 - idx;
        if (removeIdx < idx) {
          _pageModels.removeRange(0, removeIdx);
        } else if (removeIdx > idx) {
          _pageModels.removeRange(removeIdx + 1, len - 1);
        }
      }
    }
    print('  ');

    for (var i = 0; i < pages; i++) {
      tempMonth = month - (middlePage - i);
      if (tempMonth <= 0) {
        int diffYear = -tempMonth ~/ 12 + 1;
        tempYear = year - diffYear;
        // print('<0: diff year = $tempYear');
        tempMonth += 12 * diffYear;
        int idx = _pageModels.indexWhere((m) => m.month == tempMonth && m.year == tempYear);
        if (idx == -1) {
          List<CalendarModel> ms = mapToGridView(tempYear, tempMonth, firstWeekday: firstWeekday, fixedLines: fixedLines);
          CalendarPageModel pageModel = CalendarPageModel(year: tempYear, month: tempMonth, models: ms);
          _pageModels.insert(0, pageModel);
        }
        // print('<=0: year = $tempYear, month = $tempMonth');
      } else if (tempMonth > 12) {
        int diffYear = tempMonth ~/ 12;
        tempMonth -= 12 * diffYear - 1;
        tempYear = year + diffYear;
        // print('>12: year = $tempYear, month = $tempMonth');
        int idx = _pageModels.indexWhere((m) => m.month == tempMonth && m.year == tempYear);
        if (idx == -1) {
          List<CalendarModel> ms = mapToGridView(tempYear, tempMonth, firstWeekday: firstWeekday, fixedLines: fixedLines);
          CalendarPageModel pageModel = CalendarPageModel(year: tempYear, month: tempMonth, models: ms);
          _pageModels.add(pageModel);
        }
      } else {
        // print('same: year = $year, month = $tempMonth');
        int idx = _pageModels.indexWhere((m) => m.month == tempMonth && m.year == year);
        if (idx == -1) {
          List<CalendarModel> ms = mapToGridView(year, tempMonth, firstWeekday: firstWeekday, fixedLines: fixedLines);
          CalendarPageModel pageModel = CalendarPageModel(year: year, month: tempMonth, models: ms);
          if (_pageModels.isEmpty) {
            _pageModels.add(pageModel);
          } else {
            final firstPage = _pageModels.first;
            final lastPage = _pageModels.last;
            if (firstPage.year == lastPage.year) {
              if (year < firstPage.year) {
                _pageModels.insert(0, pageModel);
              } else if (lastPage.year > year) {
                _pageModels.add(pageModel);
              } else if (firstPage.year == year) {
                // print('firstPage.year == year: ${firstPage.year == year}');
                if (firstPage.month > tempMonth) {
                  // print('firstPage.month >= tempMonth: ${firstPage.month >= tempMonth}');
                  _pageModels.insert(0, pageModel);
                } else if (lastPage.month < tempMonth) {
                  // print('lastPage.month > tempMonth: ${lastPage.month > tempMonth}');
                  _pageModels.add(pageModel);
                }
              }
            } else {
              if (firstPage.year == year && firstPage.month > tempMonth) {
                _pageModels.insert(0, pageModel);
              } else if (lastPage.year == year && lastPage.month < tempMonth) {
                _pageModels.add(pageModel);
              }
            }
          }
        }
      }
    }
    return _pageModels;
  }

  // 日期转视图，可设置首日是否为星期日，适用非二维图如GridView
  List<CalendarModel> mapToCustomView(
    int year, 
    int month, 
    {
      int firstWeekday = 7, 
      bool fixedLines = true
    }
  ) {
    final numberOfLines = DateUtil.numberOfLinesOfMonth(year, month, fixedLines);
    int firstDayWeekday = DateUtil.firstWeekdayOfMonthForYearMonth(year, month);
    int curMonthDays = DateUtil.daysOfMonth(year, month);

    int preYear = year;
    int preMonth = month - 1;
    if (preMonth <= 0) {
      preMonth = 12;
      preYear--;
    }
    int nextYear = year;
    int nextMonth = month + 1;
    if (nextMonth > 12) {
      nextMonth = 12;
      nextYear++;
    }
    final totalNums = numberOfLines * 7;
    int preDays = diffWeekday(firstWeekday, firstDayWeekday);
    int preMonthDays = DateUtil.daysOfMonth(preYear, preMonth) - preDays;
    int tempCurMonthDays = curMonthDays + preDays;
    int currentMonthDayIdx = 1;
    
    String key = '';
    List<CalendarModel> models = [];
    for (var i = 0; i < totalNums; i++) {
      //pre month      
      if (i < preDays) {
        key = _key(preYear, preMonth, ++preMonthDays);
      } else if (i >= tempCurMonthDays) {
        // next month
        key = _key(nextYear, nextMonth, i-tempCurMonthDays+1);
      } else {
        key = _key(year, month, currentMonthDayIdx++);
      }
      if (key.isNotEmpty) {
        models.add(_modelsMap[key]);
      }
    }
    return models;
  }

}