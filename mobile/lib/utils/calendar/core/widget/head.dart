import 'package:flutter/material.dart';
import 'package:znk/utils/base/device.dart';
import 'package:znk/utils/calendar/core/data/util.dart';
import 'package:znk/utils/calendar/core/default/default_head.dart';
import 'package:znk/utils/calendar/custom/custom_head.dart';

class CalendarHead extends StatefulWidget {
  final CustomHead headTool;
  CalendarHead({Key key, this.headTool}) : super(key: key);

  _CalendarHeadState createState() => _CalendarHeadState();
}

class _CalendarHeadState extends State<CalendarHead> {
  CustomHead _headTool;
  @override
  Widget build(BuildContext context) {
    _headTool = widget.headTool ?? DefaultHead();

    List<Widget> children = [
      Expanded(
        child: _headTool.leftView,
      ),
      Expanded(
        flex: 2,
        child: _headTool.titleView,
      ),
      Expanded(
        child: _headTool.rightView,
      ),
    ];
    if (_headTool.statusView != null) {
      children = [
          Expanded(
            child: _headTool.leftView,
          ),
          Expanded(
            flex: 2,
            child: _headTool.titleView,
          ),
          Expanded(
            child: _headTool.statusView,
          ),
          Expanded(
            child: _headTool.rightView,
          )
      ];
    }
    return Container(
      width: Device.width,
      height: Device.relativeHeight(50),
       child: Stack(
         children: <Widget>[
           _headTool.backgroundView,
           Container(
             child: Row(
                children: children,
             )             
           ),
         ],
       ),
    );
  }
}