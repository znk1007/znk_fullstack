import 'package:flutter/material.dart';

import 'delegates/login_view_delegate.dart';

class DefaultLoginPage implements LoginViewDelegate {
  @override
  // TODO: implement account
  String get account => throw UnimplementedError();

  @override
  // TODO: implement keepPassword
  bool get keepPassword => throw UnimplementedError();

  @override
  // TODO: implement password
  String get password => throw UnimplementedError();

  @override
  // TODO: implement verifyCode
  String get verifyCode => throw UnimplementedError();
  
}

class LoginPage<T extends LoginViewDelegate> extends StatefulWidget {
  T _value;
  LoginPage({
    Key key,
    T value,
  }) {
    _value = value;
  }

  @override
  _LoginPageState createState() => _LoginPageState();
}
class _LoginPageState extends State<LoginPage> {
  @override
  Widget build(BuildContext context) {
    widget._value.account;

    return Container(
       child: child,
    );
  }
}