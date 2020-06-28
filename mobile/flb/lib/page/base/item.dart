import 'package:flb/page/home/home.dart';
import 'package:flb/page/my/my.dart';
import 'package:flutter/material.dart';

class TabbarItem extends BottomNavigationBarItem {
  final String identifier;//唯一标识
  final int index;
  final BottomNavigationBarItem item;

  const TabbarItem({
    this.identifier,
    this.index, 
    @required this.item,
  });
}

class TabbarItems extends ChangeNotifier {
  //分栏项目
  List<TabbarItem> _items = [];
  //分栏项目
  List<TabbarItem> get items => _items;
  //添加分栏项目
  void add(List<TabbarItem> items) {
    if (items.length == 0) {
      _setDefaultItems();
    } else {
      _items.addAll(items);
    }
    
    notifyListeners();
  }
  //默认分栏项目集合
  void _setDefaultItems() {
    _items = [
      TabbarItem(
        item: BottomNavigationBarItem(
          icon: Icon(Icons.home), 
          title: Text('首页'),
        ),
        index: 0,
      ), 
      TabbarItem(
        item: BottomNavigationBarItem(
          icon: Icon(Icons.category), 
          title: Text('分类'),
        ),
        index: 1,
      ),
      TabbarItem(
        item: BottomNavigationBarItem(
          icon: Icon(Icons.shopping_cart), 
          title: Text('购物车'),
        ),
        index: 2,
      ),
      TabbarItem(
        item: BottomNavigationBarItem(
          icon: Icon(Icons.person), 
          title: Text('我的'),
        ),
        index: 3,
      ),
    ];
  }
}