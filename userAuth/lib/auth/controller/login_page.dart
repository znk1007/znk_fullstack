import 'package:flutter/material.dart';
import 'package:userAuth/auth/default/default_login_view.dart';

import '../delegate/login_view_delegate.dart';

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
    var child = (widget._value != null) ? widget._value : DefaultLoginView();

    return Container(
       child: child,
    );
  }

  void _handleLoginWidget(LoginViewDelegate delegate) {
    
  }
}