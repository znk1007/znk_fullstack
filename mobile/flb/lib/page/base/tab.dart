import 'dart:math';

import 'package:flb/page/base/item.dart';
import 'package:flutter/material.dart';
class TabPage extends StatefulWidget {
  final List<TabbarItem> items;
  List<Widget> _pages = [];
  TabPage({Key key, this.items}) : assert(items != null) {
    _pages = items.map((e) => e.page).toList();
  }

  @override
  _TabPageState createState() => _TabPageState();
}

class _TabPageState extends State<TabPage> {
  //_curPageIdx 当前下标
  int _curPageIdx = 0;

  Widget _currentPage() {
    int max = widget._pages.length - 1;
    if (max < 0) {
      return Container();
    }
    int pageIdx = min(_curPageIdx, max);
    print('current page $pageIdx');
    return widget._pages[pageIdx];
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: _currentPage(),
      bottomNavigationBar: BottomNavigationBar(
        items:widget.items.map((e) => e.item).toList(),
        onTap: (value) {
          setState(() {
            _curPageIdx = value;
          });
        },
        currentIndex: _curPageIdx,
        type: BottomNavigationBarType.fixed,
        selectedFontSize: 12,
        selectedItemColor: Colors.red[900],
        unselectedItemColor: Colors.grey[900],
      ),
    );
  }
}