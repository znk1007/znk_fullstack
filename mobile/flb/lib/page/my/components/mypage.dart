import 'package:flb/page/base/hud.dart';
import 'package:flb/page/my/components/myprofile.dart';
import 'package:flb/util/random/color.dart';
import 'package:flutter/material.dart';

class MyPage extends StatelessWidget {
  static String id = 'my';

  const MyPage({Key key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      color: RandomHandler.randomColor,
      child: Column(
        children: [
          MyProfile(),
        ],
      ),
    );
  }
}
