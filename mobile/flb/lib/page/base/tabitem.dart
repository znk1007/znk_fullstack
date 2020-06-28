import 'package:flb/page/home/home.dart';
import 'package:flb/page/my/my.dart';
import 'package:flutter/material.dart';

class TabbarItem extends BottomNavigationBarItem {

  final int index;
  final Widget page;
  final BottomNavigationBarItem item;

  const TabbarItem({
    this.index, 
    @required this.page,
    @required this.item,
  });
}

class TabbarItemHandler {
  //分栏项目
  static List<TabbarItem> _items = [];
  //分栏项目
  static get items => _items;
  //添加分栏项目
  static void add(TabbarItem item) {
    _items.add(item);
  }
  //默认分栏项目集合
  static void setDefaultItems() {
    _items = [
      TabbarItem(
        page: HomePage(), 
        item: BottomNavigationBarItem(
          icon: Icon(Icons.home), 
          title: Text('首页'),
        ),
        index: 0,
      ), 
      TabbarItem(
        page: HomePage(), 
        item: BottomNavigationBarItem(
          icon: Icon(Icons.category), 
          title: Text('分类'),
        ),
        index: 1,
      ),
      TabbarItem(
        page: HomePage(), 
        item: BottomNavigationBarItem(
          icon: Icon(Icons.shopping_cart), 
          title: Text('购物车'),
        ),
        index: 2,
      ),
      TabbarItem(
        page: MyPage(), 
        item: BottomNavigationBarItem(
          icon: Icon(Icons.person), 
          title: Text('我的'),
        ),
        index: 3,
      ),
    ];
  }
}