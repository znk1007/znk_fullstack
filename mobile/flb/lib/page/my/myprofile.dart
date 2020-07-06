import 'package:cached_network_image/cached_network_image.dart';
import 'package:flb/model/style/mystyle.dart';
import 'package:flb/model/user/user.dart';
import 'package:flb/page/my/model/my.dart';
import 'package:flb/pkg/screen/screen.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class MyProfileView extends StatelessWidget {
  final MyPageStyle style;
  final UserModel userModel;
  final MyCompany company;
  const MyProfileView({Key key, this.style, this.userModel, this.company})
      : assert(style != null),
        assert(userModel != null),
        super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      height: style.profileBgHeight,
      width: style.profileBgWidth,
      color: style.profileBgColor,
      child: Stack(
        children: [
          //头像
          GestureDetector(
            child: Container(
              child: ClipOval(
                child: (userModel.currentUser != null &&
                        userModel.currentUser.photo.length > 0)
                    ? CachedNetworkImage(
                        color: Colors.white,
                        fit: BoxFit.cover, //全圆
                        imageUrl: (userModel.currentUser != null &&
                                userModel.currentUser.photo.length > 0)
                            ? userModel.currentUser.photo
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
            onTap: () {
              print('header tap');
            },
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
                  if (userModel.isLogined) {
                    return;
                  }
                  _loginOrRegist();
                },
                child: userModel.isLogined
                    ? (company.info != null && company.info.name.length > 0)
                        ? Container(
                            child: Column(
                              children: [
                                Text(
                                  userModel.currentUser.nickname.length > 0
                                      ? userModel.currentUser.nickname
                                      : '昵称',
                                  style: TextStyle(
                                      color: Colors.white,
                                      fontSize: 16,
                                      fontWeight: FontWeight.w600),
                                  textAlign: TextAlign.left,
                                ),
                                Text(
                                  company.info.name,
                                  style: TextStyle(
                                      color: Colors.grey[200],
                                      fontSize: 13,
                                      fontWeight: FontWeight.w600),
                                  textAlign: TextAlign.left,
                                ),
                              ],
                            ),
                          )
                        : Container(
                            width: style.profileBgWidth -
                                (style.avatarMargin.left + style.avatarL) -
                                10,
                            child: Text(
                              userModel.currentUser.nickname.length > 0
                                  ? userModel.currentUser.nickname
                                  : '昵称',
                              style: TextStyle(
                                  color: Colors.white,
                                  fontSize: 16,
                                  fontWeight: FontWeight.w600),
                              textAlign: TextAlign.left,
                            ),
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
                borderRadius:
                    BorderRadius.all(Radius.circular(style.eqHeight / 2.0))),
            child: Row(
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
                            child: Text(company.info != null
                                ? company.info.integ
                                : '0'),
                          ),
                          Container(
                            margin: EdgeInsets.only(left: Screen.setWidth(30)),
                            child: Text(userModel.isLogined ? '我的积分' : '积分'),
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
                            child: Text(company.info != null
                                ? company.info.redPack
                                : '0'),
                          ),
                          Container(
                            margin: EdgeInsets.only(right: Screen.setWidth(30)),
                            child: Text(userModel.isLogined ? '我的红包' : '红包'),
                          )
                        ],
                      )),
                )
              ],
            ),
          ),
        ],
      ),
    );
  }

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
