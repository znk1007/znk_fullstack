import 'package:flb/api/api.dart';
import 'package:flb/models/main/main.dart';
import 'package:flb/models/style/style.dart';
import 'package:flb/pkg/screen/screen.dart';
import 'package:flb/util/random/random.dart';
import 'package:flb/viewmodels/main/main.dart';
import 'package:flb/views/base/base.dart';
import 'package:flb/views/main/banner.dart';
import 'package:flb/views/main/magic.dart';
import 'package:flb/views/main/msgnum.dart';
import 'package:flb/views/main/nav.dart';
import 'package:flb/views/main/search.dart';
import 'package:flb/views/main/seckill.dart';
import 'package:flutter/material.dart';
import 'package:flutter_easyrefresh/easy_refresh.dart';
import 'package:provider/provider.dart';

class ZNKMainPage extends StatelessWidget {
  static const String id = 'home';
  ZNKMainPage({Key key}) : super(key: key);
  //刷新控制
  EasyRefreshController _refreshController = EasyRefreshController();

  @override
  Widget build(BuildContext context) {
    double bannerHeight = ZNKScreen.setWidth(195);
    double navHeight = bannerHeight * (2 / 3.0);
    double magicHeight = 35;
    double seckillHeight = navHeight;
    ZNKApi api = Provider.of(context);
    return ZNKBaseView<ZNKMainViewModel>(
      model: ZNKMainViewModel(api: api),
      onReady: (mainVM) async {
        mainVM.fetchMainLayoutConfig();
      },
      builder: (context, mainVM, child) {
        final Size searchSize = Size(ZNKScreen.screenWidth - 70, 31.0);
        return Consumer<ThemeStyle>(builder: (ctx, style, child) {
          return Container(
              color: style.backgroundColor,
              height: ZNKScreen.screenHeight,
              child: EasyRefresh(
                controller: _refreshController,
                onRefresh: () async {
                  print('on refresh');
                },
                onLoad: () async {
                  print('on load');
                },
                child: Column(
                  children: [
                    Stack(
                      children: [
                        ZNKBannerView(
                          show: mainVM.showModule(ZNKMainModule.slide),
                          bannerHeight: bannerHeight,
                        ),
                        Row(
                          children: [
                            ZNKSearchView(
                              show: mainVM.showModule(ZNKMainModule.search),
                              searchSize: searchSize,
                            ),
                            ZNKMsgNumView(
                              style: style,
                              marginTop: searchSize.height / 2.0,
                              show: mainVM.showModule(ZNKMainModule.msessage),
                            ),
                          ],
                        ),
                      ],
                    ),
                    ZNKNavView(
                      show: mainVM.showModule(ZNKMainModule.nav),
                      navHeight: navHeight,
                    ),
                    ZNKMagicView(
                        show: mainVM.showModule(ZNKMainModule.magic),
                        magicHeight: magicHeight),
                    Container(
                        margin: EdgeInsets.only(top: 10),
                        child: Row(
                          children: [
                            ZNKSeckillView(
                                show: mainVM.showModule(ZNKMainModule.seckill),
                                style: style,
                                seckillHeight: seckillHeight),
                            _combineModule(mainVM, style, seckillHeight),
                          ],
                        )),
                  ],
                ),
              ));
        });
      },
    );
  }

  Widget _combineModule(
      ZNKMainViewModel mainVM, ThemeStyle style, double combineHeight) {
    return Container(
      width: ZNKScreen.screenWidth / 2.0,
      height: combineHeight,
      color: RandomHandler.randomColor,
    );
  }
}
