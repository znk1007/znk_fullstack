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
  bool isSchedule;// 是否有事项
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
  
  

}