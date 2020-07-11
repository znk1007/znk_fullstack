import 'dart:math';

import 'package:cached_network_image/cached_network_image.dart';
import 'package:flb/pkg/screen/screen.dart';
import 'package:flutter/material.dart';

//滚动偏移回调
typedef ZNKSlideOffsetWrapHandler = void Function(void Function(double offset));
//滑块偏移回调
typedef ZNKSlideOffsetHandler = void Function(double offset);

class ZNKGridItem {
  //标题
  final String title;
  //路径
  final String path;
  //唯一标识
  final String identifier;

  ZNKGridItem({this.identifier, this.title, this.path});
}

class ZNKGrid extends StatelessWidget {
  ZNKGrid({
    Key key,
    this.items = const [],
    this.numberOfRow = 2,
    this.scrollDirection = Axis.horizontal,
    this.height = 135,
    this.childAspectRatio = 0.9,
    this.iconSize = const Size(35, 35),
  }) {
    _controller.addListener(() {
      if (_slideHandler != null) {
        double max = _controller.position.maxScrollExtent;
        _slideHandler(_controller.offset / max);
      }
    });
  }
  //数据源
  final List<ZNKGridItem> items;
  //行数
  final int numberOfRow;
  //滚动方向
  final Axis scrollDirection;
  //高度
  final double height;
  //宽高比例
  final double childAspectRatio;
  //icon 大小
  final Size iconSize;
  //滚动控制器
  ScrollController _controller = ScrollController();
  //滑动偏移回调
  ZNKSlideOffsetHandler _slideHandler;

  @override
  Widget build(BuildContext context) {
    double slideHeight = 3;
    return Container(
      height: this.height + slideHeight,
      width: ZNKScreen.screenWidth,
      child: Column(
        children: [
          Container(
            width: ZNKScreen.screenWidth,
            height: this.height,
            child: GridView.builder(
                controller: _controller,
                scrollDirection: Axis.horizontal,
                itemCount: this.items.length,
                gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
                    crossAxisCount: this.numberOfRow,
                    childAspectRatio: this.childAspectRatio),
                itemBuilder: (ctx, index) {
                  ZNKGridItem item = this.items[index];
                  return GestureDetector(
                    child: Container(
                      child: Column(
                        children: [
                          Container(
                            margin: EdgeInsets.only(top: 5),
                            alignment: Alignment.topCenter,
                            width: this.iconSize.width,
                            height: this.iconSize.height,
                            child: (item.path.startsWith('http://') ||
                                    item.path.startsWith('https://'))
                                ? CachedNetworkImage(imageUrl: item.path)
                                : Image.asset(item.path),
                          ),
                          Container(
                            margin: EdgeInsets.only(top: 5),
                            child: Text(
                              item.title,
                              style: TextStyle(
                                  fontSize: 10, color: Color(0xFF666666)),
                            ),
                          ),
                        ],
                      ),
                    ),
                    onTap: () {
                      print('select at index: $index');
                    },
                  );
                }),
          ),
          ZNKSlideBar(handler: (ZNKSlideOffsetHandler handler) {
            _slideHandler = handler;
          }),
        ],
      ),
    );
  }
}

class ZNKSlideBar extends StatefulWidget {
  ZNKSlideBar(
      {Key key,
      this.margin = EdgeInsets.zero,
      this.width = 44,
      this.height = 3,
      this.sliderWidth = 22,
      this.tintColor = Colors.grey,
      this.trackColor = Colors.red,
      @required this.handler})
      : super(key: key);
  //宽度
  final double width;
  //高度
  final double height;
  //滑块宽度
  final double sliderWidth;
  //普通颜色
  final Color tintColor;
  //滑块颜色
  final Color trackColor;
  //滚动回调
  final ZNKSlideOffsetWrapHandler handler;
  //边距
  final EdgeInsets margin;

  _ZNKSlieBarState state = _ZNKSlieBarState();

  @override
  _ZNKSlieBarState createState() => state;
}

class _ZNKSlieBarState extends State<ZNKSlideBar> {
  @override
  void initState() {
    super.initState();
    ZNKSlideOffsetHandler slideHandler = (double offset) {
      setState(() {
        _offset = offset;
      });
    };
    if (widget.handler != null) {
      widget.handler(slideHandler);
    }
  }

  double _offset = 0.0;

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: widget.margin,
      width: widget.width,
      height: widget.height,
      decoration: BoxDecoration(
          color: widget.tintColor,
          borderRadius: BorderRadius.circular(widget.height / 2.0)),
      child: Container(
        // width: widget.sliderWidth,
        height: widget.height,
        margin: EdgeInsets.only(
            left: max(
                min(_offset * (widget.width - widget.sliderWidth),
                    widget.width - widget.sliderWidth),
                0),
            right: widget.width -
                min(
                    _offset * (widget.width - widget.sliderWidth) +
                        widget.sliderWidth,
                    widget.width)),
        decoration: BoxDecoration(
          color: widget.trackColor,
          borderRadius: BorderRadius.circular(widget.height / 2.0),
        ),
      ),
    );
  }
}
