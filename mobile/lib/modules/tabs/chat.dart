import 'package:flutter/material.dart';

class Chat extends StatefulWidget {
  Chat({Key key}) : super(key: key);

  _ChatState createState() => _ChatState();
}

class _ChatState extends State<Chat> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('消息'),
        centerTitle: true,
      ),
    );
  }
}