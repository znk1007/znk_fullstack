import 'dart:async';

import 'package:flb/util/random/color.dart';
import 'package:flutter/material.dart';

class ZNKBanner extends StatefulWidget {
  //广告数据源
  final List<Widget> banners;
  //滚动方向
  final Axis scrollDirection;
  //大小
  final Size size;
  //动画间隔
  final int interval;
  //边距
  final EdgeInsets margin;
  //停靠
  final Alignment alignment;

  _ZNKBannerState state = _ZNKBannerState();

  ZNKBanner(
      {Key key,
      this.scrollDirection = Axis.horizontal,
      this.size = const Size(300, 45),
      this.interval = 2,
      this.margin = EdgeInsets.zero,
      this.alignment = Alignment.centerLeft,
      @required this.banners})
      : assert(banners != null && banners.length > 0),
        super(key: key);

  @override
  _ZNKBannerState createState() => state;
}

class _ZNKBannerState extends State<ZNKBanner> {
  //当前页
  int _curPage = 0;
  //上一页
  int get _prePage =>
      _curPage - 1 >= 0 ? _curPage - 1 : widget.banners.length - 1;
  //下一页
  int get _nextPage => _curPage + 1 < widget.banners.length ? _curPage + 1 : 0;

  //页码控制器
  PageController _pageController =
      PageController(initialPage: 0, keepPage: true, viewportFraction: 1.0);
  //定时器
  Timer _timer;
  //监听变化
  void _listenChange() {
    _pageController.addListener(() {
      if (_pageController.offset <= 0) {
        _pageController.jumpToPage(1);
        _changePage(0);
      }
      double nextLimit = (widget.scrollDirection == Axis.horizontal)
          ? widget.size.width
          : widget.size.height;
      if (_pageController.offset >= 2 * nextLimit) {
        _pageController.jumpToPage(1);
        _changePage(2);
      }
    });
  }

  _ZNKBannerState({Key key}) {
    _listenChange();
  }

  //改变页码
  void _changePage(int page) {
    setState(() {
      if (page == 2) {
        _curPage++;
        if (_curPage == widget.banners.length) {
          _curPage = 0;
        }
      }
      if (page == 0) {
        _curPage--;
        if (_curPage == -1) {
          _curPage = widget.banners.length - 1;
        }
      }
    });
  }

  @override
  void didChangeDependencies() {
    super.didChangeDependencies();
    _timer = Timer.periodic(Duration(seconds: widget.interval), (timer) {
      _pageController.animateToPage(2,
          duration: Duration(seconds: widget.interval),
          curve: Curves.fastOutSlowIn);
    });
  }

  @override
  void dispose() {
    if (_timer != null) {
      _timer.cancel();
    }
    if (_pageController != null) {
      _pageController.dispose();
    }
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () {
        print('current index: $_curPage');
      },
        child: Container(
      margin: widget.margin,
      height: widget.size.height,
      width: widget.size.width,
      color: RandomHandler.randomColor,
      alignment: Alignment.center,
      child: PageView(
        controller: _pageController,
        scrollDirection: widget.scrollDirection,
        children: [
          Container(
              child: widget.banners[_prePage],
              alignment: widget.alignment,
              padding: EdgeInsets.zero),
          Container(
              child: widget.banners[_curPage],
              alignment: widget.alignment,
              padding: EdgeInsets.zero),
          Container(
              child: widget.banners[_nextPage],
              alignment: widget.alignment,
              padding: EdgeInsets.zero),
        ],
      ),
    ));
  }
}
