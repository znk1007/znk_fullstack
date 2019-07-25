import 'dart:ui';

import 'package:flutter/material.dart';
import 'package:znk/utils/base/device.dart';
import 'package:znk/utils/calendar/custom/custom_head.dart';

class DefaultHead implements CustomHead {
  String _statusText = '';
  String _title = 'data';
  final _titleWidth = Device.relativeWidth(150);
  @override
  Widget get backgroundView => Container(
    color: Colors.grey[100],
  );

  @override
  Widget get leftView => Container(
    child: FlatButton(
      child: Icon(
        Icons.arrow_left,
        color: Colors.black,
      ),
      onPressed: () {
        print('left view press');
      },
    ),
  );

  @override
  Widget get rightView => Container(
    child: FlatButton(
      child: Icon(
        Icons.arrow_right,
        color: Colors.black
      ),
      onPressed: () {
        print('right view press');
      },
    ),
  );


  @override
  Widget get statusView => null;


  @override
  Widget get titleView => Container(
    alignment: Alignment.center,
    child: Text(
      _title
    ),
  );


  

  @override
  set statuText(String txt) {
    _statusText = txt ?? '';
  }

  @override
  set title(String txt) {
    _title = txt ?? '';
  }
  
  
}