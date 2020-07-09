import 'package:flb/models/style/style.dart';
import 'package:flb/models/main.dart';
import 'package:flb/pkg/banner/banner.dart';
import 'package:flb/viewmodels/main.dart';
import 'package:flb/views/base/base.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class ZNKMainPage extends StatelessWidget {
  static const String id = 'home';
  const ZNKMainPage({Key key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ZNKBaseView<ZNKMainViewModel>(
      model: ZNKMainViewModel(api: Provider.of(context)),

    );
  }
}

class MainPage extends StatefulWidget {
  static String id = 'home';

  MainPage({Key key}) : super(key: key);

  @override
  _MainPageState createState() => _MainPageState();
}

class _MainPageState extends State<MainPage> {

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
  void didUpdateWidget(MainPage oldWidget) {
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
