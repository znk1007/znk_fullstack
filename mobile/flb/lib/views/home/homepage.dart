import 'package:flb/models/style/style.dart';
import 'package:flb/models/home.dart';
import 'package:flb/pkg/banner/banner.dart';
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
  void initState() {
    super.initState();
  }

  @override
  void didChangeDependencies() {
    super.didChangeDependencies();
    print('didChangeDependencies');
  }

  @override
  void didUpdateWidget(HomePage oldWidget) {
    super.didUpdateWidget(oldWidget);
    print('didUpdateWidget');
  }

  @override
  Widget build(BuildContext context) {
    return MultiProvider(
      providers: [ChangeNotifierProvider(create: (_) => ZNKBannerModel())],
      child: Consumer<ThemeStyle>(builder: (ctx, style, child) {
        return Container(
          child: Stack(
            children: [
              Container(
                child: Row(
                  children: [
                    Container(
                      child: Consumer<ZNKBannerModel>(
                          builder: (ctx, bannerModel, child) {
                        return bannerModel.recommends.length > 0 ? ZNKBanner(
                          banners: bannerModel.recommends
                              .map((e) => Container(child: Text(e)))
                              .toList(),
                          margin: EdgeInsets.only(left: 50, top: 80),
                          scrollDirection: Axis.horizontal,
                          alignment: Alignment.centerLeft,
                          didSelected: (index) {
                            print('did selected: $index');
                          },
                        ) : Container();
                      }),
                    ),
                  ],
                ),
              ),
            ],
          ),
        );
      }),
    );
  }
}
