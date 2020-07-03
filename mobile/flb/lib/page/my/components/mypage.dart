import 'package:flb/model/user/user.dart';
import 'package:flb/page/my/components/myprofile.dart';
import 'package:flb/page/my/model/my.dart';
import 'package:flb/util/random/color.dart';
import 'package:flb/util/screen/screen.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class MyPage extends StatelessWidget {
  static String id = 'my';

  const MyPage({Key key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    double topHeight = Screen.setHeight(180);
    return MultiProvider(
      providers: [
        ChangeNotifierProvider(create: (ctx) => MyModelHandler()),
      ],
      child: Container(
        color: RandomHandler.randomColor,
        child: Column(
          children: [
            MyProfile(
              profileHeight: topHeight,
            ),
          ],
        ),
      ),
    );
  }
}
