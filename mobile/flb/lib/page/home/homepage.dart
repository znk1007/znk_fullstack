import 'package:flb/util/random/color.dart';
import 'package:flutter/material.dart';

class HomePage extends StatefulWidget {

  static String id = 'home';

  HomePage({Key key}) : super(key: key);

  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  @override
  Widget build(BuildContext context) {
    return Container(
       child: Text('首页'),
       color: RandomHandler.randomColor,
    );
  }
}