import 'package:flutter/material.dart';
import 'package:znk/utils/base/device.dart';
import 'package:znk/utils/calendar/core/data/model.dart';
import 'package:znk/utils/calendar/core/default/default_date_help.dart';
import 'package:znk/utils/calendar/custom/custom_date_help.dart';

class CalendarBody extends StatefulWidget {
  final CustomDateHelper helper;
  CalendarBody({Key key, this.helper}) : super(key: key);

  final state = _CalendarBodyState();

  _CalendarBodyState createState() => state;
}

class _CalendarBodyState extends State<CalendarBody> {

  PageController _controller;
  int _totalPage = 0;
  CustomDateHelper _helper;
  int _diffPages = 0;

  DateTime _now = DateTime.now();
  int _currentYear = 0;
  int _currentMonth = 0;
  
  @override
  void initState() {
    super.initState();
    _totalPage = CalendarManager.instance.totalPage;
    _helper = widget.helper ?? DefaultDateHelper();
    _totalPage = _helper.numberOfPage / 2 == 0 ? 3 : _helper.numberOfPage;
    int initPage = (_totalPage-1) ~/ 2;
    _diffPages = initPage;
    _controller = PageController(initialPage: (_totalPage-1) ~/ 2);
    _currentYear = _now.year;
    _currentMonth = _now.month;
  }
  @override
  Widget build(BuildContext context) {
    // final tests = [1,2,3,4,5,6];
    // tests.removeRange(4, tests.length-1);
    // print('');
    // for (var t in tests) {
    //   print('remove: $t');
    // }
    
    // int n = 100;
    // final models = CalendarManager.instance.mapToGridViews(2019, 1, pages: 2 * n + 1);
    // for (var m in models) {
    //   print('year: ${m.year}');
    //   print('month: ${m.month}');
    // }
    // for (var model in models) {
    //   print(' ');
    //   print('grid view: ${model.dateTime} weekday: ${model.dateTime.weekday} == column: ${model.column}, row: ${model.row}');
    // }
    // final models = CalendarManager.instance.mapToCustomView(2019, 7);
    // for (var model in models) {
    //   print('custom view: ${model.dateTime}');
    // }
    return Container(
      width: Device.width,
      height: 300,
      child: PageView.builder(
        itemCount: _totalPage,
        controller: _controller,
        itemBuilder: (BuildContext ctx, int idx) {
          print('idx: $idx');
          return Container(
            width: Device.width,
            height: 200,
            color: Colors.green,
          );
        },
        onPageChanged: (int current) {
          print('diff page: ${current-_diffPages}');
        },
      ),
    );
  }
  @override
  void dispose() { 
    _controller.dispose();
    super.dispose();
  }
}