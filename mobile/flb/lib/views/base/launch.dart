import 'package:flb/util/random/random.dart';
import 'package:flutter/material.dart';

class LaunchPage extends StatelessWidget {
  const LaunchPage({Key key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      child: Text('启动页'),
      color: RandomHandler.randomColor,
    );
  }
}