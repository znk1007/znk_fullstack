import '../other/utils/tools/screen_helper.dart';
import '../view/default/default_login_view.dart';
import '../view/login_view_delegate.dart';
import 'package:flutter/material.dart';

class LoginPage<T extends LoginViewDelegate> extends StatefulWidget {
  final T value;
  LoginPage({
    @required this.value,
  }) {
    ScreenHelper.setDesignParams(414, 736);
  }

  @override
  _LoginPageState createState() => _LoginPageState();
}
class _LoginPageState extends State<LoginPage> {

  @override
  Widget build(BuildContext context) {
    ScreenHelper.setContext(context);
    var child = (widget.value != null) ? widget.value : DefaultLoginView();

    return Scaffold(
      body: child,
    );
  }

  void _handleLoginWidget(LoginViewDelegate delegate) {
    
  }
}