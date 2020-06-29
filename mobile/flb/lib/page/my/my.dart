import 'package:flb/util/random/color.dart';
import 'package:flutter/material.dart';

class MyPage extends StatelessWidget {

  static String id = 'my';

  const MyPage({Key key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      child: Text('我的'),
      color: RandomHandler.randomColor,
    );
  }
}