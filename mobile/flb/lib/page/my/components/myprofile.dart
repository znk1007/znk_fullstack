import 'package:flb/model/user/user.dart';
import 'package:flb/util/screen/screen.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class MyProfile extends StatelessWidget {
  const MyProfile({Key key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    bool isLogined = context.watch<UserModel>().isLogined ?? false;
    return Container(
      child: isLogined ? _loginedWidget(context) : _unLoginedWidget(context),
    );
  }

  Widget _loginedWidget(BuildContext context) {
    print('scale height: ${Screen.scaleHeight}');
    return Container(
        color: Colors.red[500],
        height: Screen.setHeight(150).toDouble(),
        width: Screen.screenWidth,
        child: Stack(
          children: [
            
          ],
        ));
  }

  Widget _unLoginedWidget(BuildContext context) {
    return Container();
  }
}
