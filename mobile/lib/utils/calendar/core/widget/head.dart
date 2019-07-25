import 'package:flutter/material.dart';
import 'package:znk/utils/base/device.dart';
import 'package:znk/utils/calendar/core/data/util.dart';
import 'package:znk/utils/calendar/core/default/default_head.dart';
import 'package:znk/utils/calendar/custom/custom_head.dart';

class CalendarHead extends StatefulWidget {
  CustomHead headTool;
  CalendarHead({Key key, this.headTool}) : super(key: key);

  _CalendarHeadState createState() {
    this.headTool = this.headTool ?? DefaultHead();
    
    return _CalendarHeadState();
  }
}

class _CalendarHeadState extends State<CalendarHead> {
  @override
  Widget build(BuildContext context) {
    DateUtil.chineseWeek(DateTime.now());
    List<Widget> children = [
      Expanded(
        child: widget.headTool.leftView,
      ),
      Expanded(
        flex: 2,
        child: widget.headTool.titleView,
      ),
      Expanded(
        child: widget.headTool.rightView,
      ),
    ];
    if (widget.headTool.statusView != null) {
      children = [
          Expanded(
            child: widget.headTool.leftView,
          ),
          Expanded(
            flex: 2,
            child: widget.headTool.titleView,
          ),
          Expanded(
            child: widget.headTool.statusView,
          ),
          Expanded(
            child: widget.headTool.rightView,
          )
      ];
    }
    return Container(
      width: Device.width,
      height: Device.relativeHeight(50),
       child: Stack(
         children: <Widget>[
           widget.headTool.backgroundView,
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