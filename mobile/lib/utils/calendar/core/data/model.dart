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

  int _firstYear = 1930;

  int _lastYear = 2100;
  // 预加载数据
  Future preLoad({int startYear = 1900, int endYear = 2100}) async {
    int s = startYear < _firstYear ? startYear : _firstYear;
    _firstYear = s;
    int e = endYear > _lastYear ? endYear : _lastYear;
    _lastYear = e;
    for (var i = _firstYear; i < _lastYear; i++) {
      
    }
  }

}