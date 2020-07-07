import 'package:flb/model/style/style.dart';
import 'package:flb/pkg/banner/banner.dart';
import 'package:flb/pkg/screen/screen.dart';
import 'package:flb/pkg/search/search.dart';
import 'package:flb/util/random/color.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class HomePage extends StatefulWidget {
  static String id = 'home';

  HomePage({Key key}) : super(key: key);

  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  @override
  Widget build(BuildContext context) {
    return Consumer<ThemeStyle>(builder: (ctx, style, child) {
      bool test = true;
      return test
          ? _bannerView()
          : Container(
              child: SearchView(
                  style: SearchStyle(
                      margin: EdgeInsets.only(
                          left: 50, top: Screen.safeTopArea + 10),
                      width: 300,
                      cornerRadius: 45 / 2.0)));
    });
  }

  Widget _bannerView() {
    List<Widget> models = [
      Container(child: Text('测试一'),color: RandomHandler.randomColor),
      Container(child: Text('测试二'),color: RandomHandler.randomColor),
      Container(child: Text('测试三'),color: RandomHandler.randomColor),
      Container(child: Text('测试四'),color: RandomHandler.randomColor),
      Container(child: Text('测试五'),color: RandomHandler.randomColor),
      Container(child: Text('测试六'),color: RandomHandler.randomColor),
      Container(child: Text('测试日'),color: RandomHandler.randomColor),
    ];
    return ZNKBanner(
      banners: models,
      margin: EdgeInsets.only(left: 50, top: 80),
      scrollDirection: Axis.horizontal,
      alignment: Alignment.centerLeft,
    );
  }
}
