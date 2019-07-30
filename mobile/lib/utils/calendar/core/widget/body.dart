import 'package:flutter/material.dart';
import 'package:znk/utils/calendar/core/data/model.dart';

class CalendarBody extends StatefulWidget {
  CalendarBody({Key key}) : super(key: key);

  _CalendarBodyState createState() => _CalendarBodyState();
}

class _CalendarBodyState extends State<CalendarBody> {
  PageController _controller;
  int _totalPage = 0;
  @override
  void initState() {
    super.initState();
    _controller = PageController(initialPage: CalendarManager.instance.currentPage);
    _totalPage = CalendarManager.instance.totalPage;
  }
  @override
  Widget build(BuildContext context) {
    return Container(
       child: PageView.builder(
         itemCount: _totalPage,
         controller: _controller,
         itemBuilder: (BuildContext ctx, int idx) {
           print('idx: $idx');
         },
         
       ),
    );
  }
}