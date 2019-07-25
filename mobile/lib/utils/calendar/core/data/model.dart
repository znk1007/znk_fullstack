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
    return dateTime.difference(DateTime.now()).inDays == 0; 
  }
  bool isSelected;
  bool isSchedule;// 是否有事项
}