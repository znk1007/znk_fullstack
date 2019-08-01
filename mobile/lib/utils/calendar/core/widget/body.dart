import 'package:flutter/material.dart';
import 'package:znk/utils/base/device.dart';
import 'package:znk/utils/calendar/core/data/model.dart';
import 'package:znk/utils/calendar/core/default/default_date_help.dart';
import 'package:znk/utils/calendar/custom/custom_date_help.dart';

class CalendarBody extends StatefulWidget {
  final CustomDateHelper helper;
  CalendarBody({Key key, this.helper}) : super(key: key);

  _CalendarBodyState createState() => _CalendarBodyState();
}

class _CalendarBodyState extends State<CalendarBody> {

  PageController _controller;
  int _totalPage = 0;
  CustomDateHelper _helper;
  
  @override
  void initState() {
    super.initState();
    _totalPage = CalendarManager.instance.totalPage;
    _helper = widget.helper ?? DefaultDateHelper();
    _totalPage = _helper.numberOfPage / 2 == 0 ? 3 : _helper.numberOfPage;
    _controller = PageController(initialPage: (_totalPage-1) ~/ 2);
    
  }
  @override
  Widget build(BuildContext context) {
    CalendarManager.instance.load(DateTime(2018,10,31), DateTime(2019, 9, 30));
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
          print('current page: $current');
        },
      ),
    );
  }
}