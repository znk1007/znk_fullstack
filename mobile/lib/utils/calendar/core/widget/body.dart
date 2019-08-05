import 'package:flutter/material.dart';
import 'package:flutter/src/rendering/sliver.dart';
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
  int _totalPage = 0;
  CustomDateHelper _helper;
  int _diffPages = 0;
  List<CalendarPageModel> _pageModels;
  double _calendarHeight = 268;
  double _totalHeight = 300;
  DateTime _currentTime = DateTime.now();
  
  @override
  void initState() {
    super.initState();
    _helper = widget.helper ?? DefaultDateHelper();
    _loadBase(_helper);
  }
  

  @override
  Widget build(BuildContext context) {
    _loadPageModels();
    return Container(
      child: Column(
        children: <Widget>[
          Container(
            width: Device.width,
            height: (_totalHeight - _calendarHeight).toDouble(),
          ),
          Container(
            width: Device.width,
            height: _calendarHeight,
            child: PageView.builder(
              itemCount: _totalPage,
              controller: _controller,
              itemBuilder: (BuildContext pageViewCtx, int pageIdx) {
                CalendarPageModel pageModel = _pageModels[pageIdx];
                return Container(
                  width: Device.width,
                  height: _calendarHeight,
                  child: GridView.builder(
                    itemBuilder: (BuildContext gridViewCtx, int gridIdx) {
                      List<CalendarModel> models = pageModel.models;
                      return Container();//models.isEmpty ? Container() : _gridItem(models[gridIdx]);
                    },
                    gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
                      crossAxisCount: 7,
                    ),
                  ),
                );
              },
              onPageChanged: (int current) {
                print('diff page: ${current-_diffPages}');
              },
            ),
          ),
        ],
      )
    );
  }
  @override
  void dispose() { 
    _controller.dispose();
    super.dispose();
  }

  // 加载基本数据
  void _loadBase(CustomDateHelper helper) {
      _totalPage = _helper.numberOfPage / 2 == 0 ? 3 : _helper.numberOfPage;
      int initPage = (_totalPage-1) ~/ 2;
      _diffPages = initPage;
      _controller = PageController(initialPage: (_totalPage-1) ~/ 2);
  }

  // 加载日历模型
  void _loadPageModels() {
    _pageModels = CalendarManager.instance.mapToGridViews(_currentTime.year, _currentTime.month, firstWeekday: _helper.firstWeekday, pages: _totalPage, keep: _helper.keepCache);
  }

  Widget _gridItem(CalendarModel model) {
    print('model date time: ${model.dateTime}');
    return Container(

    );
  }
}