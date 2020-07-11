import 'package:flb/models/style/style.dart';
import 'package:flb/viewmodels/main/msgnum.dart';
import 'package:flb/views/base/base.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class ZNKMsgNumView extends StatelessWidget {
  final ThemeStyle style;
  final double marginTop;
  final bool show;
  const ZNKMsgNumView({
    Key key,
    @required this.style,
    @required this.marginTop,
    @required this.show,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    double msgSize = 12;
    double msgIconSize = 23;
    return ZNKBaseView<ZNKMsgNumViewModel>(
      model: ZNKMsgNumViewModel(api: Provider.of(context)),
      onReady: (model) async => model.fetch(),
      builder: (context, model, child) {
        return !this.show
            ? Container()
            : Container(
                width: msgSize,
                height: msgSize,
                margin: EdgeInsets.only(
                  left: msgIconSize + msgSize,
                  top: this.marginTop,
                ),
                decoration: BoxDecoration(
                  color: style.redColor,
                  borderRadius: BorderRadius.circular(msgSize / 2.0),
                ),
                child: Text(
                  model.msgNum,
                  textAlign: TextAlign.center,
                  style: TextStyle(
                    fontSize: 10,
                    color: Colors.white,
                  ),
                ),
              );
      },
    );
  }
}
