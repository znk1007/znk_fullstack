import 'package:flutter/material.dart';
import 'package:znkauth/viewmodel/network/grpc/client.dart';
import '../viewmodel/db/usertbl.dart';
import '../model/delegate/auth.dart';
/* 回调 */
typedef Callback = void Function(bool succ, String msg) ;
class AuthPage extends StatefulWidget {
  /* 配置 */
  final ZnkAuthConfig config;
  /* 数据库 */
  final UserTBL userTBL;
  AuthPage({Key key, this.config, this.userTBL}) : super(key: key);

  @override
  _AuthPageState createState() => _AuthPageState();
}

class _AuthPageState extends State<AuthPage> {
  @override
  Widget build(BuildContext context) {
    ZnkAuthRpc.shared.setRpc(useTls: true, useTlsCA: true, host: 'localhost', port: 8080, ctx: context);
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