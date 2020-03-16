import 'package:flutter/material.dart';
import 'package:znk_auth/znk_auth.dart';

/* 回调 */
typedef Callback = void Function(bool succ, String msg) ;
class AuthPage extends StatefulWidget {
  final ZnkAuthConfig config;
  AuthPage({Key key, this.config}) : super(key: key);

  @override
  _AuthPageState createState() => _AuthPageState();
}

class _AuthPageState extends State<AuthPage> {
  @override
  Widget build(BuildContext context) {
    if (widget.config == null) {
      return GestureDetector(
        child: Container(
          child:Text("缺少配置参数"),
          alignment: Alignment.center,
        ),
        onTap: () {
          if (Navigator.canPop(context)) {
            Navigator.pop(context);
          }
        },
      );
    }
    return Container(
       child: Text('data'),
    );
  }
}