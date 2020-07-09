import 'package:flutter/material.dart';

class TabbarItem {
  final String identifier; //唯一标识
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
