import 'dart:ui';

import 'package:flutter/material.dart';
import 'package:znk/utils/calendar/custom/custom_head.dart';

class DefaultHead implements CustomHead {

  @override
  Widget get backgroundView => Container(
    color: Colors.grey[100],
  );

  @override
  Widget get leftView => Container(
    margin: leftViewPosition,
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
  EdgeInsets get leftViewPosition => EdgeInsets.only(left: 20);

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
  EdgeInsets get rightViewPosition => null;

  @override
  Widget get statusView => Container();


  @override
  EdgeInsets get statusViewPostion => null;

  @override
  Widget get titleView => Container(
    child: Text(
      title
    ),
  );

  @override
  String get title => 'date';

  @override
  EdgeInsets get titleViewPosition => null;

  @override
  String statusText;

  @override
  void set statuText(String txt) {
    // TODO: implement statuText
  }

  @override
  void set title(String txt) {
    // TODO: implement title
  }
  
  
}