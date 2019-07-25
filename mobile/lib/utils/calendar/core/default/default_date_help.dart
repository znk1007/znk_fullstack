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

  
  
}