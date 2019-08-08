import 'package:flutter/material.dart';
import 'package:flutter/src/rendering/sliver_grid.dart';
import 'package:znk/utils/base/device.dart';
import 'package:znk/utils/calendar/core/data/model.dart';
import 'package:znk/utils/calendar/core/data/util.dart';
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
  CustomDateHelper _helper;
  int _diffPages = 0;
  List<CalendarPageModel> _pageModels;
  double _calendarHeight = 368;
  double _totalHeight = 400;
  DateTime _currentTime = DateTime.now();

  CalendarModel _currentModel = CalendarModel()..dateTime = DateTime.now();

  List<int> _weekdays;
  
  @override
  void initState() {
    super.initState();
    _helper = widget.helper ?? DefaultDateHelper();
    _loadBase(_helper);
    _loadPageModels();
  }
  

  @override
  Widget build(BuildContext context) {
    return Container(
      height: _totalHeight,
      child: Column(
        children: <Widget>[
          Container(
            width: Device.width,
            height: _totalHeight - _calendarHeight,
            child: GridView.builder(
              physics: NeverScrollableScrollPhysics(),
              itemBuilder: (BuildContext weekdayCtx, int weekdayIdx) {
                return _weekdayItem(_weekdays[weekdayIdx], weekdayIdx);
              },
              gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
                crossAxisCount: _weekdays.length,
                childAspectRatio: 1.8,
              ),
              itemCount: _weekdays.length,
            ),
          ),
          Container(
            width: Device.width,
            height: _calendarHeight,
            child: PageView.builder(
              itemCount: _helper.numberOfPage,
              controller: _controller,
              itemBuilder: (BuildContext pageViewCtx, int pageIdx) {
                CalendarPageModel pageModel = _pageModels[pageIdx];
                return GridView.builder(
                  itemBuilder: (BuildContext gridViewCtx, int gridIdx) {
                    List<CalendarModel> models = pageModel.models;
                    return _gridCalendarItem(models[gridIdx]);
                  },
                  physics: NeverScrollableScrollPhysics(),
                  gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
                    crossAxisCount: _weekdays.length,
                    childAspectRatio: 1.2,
                  ),
                  itemCount: pageModel.models.length,
                );
              },
              onPageChanged: (int current) {
                setState(() {
                  int diff = current - _diffPages;
                  int tempDiff = DateUtil.abs(diff);
                  print('temp diff: $tempDiff');
                  print('_helper.numberOfPage - tempDiff: ${_helper.numberOfPage - tempDiff}');
                  _currentModel = CalendarModel()..dateTime = DateUtil.diffMonths(_currentTime, diff);
                });
              },
            ),
          ),
        ],
      )
    );. ·
  }
  @override
  void dispose() { 
    _controller.dispose();
    super.dispose();
  }

  // 加载基本数据
  void _loadBase(CustomDateHelper helper) {
      int totalPage = _helper.numberOfPage / 2 == 0 ? 3 : _helper.numberOfPage;
      int initPage = (totalPage-1) ~/ 2;
      _diffPages = initPage;
      _controller = PageController(initialPage: initPage);
      _weekdays = DateUtil.weekdays(firstWeekday: helper.firstWeekday);
  }

  // 加载日历模型
  void _loadPageModels() {
    _pageModels = CalendarManager.instance.mapToGridViews(_currentModel.dateTime.year, _currentModel.dateTime.month, firstWeekday: _helper.firstWeekday, pages: _helper.numberOfPage, keep: _helper.keepCache);
  }

  Widget _weekdayItem(int weekday, int idx) {
    int len = _weekdays.length;
    return Container(
      width: Device.width / len,
      height: _totalHeight - _calendarHeight,
      decoration: BoxDecoration(
        border: _helper.weekdaySeparatorWidth != 0 ? Border(
          right: BorderSide(
            color: (idx != len - 1) ? _helper.weekdaySeparatorColor : Colors.white,
          ),
        ) : null,
      ),
      alignment: Alignment.center,
      child: Text(
        '${DateUtil.cnWeek(weekday, _helper.weekType)}',
        textAlign: TextAlign.center,
      ),
    );
  }

  // 日历具体内容
  Widget _gridCalendarItem(CalendarModel model) {
    return Container(
      alignment: Alignment.center,
      child: Stack(
        children: <Widget>[
          Text(
            '${model.dateTime.day}',
            style: TextStyle(
              color: _currentModel.isSameMonth(model) ? _helper.innerDayColor : _helper.outerDayColor,
            ),
          ),
        ],
      ),
    );
  }
}