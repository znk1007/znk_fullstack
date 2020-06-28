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
}

final List<TabbarItem> tabbarItems = <TabbarItem>[
  TabbarItem(
    index: 0,

    item: BottomNavigationBarItem(
      icon: Icon(Icons.home),
      title: Text('首页'),
    ),
  ),
  TabbarItem(
    index: 1,
    item: BottomNavigationBarItem(
      icon: Icon(Icons.sort),
      title: Text('分类'),
    ),
  ),
  TabbarItem(
    index: 2,
    item: BottomNavigationBarItem(
      icon: Icon(Icons.shop),
      title: Text('购物车标题'),
    ),
  ),
  TabbarItem(
    index: 3,
    item: BottomNavigationBarItem(
      icon: Icon(Icons.people),
      title: Text('我的'),
    ),
  ),
];