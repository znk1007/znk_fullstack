import 'package:flb/models/style/style.dart';
import 'package:flb/pkg/screen/screen.dart';
import 'package:flb/util/random/random.dart';
import 'package:flb/viewmodels/main/seckill.dart';
import 'package:flb/views/base/base.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class ZNKSeckillView extends StatefulWidget {
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
  _ZNKSeckillViewState createState() => _ZNKSeckillViewState();
}

class _ZNKSeckillViewState extends State<ZNKSeckillView> {
  @override
  Widget build(BuildContext context) {
    return ZNKBaseView<ZNKSeckillViewModel>(
      model: ZNKSeckillViewModel(
        api: Provider.of(context),
      ),
      onReady: (model) async => model.fetch(),
      builder: (context, model, child) => (widget.show &&
              model.seckill != null &&
              model.seckill.items.length > 0)
          ? Container(
              width: ZNKScreen.screenWidth / 2.0,
              height: widget.seckillHeight,
              color: RandomHandler.randomColor,
              child: Column(
                children: [
                  Row(
                    children: [
                      Container(),
                      Container(),
                    ],
                  ),
                ],
              ),
            )
          : Container(),
    );
  }
}
