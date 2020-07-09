import 'package:flb/pkg/screen/screen.dart';
import 'package:flb/state/home.dart';
import 'package:flb/viewmodels/home.dart';
import 'package:flb/views/base/base.dart';
import 'package:flb/views/base/launch.dart';
import 'package:flb/views/tabbar/tabbar.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class ZNKHomePage extends StatelessWidget {
  const ZNKHomePage({Key key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    //设置屏幕
    ZNKScreen.setContext(context);
    return ZNKBaseView<ZNKHomeModel>(
      model: ZNKHomeModel(api: Provider.of(context)),
      onReady: (model) async {
        model.fetch();
      },
      builder: (context, model, child) {
        return Container(
          child: (model.state == ZNKHomeLoadState.launching ||
                  model.items == null ||
                  model.items.length == 0)
              ? LaunchPage()
              : ZNKTabbar(items: model.items),
        );
      },
    );
  }
}
