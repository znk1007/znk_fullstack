import 'package:flutter/material.dart';
import 'package:znk/core/user/index.dart';
import 'package:znk/utils/calendar/core/data/model.dart';
import 'package:znk/utils/calendar/core/data/util.dart';
import 'package:znk/utils/calendar/core/widget/head.dart';

class Schedule extends StatefulWidget {
  UserRepository _userRepository;
  Schedule({Key key, @required UserRepository userRepository}) : 
    assert(userRepository != null),
    _userRepository = userRepository,
    super(key: key);
  _ScheduleState createState() => _ScheduleState();
}

class _ScheduleState extends State<Schedule> {
  
  @override
  Widget build(BuildContext context) {
    final models = CalendarManager.instance.getModels(2019, 7, 29, 7);
    if (models != null) {
      for (var m in models) {
        print('year ${m.year}');
        print('month ${m.month}');
        print('day ${m.day}');
      }
    }
    return Scaffold(
      appBar: AppBar(
        elevation: 0,
        title: Text(
          '日程',
          style: TextStyle(
            color: Colors.black,
          ),
        ),
        centerTitle: true,
        backgroundColor: Colors.white,
      ),
      body: Container(
        child: Column(
          children: <Widget>[
            CalendarHead(),
          ],
        ),
      ),
    );
  }
}