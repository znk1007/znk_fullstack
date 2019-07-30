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

  int get totalPage {
    return _pages;
  }

  Map<String, CalendarModel> _selectedModels;

  // 模型
  Map<String, CalendarModel> _modelsMap;
  // 当前页码
  int currentPage = 0;
  // 总页数
  int _pages = 0;

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
    _pages = (endYear - startYear) * 12;
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

  List<CalendarModel> load(int startYear, int endYear) {

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