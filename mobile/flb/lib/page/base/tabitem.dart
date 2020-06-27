import 'package:flutter/material.dart';

const List<BottomNavigationBarItem> tabbarItems = <BottomNavigationBarItem>[
  BottomNavigationBarItem(
      icon: Text('首页'),
      activeIcon: Text('选中'),
      title: Text('首页标题'),
  ),
  BottomNavigationBarItem(
      icon: Text('分类'),
      activeIcon: Text('选中'),
      title: Text('分类标题'),
  ),
  BottomNavigationBarItem(
      icon: Text('购物车'),
      activeIcon: Text('选中'),
      title: Text('购物车标题'),
  ),
  BottomNavigationBarItem(
      icon: Text('我的'),
      activeIcon: Text('选中'),
      title: Text('我的标题'),
  )

];