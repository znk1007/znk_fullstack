import 'package:flb/model/style/mystyle.dart';
import 'package:flb/util/screen/screen.dart';
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
      child: ListView.builder(
        itemCount: 1,
        itemBuilder: (BuildContext context, int index) {
          return ListTile(
            
          );
        },
      ),
    );
  }
}
