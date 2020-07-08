import 'dart:math';

import 'package:flb/views/base/hud.dart';
import 'package:flb/views/tabbar/item.dart';
import 'package:flb/pkg/screen/screen.dart';
import 'package:flutter/material.dart';
import 'package:loading_overlay/loading_overlay.dart';
import 'package:provider/provider.dart';

class ZNKTabbar extends StatefulWidget {
  //页面集合
  List<Widget> _pages = [];
  ZNKTabbar({Key key, List<TabbarItem> items}) : assert(items.length > 2) {
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
    //设置屏幕
    Screen.setContext(context);
    return Scaffold(
      body: LoadingOverlay(
        child: SingleChildScrollView(
          physics: NeverScrollableScrollPhysics(),
          child: Container(
            padding: const EdgeInsets.all(0),
            child: _currentPage(),
          ),
        ),
        isLoading: _isLoading,
        // demo of some additional parameters
        opacity: 0,
        progressIndicator: CircularProgressIndicator(),
      ), //_currentPage(),
      bottomNavigationBar: Consumer<TabbarItems>(
          builder: (ctx, t, w) => BottomNavigationBar(
                items: t.items.map((e) => e.item).toList(),
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
              )),
    );
  }
}
