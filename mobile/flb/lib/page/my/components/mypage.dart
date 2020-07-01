import 'package:flb/page/my/components/myprofile.dart';
import 'package:flb/util/random/color.dart';
import 'package:flb/util/screen/screen.dart';
import 'package:flutter/material.dart';

class MyPage extends StatelessWidget {
  static String id = 'my';

  const MyPage({Key key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    double topHeight = Screen.setHeight(180);
    return Container(
      color: RandomHandler.randomColor,
      child: Column(
        children: [
          MyProfile(
            profileHeight: topHeight,
          ),
        ],
      ),
    );
  }
}
