import 'package:znk/utils/calendar/custom/custom_date_help.dart';

class DefaultDateHelper implements CustomDateHelper {
  @override
  WeekType get weekType => WeekType.single;

  @override
  set weekType(WeekType _type) {
    weekType = _type;
  }

  @override
  MonthType get monthType => MonthType.short;

  @override
  set monthType(MonthType _monthType) {
    monthType = _monthType;
  }

  @override
  int get firstWeekday => 7;

  @override
  int get numberOfPage => 5;

  @override
  bool get keepCache => true;

  
}