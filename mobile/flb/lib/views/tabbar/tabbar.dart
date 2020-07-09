import 'dart:math';

import 'package:flb/models/tabbar.dart';
import 'package:flb/views/base/hud.dart';
import 'package:flutter/material.dart';
import 'package:loading_overlay/loading_overlay.dart';

class ZNKTabbar extends StatefulWidget {
  //页面集合
  List<Widget> _pages = [];
  //数据源集合
  final List<TabbarItem> items;
  ZNKTabbar({Key key, this.items}) : assert(items.length > 2) {
    _pages = items.map((e) => e.page).toList();
    //加载框
    Hud().wrap(this);
  }
  //状态初始化
  _ZNKTabbarState state = _ZNKTabbarState();

  @override
  _ZNKTabbarState createState() => state;
}

class _ZNKTabbarState extends State<ZNKTabbar> {
  //_curPageIdx 当前下标
  int _curPageIdx = 0;
  bool _isLoading = false;
  /* 显示加载框 */
  void showLoading() {
    setState(() {
      _isLoading = true;
    });
  }

  /* 隐藏加载框 */
  void hideLoading() {
    setState(() {
      _isLoading = false;
    });
  }

  Widget _currentPage() {
    int max = widget._pages.length - 1;
    if (max < 0) {
      return Container();
    }
    int pageIdx = min(_curPageIdx, max);
    return widget._pages[pageIdx];
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: LoadingOverlay(
        child: SingleChildScrollView(
          physics: NeverScrollableScrollPhysics(),
          child: Container(
            padding: EdgeInsets.all(0),
            child: _currentPage(),
          ),
        ),
        isLoading: _isLoading,
        // demo of some additional parameters
        opacity: 0,
        progressIndicator: CircularProgressIndicator(),
      ), //_currentPage(),
      bottomNavigationBar: BottomNavigationBar(
        items: widget.items.map((e) => e.item).toList(),
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
