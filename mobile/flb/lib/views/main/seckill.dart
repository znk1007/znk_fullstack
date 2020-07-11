import 'package:cached_network_image/cached_network_image.dart';
import 'package:flb/models/main/seckill.dart';
import 'package:flb/models/style/style.dart';
import 'package:flb/pkg/screen/screen.dart';
import 'package:flb/util/random/random.dart';
import 'package:flb/viewmodels/main/seckill.dart';
import 'package:flb/views/base/base.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class ZNKSeckillView extends StatelessWidget {
  final bool show;
  final ThemeStyle style;
  final seckillHeight;
  ZNKSeckillView({
    Key key,
    @required this.show,
    @required this.style,
    @required this.seckillHeight,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    double containWidth = ZNKScreen.screenWidth / 2.0;
    double topHeight = this.seckillHeight * (1 / 4.0);
    return ZNKBaseView<ZNKSeckillViewModel>(
      model: ZNKSeckillViewModel(
        api: Provider.of(context),
      ),
      onReady: (model) async => model.fetch(),
      builder: (context, seckillModel, child) => (this.show &&
              seckillModel.seckill != null &&
              seckillModel.seckill.items.length > 0)
          ? Container(
              child: Column(
                children: [
                  Container(
                    height: topHeight,
                    width: containWidth,
                    child: Row(
                      children: [
                        Container(
                          width: containWidth / 2.0,
                          padding: EdgeInsets.only(left: 10),
                          child: Text(
                            seckillModel.seckill.title,
                            style: TextStyle(
                              color: Color(0xFFD81E06),
                              fontFamily: 'PingFangSC-Medium',
                              fontSize: 16,
                              fontWeight: FontWeight.w500,
                            ),
                          ),
                        ),
                        Container(
                          width: containWidth / 2.0,
                          child: ZNKBaseView<ZNKSeckillCountDownViewModel>(
                            model: ZNKSeckillCountDownViewModel(
                              api: Provider.of(context),
                            ),
                            onReady: (countdown) => countdown
                                .fetch(int.parse(seckillModel.seckill.time)),
                            builder: (context, countdown, child) => Row(
                              children: [
                                Container(
                                  margin: EdgeInsets.only(
                                    left: ZNKScreen.setWidth(30),
                                  ),
                                  width: 16,
                                  height: 16,
                                  alignment: Alignment.center,
                                  decoration: BoxDecoration(
                                    color: this.style.redColor,
                                    borderRadius: BorderRadius.circular(3),
                                  ),
                                  child: Text(
                                    countdown.hour,
                                    style: TextStyle(
                                      color: Colors.white,
                                      fontSize: 11,
                                    ),
                                  ),
                                ),
                                Container(
                                  child: Text(
                                    ':',
                                    style: TextStyle(
                                      color: this.style.redColor,
                                    ),
                                  ),
                                ),
                                Container(
                                  width: 16,
                                  height: 16,
                                  alignment: Alignment.center,
                                  decoration: BoxDecoration(
                                    color: this.style.redColor,
                                    borderRadius: BorderRadius.circular(3),
                                  ),
                                  child: Text(
                                    countdown.minute,
                                    style: TextStyle(
                                      color: Colors.white,
                                      fontSize: 11,
                                    ),
                                  ),
                                ),
                                Container(
                                  child: Text(
                                    ':',
                                    style:
                                        TextStyle(color: this.style.redColor),
                                  ),
                                ),
                                Container(
                                  width: 16,
                                  height: 16,
                                  alignment: Alignment.center,
                                  decoration: BoxDecoration(
                                    color: this.style.redColor,
                                    borderRadius: BorderRadius.circular(3),
                                  ),
                                  child: Text(
                                    countdown.second,
                                    style: TextStyle(
                                      color: Colors.white,
                                      fontSize: 11,
                                    ),
                                  ),
                                ),
                              ],
                            ),
                          ),
                        ),
                      ],
                    ),
                  ),
                  Container(
                    width: containWidth - 20,
                    height: this.seckillHeight - topHeight,
                    child: ListView.builder(
                      padding: EdgeInsets.all(5),
                      scrollDirection: Axis.horizontal,
                      itemCount: seckillModel.seckill.items.length,
                      itemBuilder: (BuildContext context, int index) {
                        ZNKSeckillItem item = seckillModel.seckill.items[index];
                        return GestureDetector(
                          onTap: () {},
                          child: Column(
                            children: [
                              Container(
                                width: (containWidth - 20) / 3.0,
                                child: (item.path.startsWith('http://') ||
                                        item.path.startsWith('https://'))
                                    ? CachedNetworkImage(imageUrl: item.path)
                                    : Image.asset(
                                        item.path,
                                        fit: BoxFit.scaleDown,
                                      ),
                              ),
                              Container(),
                              Container(),
                              Container(),
                            ],
                          ),
                        );
                      },
                    ),
                  ),
                ],
              ),
            )
          : Container(),
    );
  }
}
