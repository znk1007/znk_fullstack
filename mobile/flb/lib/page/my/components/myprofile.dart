import 'dart:math';

import 'package:cached_network_image/cached_network_image.dart';
import 'package:flb/model/user/user.dart';
import 'package:flb/util/db/protos/generated/user/user.pb.dart';
import 'package:flb/util/screen/screen.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class MyProfile extends StatelessWidget {
  final double profileHeight;
  const MyProfile({Key key, this.profileHeight}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    UserModel userModel = context.watch<UserModel>();
    bool isLogined = userModel.isLogined ?? false;
    User user = userModel.currentUser;
    return Container(
      child: isLogined ? _loginedWidget(user) : _unLoginedWidget(context),
    );
  }

  //已登录页面布局
  Widget _loginedWidget(User user) {
    double avatarS = Screen.setWidth(60).toDouble();
    double height = max(Screen.setHeight(160).toDouble(), profileHeight);
    double lrW = Screen.setWidth(110).toDouble();
    double lrH = Screen.setHeight(40).toDouble();
    double avatarTop = (height - avatarS) / 2.0 - 10;
    double lrTop = (height - lrH) / 2.0 - 10;

    return Container(
        color: Colors.red[400],
        height: height,
        width: Screen.screenWidth,
        child: Stack(
          children: [
            Container(
              child: ClipOval(
                child: CachedNetworkImage(
                  fit: BoxFit.cover, //全圆
                  imageUrl:
                      'http://b-ssl.duitang.com/uploads/item/201509/22/20150922134955_vfEWL.jpeg',
                  placeholder: (context, url) => Icon(Icons.person),
                ),
              ),
              width: avatarS,
              height: avatarS,
              margin: EdgeInsets.only(left: 50, top: avatarTop),
            ),
            /* 适用本地资源 */
            // Container(
            //   width: avatarS,
            //   height: avatarS,
            //   margin: EdgeInsets.only(top: (height - avatarS) / 2.0, left: 50),
            //   decoration: BoxDecoration(
            //     shape: BoxShape.circle,
            //     color: Colors.blue,
            //   ),
            // ),
            Container(
              width: lrW,
              height: lrH,
              margin: EdgeInsets.only(top: lrTop, left: 50 + avatarS),
              child: FlatButton(
                  onPressed: () {
                    _loginOrRegist();
                  },
                  child: Text(
                    '登录/注册',
                    style: TextStyle(
                      color: Colors.white,
                      fontSize: 16,
                      fontWeight: FontWeight.bold,
                    ),
                  )),
            ),
            _privateEquality(context),
          ],
        ));
  }

  //已登录页面
  Widget _unLoginedWidget(BuildContext context) {
    return Container();
  }

  //未登录页面
  void _loginOrRegist() {
    print('登录/注册');
  }

  //积分、红包视图
  Widget _privateEquality(BuildContext context) {
    double height = Screen.setWidth(80).toDouble();
    return Container(
      height: height,
      width: Screen.screenWidth,
      decoration: BoxDecoration(
        color: Colors.white,
        borderRadius: BorderRadius.all(Radius.circular(height / 2.0)),
      ),
      child: Row(
        children: [],
      ),
    );
  }
}
