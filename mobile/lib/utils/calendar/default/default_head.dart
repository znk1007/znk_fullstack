import 'dart:ui';

import 'package:flutter/material.dart';
import 'package:znk/utils/calendar/custom/custom_head.dart';

class DefaultHead implements CustomHead {

  @override
  Widget get backgroundView => Container(
    color: Colors.blue,
  );

  @override
  Widget get leftView => Container(
    child: Text('<'),
  );

  @override
  // TODO: implement leftViewPosition
  EdgeInsets get leftViewPosition => null;

  @override
  // TODO: implement rightView
  Widget get rightView => Container(
    child: Text('<'),
  );

  @override
  // TODO: implement rightViewPosition
  EdgeInsets get rightViewPosition => null;

  @override
  // TODO: implement statusView
  Widget get statusView => Container();

  @override
  // TODO: implement statusViewPostion
  EdgeInsets get statusViewPostion => null;

  @override
  // TODO: implement titleView
  Widget get titleView => Container(
    child: Text('<'),
  );

  @override
  // TODO: implement titleViewPosition
  EdgeInsets get titleViewPosition => null;
  
  
}