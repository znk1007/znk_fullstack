import 'package:cached_network_image/cached_network_image.dart';
import 'package:flb/pkg/screen/screen.dart';
import 'package:flb/util/random/random.dart';
import 'package:flutter/material.dart';

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
  const ZNKGrid({
    Key key,
    this.items = const [],
    this.numberOfRow = 2,
    this.scrollDirection = Axis.horizontal,
    this.height = 148,
    this.childAspectRatio = 0.98,
  }) : super(key: key);
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
  //滚动控制器
  ScrollController _controller = ScrollController();

  @override
  Widget build(BuildContext context) {
    return Container(
      height: this.height,
      width: ZNKScreen.screenWidth,
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
                color: RandomHandler.randomColor,
                child: Column(
                  children: [
                    Container(
                      margin: EdgeInsets.only(top: 5),
                      alignment: Alignment.topCenter,
                      width: 35,
                      height: 35,
                      child: (item.path.startsWith('http://') ||
                              item.path.startsWith('https://'))
                          ? CachedNetworkImage(imageUrl: item.path)
                          : Image.asset(item.path),
                    ),
                    Container(
                      margin: EdgeInsets.only(top: 3),
                      child: Text(
                        item.title,
                        style:
                            TextStyle(fontSize: 10, color: Color(0xFF666666)),
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
    );
  }
}
