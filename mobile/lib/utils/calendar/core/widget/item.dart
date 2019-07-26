import 'package:flutter/material.dart';
import 'package:znk/utils/calendar/core/data/model.dart';
import 'package:znk/utils/calendar/custom/custom_date_help.dart';

class CalendarItem extends StatefulWidget {
  CustomDateHelper dateHelper;
  CalendarModel model;
  CalendarItem({Key key, @required this.dateHelper, @required this.model}) : super(key: key);

  _CalendarItemState createState() => _CalendarItemState();
}

class _CalendarItemState extends State<CalendarItem> {
  @override
  Widget build(BuildContext context) {
    return Container(
       child: Container(),
    );
  }
}