import 'dart:ui';

import 'package:cached_network_image/cached_network_image.dart';
import 'package:flb/models/main/combine.dart';
import 'package:flb/models/style/style.dart';
import 'package:flb/pkg/screen/screen.dart';
import 'package:flb/viewmodels/main/combine.dart';
import 'package:flb/views/base/base.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class ZNKCombineView extends StatelessWidget {
  const ZNKCombineView({
    Key key,
    @required this.style,
    @required this.width,
    @required this.height,
    @required this.show,
  }) : super(key: key);
  //主题样式
  final ThemeStyle style;
  //宽度
  final double width;
  //高度
  final double height;
  //是否显示
  final bool show;

  @override
  Widget build(BuildContext context) {
    double containWidth = this.width;
    double topHeight = this.height * (1 / 4.0);
    return ZNKBaseView<ZNKCombineViewModel>(
        model: ZNKCombineViewModel(
          api: Provider.of(context),
        ),
        onReady: (model) => model.fetch(),
        builder: (context, model, child) {
          List<ZNKCombineItem> items = model.combine.items;
          return Container(
            width: containWidth,
            child: Column(
              children: [
                Container(
                  height: topHeight,
                  width: containWidth,
                  padding: EdgeInsets.only(
                    left: ZNKScreen.setWidth(5),
                  ),
                  alignment: Alignment.centerLeft,
                  child: Text(
                    model.combine.title,
                    style: TextStyle(
                      fontFamily: 'PingFangSC-Medium',
                      fontSize: 16,
                      fontWeight: FontWeight.w600,
                    ),
                    overflow: TextOverflow.ellipsis,
                  ),
                ),
                Container(
                  height: this.height - topHeight,
                  width: containWidth - ZNKScreen.setWidth(15),
                  child: ListView.builder(
                    scrollDirection: Axis.horizontal,
                    itemCount: items.length,
                    itemBuilder: (ctx, index) {
                      ZNKCombineItem item = items[index];
                      return Column(
                        children: [
                          Container(
                            margin: EdgeInsets.only(
                                left: index > 0 ? 3 : 0, top: 2),
                            width: (containWidth - 20) / 3.0,
                            height: (this.height - topHeight) / 2.0,
                            padding: EdgeInsets.all(3),
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
                                fontSize: 10,
                                color: style.dartTextColor,
                                fontWeight: FontWeight.w500,
                              ),
                            ),
                          ),
                          Container(
                            width: (containWidth - 20) / 3.0,
                            height: ZNKScreen.setWidth(18),
                            alignment: Alignment.center,
                            margin: EdgeInsets.only(top: 5),
                            decoration: BoxDecoration(
                              gradient: LinearGradient(
                                colors: [Colors.white, Colors.red[100]],
                                begin: Alignment.topCenter,
                                end: Alignment.bottomCenter,
                              ),
                            ),
                            child: Text.rich(
                              TextSpan(
                                children: [
                                  TextSpan(
                                    text: item.coinType,
                                    style: TextStyle(
                                      fontSize: 8,
                                      fontWeight: FontWeight.w500,
                                      color: this.style.redColor,
                                    ),
                                  ),
                                  TextSpan(
                                    text: item.price,
                                    style: TextStyle(
                                      fontSize: 12,
                                      fontWeight: FontWeight.w500,
                                      color: this.style.redColor,
                                    ),
                                  ),
                                ],
                              ),
                            ),
                          ),
                        ],
                      );
                    },
                  ),
                )
              ],
            ),
          );
        });
  }
}
