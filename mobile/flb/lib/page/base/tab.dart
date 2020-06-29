import 'package:flutter/material.dart';
class TabPage extends StatefulWidget {
  final List<Widget> pages;
  TabPage({Key key, this.pages}) :  assert(pages != null), super(key: key);

  @override
  _TabPageState createState() => _TabPageState();
}

class _TabPageState extends State<TabPage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold();
  }
}