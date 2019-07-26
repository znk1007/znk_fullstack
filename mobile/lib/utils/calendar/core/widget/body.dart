import 'package:flutter/material.dart';
import 'package:znk/utils/calendar/core/data/model.dart';

class CalendarBody extends StatefulWidget {
  CalendarBody({Key key}) : super(key: key);

  _CalendarBodyState createState() => _CalendarBodyState();
}

class _CalendarBodyState extends State<CalendarBody> {
  List <CalendarModel> models;
  @override
  void initState() {
    super.initState();
    models = [];
  }
  @override
  Widget build(BuildContext context) {
    return Container(
       child: PageView.builder(
         itemBuilder: (BuildContext ctx, int idx) {

         },
         itemCount: 1,
       ),
    );
  }
}