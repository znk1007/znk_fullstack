import 'package:flb/page/classify/classify.dart';
import 'package:flb/page/home/home.dart';
import 'package:flb/page/my/my.dart';
import 'package:flb/page/shop/shop.dart';
import 'package:flutter/material.dart';

class TabbarItem {
  final String identifier;//唯一标识
  final int index;
  final BottomNavigationBarItem item;
  Widget page;

  TabbarItem({
    this.identifier,
    this.index, 
    this.page,
    @required this.item,
  });
}

class TabbarItems extends ChangeNotifier {

  //分栏项目
  List<TabbarItem> _items = [];

  //分栏项目
  List<TabbarItem> get items => _items;

  //版本1分栏页面
  List<Map<String, Widget>> _pages_v1 = [
    {
      HomePage.id:HomePage()
    }, 
    {
      ClassifyPage.id:ClassifyPage()
    }, 
    {
      ShopPage.id:ShopPage()
    }, 
    {
      MyPage.id:MyPage()
    },
  ];

  //添加分栏项目
  void add(List<TabbarItem> items) {
    if (items.length == 0) {
      _setDefaultItems();
    } else {
      for (var idx = 0; idx < items.length; idx++) {
        TabbarItem item = items[idx];
        int curIdx = _pages_v1.indexWhere((elem) => elem[item.identifier] != null);
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