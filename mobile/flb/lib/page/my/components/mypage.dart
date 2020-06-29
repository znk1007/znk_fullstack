import 'package:flb/page/base/hud.dart';
import 'package:flb/util/random/color.dart';
import 'package:flutter/material.dart';

class MyPage extends StatelessWidget {

  static String id = 'my';

  MyPage({Key key}) {
    print('my page init');
    Future.delayed(Duration(seconds: 5), () {
      print('my page delay');
      Hud.shared.hide();
    });
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      color: RandomHandler.randomColor,
      child: Column(
        children: [

        ], 
      ),
      
    );
  }
  
}