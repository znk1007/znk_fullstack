import 'package:flutter/material.dart';
import 'package:znk/core/user/index.dart';

class Owner extends StatefulWidget {
  UserRepository _userRepository;
  Owner({Key key, @required UserRepository userRepository}) : 
  assert(userRepository != null),
  _userRepository = userRepository,
  super(key: key);

  _OwnerState createState() => _OwnerState();
}

class _OwnerState extends State<Owner> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('我的'),
        centerTitle: true,
      ),
    );
  }
}