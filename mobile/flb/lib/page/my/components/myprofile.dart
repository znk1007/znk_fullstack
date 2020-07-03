import 'dart:math';

import 'package:cached_network_image/cached_network_image.dart';
import 'package:flb/model/user/user.dart';
import 'package:flb/page/my/model/my.dart';
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
    return Container(
      child: isLogined
          ? _loginedWidget(context, userModel)
          : _unLoginedWidget(userModel),
    );
  }

  //已登录页面布局
  Widget _loginedWidget(BuildContext context, UserModel userModel) {
    double avatarS = Screen.setWidth(60).toDouble();
    double height = max(Screen.setHeight(200).toDouble(), profileHeight);
    double lrW = Screen.setWidth(110).toDouble();
    double lrH = Screen.setHeight(40).toDouble();
    double avatarTop = (height - avatarS) / 2.0 - 20;
    double lrTop = (height - lrH) / 2.0 - 10;
    MyModelHandler handler = context.watch<MyModelHandler>();
    double eqWidth = Screen.screenWidth;
    double eqHeight = Screen.setHeight(50).toDouble();

    List<Widget> eqChildren = handler.equalitys
        .map((e) => _equalityWidget(e.number, e.title,
            e.offsetRadio * (e.widthRadio * eqWidth), (e.widthRadio * eqWidth)))
        .toList();
    return Container(
        color: Colors.red[400],
        height: height,
        width: Screen.screenWidth,
        child: Stack(
          children: [
            //头像
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
            //登录/注册
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
            //收益
            Container(
              height: eqHeight,
              width: eqWidth,
              margin: EdgeInsets.only(top: avatarTop + avatarS + 20),
              decoration: BoxDecoration(
                color: Colors.white,
                borderRadius: BorderRadius.all(Radius.circular(height / 2.0)),
              ),
              child: Row(
                children: eqChildren,
              ),
            ),
          ],
        ));
  }

  //未登录页面
  Widget _unLoginedWidget(UserModel userModel) {
    return Container();
  }

  //未登录页面
  void _loginOrRegist() {
    print('登录/注册');
  }

  //收益模块
  Widget _equalityWidget(
      String number, String title, double left, double width) {
    return Container(
      width: width,
      margin: EdgeInsets.only(left: left),
      child: Column(
        children: [
          Text(number),
          Text(title),
        ],
      ),
    );
  }
}
