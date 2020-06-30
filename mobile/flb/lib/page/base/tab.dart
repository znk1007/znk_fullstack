import 'dart:math';

import 'package:flb/page/base/item.dart';
import 'package:flb/util/screen/screen.dart';
import 'package:flutter/material.dart';
import 'package:loading_overlay/loading_overlay.dart';

class TabPage extends StatefulWidget {
  final List<TabbarItem> items;
  List<Widget> _pages = [];
  TabPage({Key key, this.items}) : assert(items != null) {
    _pages = items.map((e) => e.page).toList();
  }
  //状态初始化
  _TabPageState state = _TabPageState();

  @override
  _TabPageState createState() => state;
}

class _TabPageState extends State<TabPage> {
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
    //设置屏幕
    Screen.setContext(context);
    return Scaffold(
      body: LoadingOverlay(
        child: SingleChildScrollView(
          child: Container(
            color: Colors.red,
            padding: const EdgeInsets.all(0),
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
