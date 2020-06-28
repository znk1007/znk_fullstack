import 'package:flutter/material.dart';

class ClassifyPage extends StatefulWidget {
  ClassifyPage({Key key}) : super(key: key);

  @override
  _ClassifyPageState createState() => _ClassifyPageState();
}

class _ClassifyPageState extends State<ClassifyPage> {
  @override
  Widget build(BuildContext context) {
    return Container(
       child: Text('分类'),
    );
  }
}