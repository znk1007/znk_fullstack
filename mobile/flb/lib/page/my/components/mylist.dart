import 'package:flb/model/style/mystyle.dart';
import 'package:flb/model/user/user.dart';
import 'package:flb/page/my/model/my.dart';
import 'package:flb/pkg/screen/screen.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class MyListView extends StatelessWidget {
  const MyListView({Key key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    MyPageStyle style = context.watch<MyPageStyle>();
    return Container(
      height: Screen.screenHeight -
          Screen.safeBottomArea -
          style.profileBgHeight -
          48,
      child: Consumer2<MyModelHandler, UserModel>(builder: (ctx, m, u, child) {
        
        return ListView.builder(
          itemCount: m.fetchMyList(u.isLogined).length,
          itemBuilder: (BuildContext context, int index) {
            
            return ListTile(
            );
          },
        );
      }),
    );
  }
}
