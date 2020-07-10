import 'package:cached_network_image/cached_network_image.dart';
import 'package:flb/api/api.dart';
import 'package:flb/models/main/main.dart';
import 'package:flb/models/style/style.dart';
import 'package:flb/pkg/banner/banner.dart';
import 'package:flb/pkg/screen/screen.dart';
import 'package:flb/pkg/search/search.dart';
import 'package:flb/viewmodels/main/main.dart';
import 'package:flb/views/base/base.dart';
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
    return ZNKBaseView<ZNKMainViewModel>(
      model: ZNKMainViewModel(api: api),
      onReady: (mainVM) async {
        mainVM.fetchMainLayoutConfig();
        mainVM.fetchRecommends();
        mainVM.fetchMsgNum();
        mainVM.fetchBanner();
        mainVM.fetchNav();
      },
      builder: (context, mainVM, child) {
        return Consumer<ThemeStyle>(builder: (ctx, style, child) {
          return Container(
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
                        _sliderModule(mainVM, bannerHeight),
                        _searchModule(mainVM),
                      ],
                    ),
                  ],
                ),
              ));
        });
      },
    );
  }

  //幻灯片模块
  Widget _sliderModule(ZNKMainViewModel mainVM, double bannerHeight) {
    return ZNKBanner(
      indicatorTrackColor: Color(0xFFD73B1E), //ox后两位表示透明度
      indicatorTintColor: Colors.white,
      size: Size(ZNKScreen.screenWidth, bannerHeight),
      banners: mainVM.banners
          .map((e) =>
              (e.path.startsWith('http://') || e.path.startsWith('https://'))
                  ? CachedNetworkImage(imageUrl: e.path)
                  : Image.asset(e.path,
                      fit: BoxFit.fill,
                      width: ZNKScreen.screenWidth,
                      height: bannerHeight))
          .toList(),
      scrollDirection: Axis.horizontal,
      alignment: Alignment.centerLeft,
      didSelected: (index) {
        print('did selected: $index');
      },
    );
  }

  //搜索框部分
  Widget _searchModule(ZNKMainViewModel mainVM) {
    Size searchSize = Size(ZNKScreen.screenWidth - 70, 31.0);
    double msgSize = 12;
    double msgIconSize = 23;
    //搜索加消息数
    return Row(
      children: [
        (mainVM.showModule(ZNKMainModule.msessage) &&
                mainVM.recommends.length > 0)
            ? ZNKSearch(
                style: ZNKSearchStyle(
                  enabled: false,
                  backgroudColor: Colors.white,
                  cornerRadius: searchSize.height / 2.0,
                  margin: EdgeInsets.only(left: 14, top: ZNKScreen.safeTopArea),
                  width: searchSize.width,
                  height: searchSize.height,
                ),
                child: ZNKBanner(
                  size: Size(searchSize.width, searchSize.height),
                  decoration: BoxDecoration(
                      borderRadius: BorderRadius.all(
                          Radius.circular(searchSize.height / 2.0))),
                  banners: mainVM.recommends
                      .map((e) => Container(
                            child: Text(
                              e,
                              style: TextStyle(color: Color(0xFF999999)),
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
            : Container(),
        mainVM.showModule(ZNKMainModule.msessage)
            ? Stack(
                children: [
                  Container(
                    child: Icon(Icons.message),
                    width: msgIconSize,
                    height: msgIconSize,
                    margin: EdgeInsets.only(
                        left: 16,
                        top: (searchSize.height +
                                ZNKScreen.safeTopArea -
                                msgIconSize / 2.0) /
                            2.0),
                  ),
                  Container(
                    width: msgSize,
                    height: msgSize,
                    margin: EdgeInsets.only(
                        left: msgIconSize + msgSize,
                        top: (searchSize.height) / 2.0),
                    decoration: BoxDecoration(
                      color: Colors.red,
                      borderRadius: BorderRadius.circular(msgSize / 2.0),
                    ),
                    child: Text(mainVM.msgNum,
                        textAlign: TextAlign.center,
                        style: TextStyle(fontSize: 10, color: Colors.white)),
                  ),
                ],
              )
            : Container()
      ],
    );
  }
}

/*

*/
