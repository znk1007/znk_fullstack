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
    return ZNKBaseView<ZNKSeckillViewModel>(
      model: ZNKSeckillViewModel(
        api: Provider.of(context),
      ),
      onReady: (model) async => model.fetch(),
      builder: (context, seckillModel, child) => (this.show &&
              seckillModel.seckill != null &&
              seckillModel.seckill.items.length > 0)
          ? Container(
              width: containWidth,
              height: this.seckillHeight,
              child: Column(
                children: [
                  Row(
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
                        padding: EdgeInsets.only(right: 10),
                        child: ZNKBaseView<ZNKSeckillCountDownViewModel>(
                          model: ZNKSeckillCountDownViewModel(
                            api: Provider.of(context),
                          ),
                          onReady: (countdown) => countdown
                              .fetch(int.parse(seckillModel.seckill.time)),
                          builder: (context, countdown, child) => Row(
                            children: [
                              Container(
                                child: Text(countdown.hour),
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
                                child: Text(countdown.minute),
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
                                child: Text(countdown.second),
                              ),
                            ],
                          ),
                        ),
                      ),
                    ],
                  ),
                ],
              ),
            )
          : Container(),
    );
  }
}
