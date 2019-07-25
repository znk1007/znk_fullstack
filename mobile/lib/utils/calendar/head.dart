import 'package:flutter/material.dart';
import 'package:znk/utils/base/device.dart';
import 'package:znk/utils/calendar/custom/custom_head.dart';
import 'package:znk/utils/calendar/default/default_head.dart';

class CalendarHead extends StatefulWidget {
  CustomHead headTool;
  CalendarHead({Key key, CustomHead headTool}) : 
    this.headTool = headTool == null ? DefaultHead() : headTool,
    super(key: key);

  _CalendarHeadState createState() => _CalendarHeadState();
}

class _CalendarHeadState extends State<CalendarHead> {
  @override
  Widget build(BuildContext context) {
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