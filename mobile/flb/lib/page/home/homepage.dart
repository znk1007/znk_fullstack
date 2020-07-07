import 'package:flb/model/style/style.dart';
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
      return Container(
          child: SearchView(
              style: SearchStyle(
                  margin:
                      EdgeInsets.only(left: 50, top: Screen.safeTopArea + 10),
                  size: Size(300, 45),
                  cornerRadius: 45 / 2.0)));
    });
  }
}
