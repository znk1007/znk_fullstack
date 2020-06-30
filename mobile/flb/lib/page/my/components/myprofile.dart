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
    num height = Screen.setHeight(120);
    print('scale height: ${Screen.scaleHeight}');
    return Container(
      child: Text('头部'),
      color: Colors.red[500],
      height: 150,
      width: Screen.screenWidth,
    );
  }

  Widget _unLoginedWidget(BuildContext context) {
    return Container();
  }
}
