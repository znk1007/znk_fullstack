import 'package:flutter/material.dart';
import 'package:znk/core/user/register/index.dart';
import 'package:znk/core/user/user_repository.dart';

class RegisterPage extends StatelessWidget {
  static const String routeName = "/register";
  final UserRepository _userRepository;

  RegisterPage({Key key, @required UserRepository userRepository}):
    assert(userRepository != null),
    _userRepository = userRepository,
    super(key: key);
    
  @override
  Widget build(BuildContext context) {
    var _registerBloc = new RegisterBloc();
    return new Scaffold(
      appBar: new AppBar(
        title: new Text("Register"),
      ),
      body: null,//new RegisterScreen(registerBloc: _registerBloc),
    );
  }
}
