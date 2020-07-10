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
    this.height = 130,
    this.childAspectRatio = 0.8,
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

  @override
  Widget build(BuildContext context) {
    return Container(
      height: this.height,
      width: ZNKScreen.screenWidth,
      child: GridView.builder(
          shrinkWrap: true,
          scrollDirection: Axis.horizontal,
          itemCount: this.items.length,
          gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
              crossAxisCount: this.numberOfRow,
              childAspectRatio: this.childAspectRatio),
          itemBuilder: (ctx, index) {
            return GestureDetector(
              child: Container(
                  color: RandomHandler.randomColor,
                  width: ZNKScreen.screenWidth / 5.0,
                  child: Text('测试')),
              onTap: () {
                print('select at index: $index');
              },
            );
          }),
    );
  }
}
