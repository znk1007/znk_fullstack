import 'package:flutter/material.dart';
import 'package:znk/core/user/index.dart';

class Chat extends StatefulWidget {
  UserRepository _userRepository;
  Chat({Key key, @required UserRepository userRepository}) : 
    assert(userRepository != null),
    _userRepository = userRepository,
    super(key: key);
  _ChatState createState() => _ChatState();
}

class _ChatState extends State<Chat> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(
          '消息',
          style: TextStyle(
            color: Colors.black,
          ),
        ),
        centerTitle: true,
        backgroundColor: Colors.white,
      ),
    );
  }
}