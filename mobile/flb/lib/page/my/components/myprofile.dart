import 'package:cached_network_image/cached_network_image.dart';
import 'package:flb/model/style/mystyle.dart';
import 'package:flb/model/user/user.dart';
import 'package:flb/page/my/model/my.dart';
import 'package:flb/util/screen/screen.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class MyProfile extends StatelessWidget {
  const MyProfile({Key key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    MyPageStyle style = context.watch<MyPageStyle>();
    return Container(
      height: style.profileBgHeight,
      width: style.profileBgWidth,
      color: style.profileBgColor,
      child: Consumer<UserModel>(
        builder: (ctx, u, child) {
          return Stack(
            children: [
              //头像
              Container(
                child: ClipOval(
                  child:
                      (u.currentUser != null && u.currentUser.photo.length > 0)
                          ? CachedNetworkImage(
                              color: Colors.white,
                              fit: BoxFit.cover, //全圆
                              imageUrl: (u.currentUser != null &&
                                      u.currentUser.photo.length > 0)
                                  ? u.currentUser.photo
                                  : '',
                              placeholder: (context, url) =>
                                  Icon(Icons.person_outline),
                            )
                          : Container(
                              child: Icon(Icons.person_outline),
                            ),
                ),
                width: style.avatarL,
                height: style.avatarL,
                margin: style.avatarMargin,
                decoration: BoxDecoration(
                  color: Colors.white,
                  borderRadius: BorderRadius.all(
                    Radius.circular(style.avatarL / 2.0),
                  ),
                ),
              ),
              //昵称、登录/注册
              Container(
                height: style.nicknameHeight,
                width: style.profileBgWidth -
                    (style.avatarMargin.left + style.avatarL) -
                    10,
                margin: EdgeInsets.only(
                    left: (style.avatarMargin.left + style.avatarL + 5),
                    top: (style.avatarMargin.top +
                        (style.avatarL - style.nicknameHeight) / 2.0)),
                child: FlatButton(
                    onPressed: () {
                      if (u.isLogined) {
                        return;
                      }
                      _loginOrRegist();
                    },
                    child: u.isLogined
                        ? Column(
                            children: [
                              Text(
                                u.currentUser.nickname.length > 0
                                    ? u.currentUser.nickname
                                    : '昵称',
                                style: TextStyle(
                                    color: Colors.white,
                                    fontSize: 16,
                                    fontWeight: FontWeight.w600),
                                textAlign: TextAlign.left,
                              )
                            ],
                          )
                        : Container(
                            width: style.profileBgWidth -
                                (style.avatarMargin.left + style.avatarL) -
                                10,
                            child: Text(
                              '登录/注册',
                              style: TextStyle(
                                  color: Colors.white,
                                  fontSize: 16,
                                  fontWeight: FontWeight.w600),
                              textAlign: TextAlign.left,
                            ),
                          )),
              ),
              //收益
              Container(
                width: style.profileBgWidth - 10,
                height: style.eqHeight,
                margin: EdgeInsets.only(
                  left: 5,
                  top: style.avatarMargin.top + style.avatarL + 20,
                ),
                decoration: BoxDecoration(
                    color: Colors.white,
                    borderRadius: BorderRadius.all(
                        Radius.circular(style.eqHeight / 2.0))),
                child: Consumer<MyModelHandler>(
                  builder: (context, m, child) {
                    return Row(
                      children: [
                        Container(
                          width: (style.profileBgWidth - 10) / 2.0,
                          child: FlatButton(
                              onPressed: () {},
                              child: Column(
                                children: [
                                  Container(
                                    margin: EdgeInsets.only(
                                        left: Screen.setWidth(30), top: 8),
                                    child: Text(m.company != null
                                        ? m.company.integ
                                        : '0'),
                                  ),
                                  Container(
                                    margin: EdgeInsets.only(
                                        left: Screen.setWidth(30)),
                                    child: Text(u.isLogined ? '我的积分' : '积分'),
                                  ),
                                ],
                              )),
                        ),
                        Container(
                          width: (style.profileBgWidth - 10) / 2.0,
                          child: FlatButton(
                              onPressed: () {},
                              child: Column(
                                children: [
                                  Container(
                                    margin: EdgeInsets.only(
                                        top: 8, right: Screen.setWidth(30)),
                                    child: Text(m.company != null
                                        ? m.company.redPack
                                        : '0'),
                                  ),
                                  Container(
                                    margin: EdgeInsets.only(
                                        right: Screen.setWidth(30)),
                                    child: Text(u.isLogined ? '我的红包' : '红包'),
                                  )
                                ],
                              )),
                        )
                      ],
                    );
                  },
                ),
              ),
            ],
          );
        },
      ),
    );
  }

  // //已登录页面布局
  // Widget _loginedWidget(BuildContext context, UserModel userModel) {
  //   double avatarS = Screen.setWidth(60).toDouble();
  //   double height = max(Screen.setHeight(200).toDouble(), 200);
  //   double lrW = Screen.setWidth(110).toDouble();
  //   double lrH = Screen.setHeight(40).toDouble();
  //   double avatarTop = (height - avatarS) / 2.0 - 20;
  //   double lrTop = (height - lrH) / 2.0 - 10;
  //   MyModelHandler handler = context.watch<MyModelHandler>();
  //   double eqWidth = Screen.screenWidth;
  //   double eqHeight = Screen.setHeight(50).toDouble();

  //   List<Widget> eqChildren = handler.equalitys
  //       .map((e) => _equalityWidget(e.number, e.title,
  //           e.offsetRadio * (e.widthRadio * eqWidth), (e.widthRadio * eqWidth)))
  //       .toList();
  //   return Container(
  //       color: Colors.red[400],
  //       height: height,
  //       width: Screen.screenWidth,
  //       child: Stack(
  //         children: [
  //           //头像
  //           Container(
  //             child: Consumer<UserModel>(
  //               builder: (context, t, child) {
  //                 return ClipOval(
  //                   child: CachedNetworkImage(
  //                     color: Colors.white,
  //                     fit: BoxFit.cover, //全圆
  //                     imageUrl: (t.currentUser != null &&
  //                             t.currentUser.photo.length > 0)
  //                         ? t.currentUser.photo
  //                         : '',
  //                     placeholder: (context, url) => Icon(Icons.person),
  //                   ),
  //                 );
  //               },
  //             ),
  //             width: avatarS,
  //             height: avatarS,
  //             margin: EdgeInsets.only(left: 50, top: avatarTop),
  //             color: Colors.white,
  //           ),
  //           //登录/注册
  //           Container(
  //             width: lrW,
  //             height: lrH,
  //             margin: EdgeInsets.only(top: lrTop, left: 50 + avatarS),
  //             child: FlatButton(
  //                 onPressed: () {
  //                   _loginOrRegist();
  //                 },
  //                 child: Text(
  //                   '登录/注册',
  //                   style: TextStyle(
  //                     color: Colors.white,
  //                     fontSize: 16,
  //                     fontWeight: FontWeight.bold,
  //                   ),
  //                 )),
  //           ),
  //           //收益
  //           Container(
  //             height: eqHeight,
  //             width: eqWidth,
  //             margin: EdgeInsets.only(top: avatarTop + avatarS + 20),
  //             decoration: BoxDecoration(
  //               color: Colors.white,
  //               borderRadius: BorderRadius.all(Radius.circular(height / 2.0)),
  //             ),
  //             child: Row(
  //               children: eqChildren,
  //             ),
  //           ),
  //         ],
  //       ));
  // }

  //未登录页面
  void _loginOrRegist() {
    print('登录/注��');
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

/*
'http://b-ssl.duitang.com/uploads/item/201509/22/20150922134955_vfEWL.jpeg'
*/
