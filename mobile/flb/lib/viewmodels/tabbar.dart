import 'package:cached_network_image/cached_network_image.dart';
import 'package:flb/api/api.dart';
import 'package:flb/models/tabbar.dart';
import 'package:flb/util/http/core/request.dart';
import 'package:flb/util/http/tab/tab.dart';
import 'package:flb/viewmodels/base.dart';
import 'package:flb/views/classify/classifypage.dart';
import 'package:flb/views/home/homepage.dart';
import 'package:flb/views/my/mypage.dart';
import 'package:flb/views/shop/shoppage.dart';
import 'package:flb/views/tabbar/item.dart';
import 'package:flutter/material.dart';

class ZNKTabbarModel extends ZNKBaseViewModel {
  ZNKTabbarModel({@required ZNKApi api}) : super(api: api);
  //数据
  List<TabbarItem> items = [];

  //获取分栏类目数据
  Future<void> fetch() async {

    ResponseResult result = await RequestHandler.get(this.api.tabbarUrl);
    result.code = -1;
    if (result.statusCode == 200 && result.data != null) {
      String code = result.data['code'];
      result.code = int.parse(code);
      List<Map> body = result.data['body'];
      if (body.length == 0) {
        result.code = -1;
      }
    }
    ResponseResult res = await TabbarItemReq.fetch();
    if (result.statusCode != 0) {
      return;
    }
    List<Map> body = res.data['body'];
    List<TabbarItem> items = [];
    for (var i = 0; i < body.length; i++) {
      Map<String, dynamic> itemMap = body[i];
      TabbarItem item = TabbarItem(
          identifier: itemMap["identifier"] ? "${itemMap['identifier']}" : "$i",
          index: i,
          item: BottomNavigationBarItem(
              icon: CachedNetworkImage(
            imageUrl: itemMap['icon'],
            placeholder: (context, url) => Icon(Icons.pages),
          )));
      items.add(item);
    }
    context.read<TabbarItems>().add(items);
  }

}

class TabbarItems extends ChangeNotifier {
  //分栏项目
  List<TabbarItem> _items = [];

  //分栏项目
  List<TabbarItem> get items => _items;

  //版本1分栏页面
  List<Map<String, Widget>> _pages_v1 = [
    {HomePage.id: HomePage()},
    {ClassifyPage.id: ClassifyPage()},
    {ShopPage.id: ShopPage()},
    {MyPage.id: MyPage()},
  ];

  //添加分栏项目
  void add(List<TabbarItem> items) {
    if (items.length == 0) {
      _setDefaultItems();
    } else {
      for (var idx = 0; idx < items.length; idx++) {
        TabbarItem item = items[idx];
        int curIdx =
            _pages_v1.indexWhere((elem) => elem[item.identifier] != null);
        if (curIdx != -1) {
          Map<String, Widget> pageMap = _pages_v1[curIdx];
          item.page = pageMap[item.identifier];
          _items.add(item);
        }
      }
    }

    notifyListeners();
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
        identifier: HomePage.id,
        page: HomePage(),
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
  }
}