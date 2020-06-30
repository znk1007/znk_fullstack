import 'package:cached_network_image/cached_network_image.dart';
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
  //已登录页面布局
  Widget _loginedWidget(BuildContext context) {
    print('scale height: ${Screen.scaleHeight}');
    return Container(
        color: Colors.red[500],
        height: Screen.setHeight(150).toDouble(),
        width: Screen.screenWidth,
        child: Stack(
          children: [
            Container(
              child: CachedNetworkImage(
                imageUrl: '',
                placeholder: (context, url) => Icon(Icons.person),
              ),
              width: Screen.setWidth(80),
              height: Screen.setWidth(80),
              color: Colors.orange,
            ),
            Container(),
          ],
        ));
  }
  //未登录页面
  Widget _unLoginedWidget(BuildContext context) {
    return Container();
  }
}
