import 'package:cached_network_image/cached_network_image.dart';
import 'package:flb/models/style/style.dart';
import 'package:flb/pkg/banner/banner.dart';
import 'package:flb/pkg/screen/screen.dart';
import 'package:flb/pkg/search/search.dart';
import 'package:flb/viewmodels/main/banner.dart';
import 'package:flb/viewmodels/main/msgnum.dart';
import 'package:flb/viewmodels/main/recommend.dart';
import 'package:flb/views/base/base.dart';
import 'package:flutter/material.dart';
import 'package:flutter_easyrefresh/easy_refresh.dart';
import 'package:provider/provider.dart';

class ZNKMainPage extends StatelessWidget {
  static const String id = 'home';
  ZNKMainPage({Key key}) : super(key: key);
  //刷新控制
  ScrollController _refreshController = ScrollController();

  @override
  Widget build(BuildContext context) {
    Size searchSize = Size(ZNKScreen.screenWidth - 80, 40.0);
    double bannerHeight = ZNKScreen.setWidth(195);
    return Consumer<ThemeStyle>(builder: (ctx, style, child) {
      return Stack(
        children: [
          //整体页面
          Container(
            height: ZNKScreen.screenHeight - style.tabbarHeight,
            child: EasyRefresh.custom(slivers: <Widget>[
              SliverList(
                  delegate: SliverChildBuilderDelegate((ctx, index) {
                return ZNKBaseView<ZNKBannerViewModel>(
                  model: ZNKBannerViewModel(api: Provider.of(context)),
                  onReady: (model) async {
                    model.fetch();
                  },
                  builder: (context, model, child) {
                    return ZNKBanner(
                      indicatorTrackColor: Color(0xFFD73B1E), //ox后两位表示透明度
                      indicatorTintColor: Colors.white,
                      size: Size(ZNKScreen.screenWidth, bannerHeight),
                      banners: model.banners
                          .map((e) => (e.path.startsWith('http://') ||
                                  e.path.startsWith('https://'))
                              ? CachedNetworkImage(imageUrl: e.path)
                              : Image.asset(e.path,
                                  width: ZNKScreen.scaleWidth,
                                  height: bannerHeight))
                          .toList(),
                      scrollDirection: Axis.horizontal,
                      alignment: Alignment.centerLeft,
                      didSelected: (index) {
                        print('did selected: $index');
                      },
                    );
                  },
                );
              }, childCount: 1)),
            ]),
          ),
          //搜索加消息数
          Row(
            children: [
              ZNKBaseView<ZNKMainRecommand>(
                model: ZNKMainRecommand(api: Provider.of(context)),
                onReady: (model) async {
                  model.fetch();
                },
                builder: (context, model, child) {
                  return model.recommends.length > 0
                      ? ZNKSearchView(
                          style: ZNKSearchStyle(
                            enabled: false,
                            backgroudColor: Colors.white,
                            cornerRadius: searchSize.height / 2.0,
                            margin: EdgeInsets.only(
                                left: 14, top: ZNKScreen.safeTopArea),
                            width: searchSize.width,
                            height: searchSize.height,
                          ),
                          child: ZNKBanner(
                            size: Size(searchSize.width, searchSize.height),
                            decoration: BoxDecoration(
                                borderRadius: BorderRadius.all(
                                    Radius.circular(searchSize.height / 2.0))),
                            banners: model.recommends
                                .map((e) => Container(
                                      child: Text(
                                        e,
                                        style:
                                            TextStyle(color: Color(0xFF999999)),
                                      ),
                                      margin: EdgeInsets.only(left: 40),
                                    ))
                                .toList(),
                            showIndicator: false,
                            scrollDirection: Axis.vertical,
                            alignment: Alignment.centerLeft,
                            didSelected: (index) {
                              print('did selected: $index');
                            },
                          ),
                        )
                      : Container();
                },
              ),
              ZNKBaseView<ZNKMsgViewModel>(
                model: ZNKMsgViewModel(api: Provider.of(context)),
                onReady: (model) async {
                  model.fetch();
                },
                builder: (context, model, child) {
                  double msgSize = 12;
                  double msgIconSize = 23;
                  return Stack(
                    children: [
                      Container(
                        child: Icon(Icons.message),
                        width: msgIconSize,
                        height: msgIconSize,
                        margin: EdgeInsets.only(
                            left: 16,
                            top: (searchSize.height +
                                    ZNKScreen.safeTopArea -
                                    msgIconSize) /
                                2.0),
                      ),
                      Container(
                        width: msgSize,
                        height: msgSize,
                        margin: EdgeInsets.only(
                            left: msgIconSize + msgSize,
                            top: (searchSize.height - msgSize) / 2.0),
                        decoration: BoxDecoration(
                          color: Colors.red,
                          borderRadius: BorderRadius.circular(msgSize / 2.0),
                        ),
                        child: Text(model.msgNum,
                            textAlign: TextAlign.center,
                            style:
                                TextStyle(fontSize: 10, color: Colors.white)),
                      ),
                    ],
                  );
                },
              )
            ],
          ),
        ],
      );
    });
  }
}
