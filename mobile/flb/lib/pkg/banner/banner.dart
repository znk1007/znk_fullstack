import 'dart:async';
import 'dart:math';

import 'package:flutter/material.dart';

//选择事件回调
typedef ZNKBannerDidSelected = Function(int index);

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
  //选择事件回调
  final ZNKBannerDidSelected didSelected;
  //动画效果
  final Curve curve;
  //是否显示指示器
  final bool showIndicator;
  //只有一个页码时，是否隐藏指示器
  final bool hideIndicatorWhileSingle;
  //指示器点大小
  final double indicatorDotSize;
  //指示器普通颜色
  final Color indicatorTintColor;
  //指示器轨迹颜色
  final Color indicatorTrackColor;

  _ZNKBannerState state = _ZNKBannerState();

  ZNKBanner(
      {Key key,
      this.scrollDirection = Axis.horizontal,
      this.size = const Size(300, 45),
      this.interval = 2,
      this.margin = EdgeInsets.zero,
      this.alignment = Alignment.centerLeft,
      this.curve = Curves.linear,
      this.didSelected,
      this.showIndicator = true,
      this.hideIndicatorWhileSingle = true,
      this.indicatorDotSize = 8.0,
      this.indicatorTintColor = Colors.lightBlue,
      this.indicatorTrackColor = Colors.blue,
      @required this.banners})
      : assert(banners != null && banners.length > 0),
        super(key: key);

  @override
  _ZNKBannerState createState() => state;
}

class _ZNKBannerState extends State<ZNKBanner>
    with AutomaticKeepAliveClientMixin {
  //当前页
  int _curPage = 0;
  //上一页
  int get _prePage =>
      _curPage - 1 >= 0 ? _curPage - 1 : widget.banners.length - 1;
  //下一页
  int get _nextPage => _curPage + 1 < widget.banners.length ? _curPage + 1 : 0;
  //页码控制器
  PageController _pageController = PageController(initialPage: 1);
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

  void _startTimer() {
    if (widget.banners.length < 2) {
      return;
    }
    _stopTimer();
    _timer = Timer.periodic(Duration(seconds: 2 + widget.interval), (timer) {
      _pageController.animateToPage(2,
          duration: Duration(seconds: widget.interval), curve: widget.curve);
    });
  }

  void _stopTimer() {
    if (_timer != null) {
      _timer.cancel();
      _timer = null;
    }
  }

  @override
  void didChangeDependencies() {
    super.didChangeDependencies();
    _startTimer();
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
    return Stack(
      children: [
        _pageView(),
        _indicator(),
      ],
    );
  }

  Widget _pageView() {
    print('pre page: $_prePage current page: $_curPage next page: $_nextPage');
    return GestureDetector(
        onTap: () {
          _stopTimer();
          if (widget.didSelected != null) {
            widget.didSelected(_curPage);
          }
          Future.delayed(Duration(seconds: 1), () {
            _startTimer();
          });
        },
        child: Container(
          color: Colors.cyan,
          margin: widget.margin,
          height: widget.size.height,
          width: widget.size.width,
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

  Widget _indicator() {
    if (!widget.showIndicator) {
      return Container();
    }
    if (widget.hideIndicatorWhileSingle && widget.banners.length == 1) {
      return Container();
    }
    return _ZNKBannerIndicator(
      margin: EdgeInsets.only(
          top: widget.margin.top +
              widget.size.height -
              widget.indicatorDotSize -
              2.0),
      controller: _pageController,
      current: _curPage,
      itemCount: widget.banners.length,
      dotMaxZoom: 1.0,
      tintColor: widget.indicatorTintColor,
      trackColor: widget.indicatorTrackColor,
      dotSize: widget.indicatorDotSize,
    );
  }

  @override
  bool get wantKeepAlive => true;
}

class _ZNKBannerIndicator extends AnimatedWidget {
  //分页控制器
  final PageController controller;
  //个数
  final int itemCount;
  //普通点颜色
  final Color tintColor;
  //选中点颜色
  final Color trackColor;
  //点大小
  final double dotSize;
  //点间距
  final double dotSpacing;
  //点放大最大范围
  final double dotMaxZoom;
  //当前点
  final int current;
  //边距
  final EdgeInsets margin;
  //选择指示器点回调
  final Function(int index) didSelectedIndex;

  _ZNKBannerIndicator(
      {this.controller,
      this.itemCount,
      this.tintColor = Colors.grey,
      this.trackColor = Colors.blue,
      this.margin = EdgeInsets.zero,
      this.current = 0,
      this.dotSize = 8.0,
      this.dotSpacing = 5.0,
      this.dotMaxZoom = 2,
      this.didSelectedIndex})
      : assert(controller != null),
        super(listenable: controller);

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: this.margin,
      child: Row(
        mainAxisAlignment: MainAxisAlignment.center,
        children: List<Widget>.generate(this.itemCount, _createDot),
      ),
    );
  }

  //创建点
  Widget _createDot(int index) {
    double selectedness =
        Curves.easeOut.transform(max(0.0, 1.0 - (current - index).abs()));
    double zoom = 1.0 + (this.dotMaxZoom - 1.0) * selectedness;
    return Container(
      width: max(this.dotSpacing * (this.itemCount / 3), 10),
      child: Material(
        color: this.current == index ? this.trackColor : this.tintColor,
        type: MaterialType.circle,
        child: Container(
          width: this.dotSize * zoom,
          height: this.dotSize * zoom,
          child: InkWell(
            onTap: () {
              if (didSelectedIndex != null) {
                didSelectedIndex(index);
              }
            },
          ),
        ),
      ),
    );
  }
}
