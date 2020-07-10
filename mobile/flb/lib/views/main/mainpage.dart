import 'package:cached_network_image/cached_network_image.dart';
import 'package:flb/api/api.dart';
import 'package:flb/models/main/main.dart';
import 'package:flb/models/style/style.dart';
import 'package:flb/pkg/banner/banner.dart';
import 'package:flb/pkg/screen/screen.dart';
import 'package:flb/viewmodels/main/collection.dart';
import 'package:flb/views/main/banner.dart';
import 'package:flb/views/main/search.dart';
import 'package:flutter/material.dart';
import 'package:flutter_easyrefresh/easy_refresh.dart';
import 'package:provider/provider.dart';

class ZNKMainPage extends StatelessWidget {
  static const String id = 'home';
  ZNKMainPage({Key key}) : super(key: key);
  //刷新控制
  EasyRefreshController _refreshController = EasyRefreshController();
  //集合
  List<ZNKCollection> _collections = [];

  @override
  Widget build(BuildContext context) {
    double bannerHeight = ZNKScreen.setWidth(195);
    ZNKApi api = Provider.of(context);
    return MultiProvider(providers: [
      Provider(create: (_) => ZNKCollectionViewModel(url: api.collectionUrl)),
    ],child: Consumer<ThemeStyle>(builder: (ctx, style, child) {
      return Stack(
        children: [
          //整体页面
          Container(
            height: ZNKScreen.screenHeight - style.tabbarHeight,
            child: Consumer5<>(builder: null)
          ),
          //搜索加消息数
          ZNKSearchView(),
        ],
      );
    }),);
  }
}

/*
EasyRefresh.custom(
                controller: _refreshController,
                onRefresh: () async {
                  print('on refresh');
                },
                onLoad: () async {
                  print('on load');
                },
                
                slivers: <Widget>[
                  SliverList(
                      delegate: SliverChildBuilderDelegate((ctx, index) {
                    return ZNKBannerView(height: bannerHeight);
                  }, childCount: 1)),
                  // SliverGrid(delegate: null, gridDelegate: null)
                ])
*/