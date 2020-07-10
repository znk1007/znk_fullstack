import 'package:cached_network_image/cached_network_image.dart';
import 'package:flb/api/api.dart';
import 'package:flb/models/tabbar.dart';
import 'package:flb/state/home.dart';
import 'package:flb/util/http/core/request.dart';
import 'package:flb/viewmodels/base.dart';
import 'package:flb/views/classify/classifypage.dart';
import 'package:flb/views/main/mainpage.dart';
import 'package:flb/views/my/mypage.dart';
import 'package:flb/views/shop/shoppage.dart';
import 'package:flutter/material.dart';

class ZNKHomeModel extends ZNKBaseViewModel {
  ZNKHomeModel({@required ZNKApi api}) : super(api: api);
  //主页加载状态
  ZNKHomeLoadState _state = ZNKHomeLoadState.launching;
  ZNKHomeLoadState get state => _state;

  //数据
  List<TabbarItem> _items = [];
  List<TabbarItem> get items => _items;
  //版本1分栏页面
  List<Map<String, Widget>> _pages_v1 = [
    {ZNKMainPage.id: ZNKMainPage()},
    {ClassifyPage.id: ClassifyPage()},
    {ShopPage.id: ShopPage()},
    {MyPage.id: MyPage()},
  ];

  void setState(ZNKHomeLoadState state) {
    _state = state;
    notifyListeners();
  }

  //获取分栏类目数据
  Future<void> fetch() async {
    if (this.api.tabbarUrl.length <= 0) {
      _setDefaultItems();
      return;
    }
    ResponseResult result = await RequestHandler.get(this.api.tabbarUrl);
    result.code = -1;
    List<Map> data = [];
    if (result.statusCode == 200 && result.data != null) {
      String code = result.data['code'];
      result.code = int.parse(code);
      data = result.data['data'];
      if (data.length == 0) {
        result.code = -1;
      }
    }
    if (result.statusCode != 0) {
      _setDefaultItems();
      return;
    }

    List<TabbarItem> items = [];
    for (var i = 0; i < data.length; i++) {
      Map<String, dynamic> itemMap = data[i];
      TabbarItem item = TabbarItem(
          identifier: itemMap['identifier'] ? '${itemMap["identifier"]}' : "$i",
          index: i,
          item: BottomNavigationBarItem(
              icon: CachedNetworkImage(
            imageUrl: itemMap['icon'],
            placeholder: (context, url) => Icon(Icons.pages),
          )));
      int curIdx =
          _pages_v1.indexWhere((elem) => (elem[item.identifier] != null));
      if (curIdx != -1) {
        items.add(item);
      }
    }
    setState(ZNKHomeLoadState.running);
  }

  //默认分栏项目集合
  void _setDefaultItems() {
    _items = [
      TabbarItem(
        item: BottomNavigationBarItem(
          icon: Icon(
            Icons.home,
          ),
          activeIcon: Icon(
            Icons.home,
          ),
          title: Text('首页'),
        ),
        index: 0,
        identifier: ZNKMainPage.id,
        page: ZNKMainPage(),
      ),
      TabbarItem(
        item: BottomNavigationBarItem(
          icon: Icon(
            Icons.category,
          ),
          activeIcon: Icon(
            Icons.category,
          ),
          title: Text('分类'),
        ),
        index: 1,
        identifier: ClassifyPage.id,
        page: ClassifyPage(),
      ),
      TabbarItem(
        item: BottomNavigationBarItem(
          icon: Icon(
            Icons.shopping_cart,
          ),
          activeIcon: Icon(
            Icons.shopping_cart,
          ),
          title: Text('购物车'),
        ),
        index: 2,
        identifier: ShopPage.id,
        page: ShopPage(),
      ),
      TabbarItem(
        item: BottomNavigationBarItem(
          icon: Icon(
            Icons.person,
          ),
          activeIcon: Icon(
            Icons.person,
          ),
          title: Text('我的'),
        ),
        index: 3,
        identifier: MyPage.id,
        page: MyPage(),
      ),
    ];
    setState(ZNKHomeLoadState.running);
  }
}
