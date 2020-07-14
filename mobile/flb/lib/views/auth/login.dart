import 'package:flb/models/style/style.dart';
import 'package:flb/pkg/screen/screen.dart';
import 'package:flb/viewmodels/auth/login.dart';
import 'package:flb/views/base/base.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class ZNKLoginView extends StatelessWidget {
  const ZNKLoginView({Key key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    ThemeStyle style = context.watch<ThemeStyle>();
    return ZNKBaseView<ZNKLoginViewModel>(
      model: ZNKLoginViewModel(
        api: Provider.of(context),
      ),
      builder: (context, model, child) {
        return Stack(
          children: [
            _backgroundView(style),
          ],
        );
      },
    );
  }

  Widget _backgroundView(ThemeStyle style) {
    return Container(
      width: ZNKScreen.screenWidth,
      height: ZNKScreen.screenHeight,
      child: Stack(
        children: [
          Container(
            color: style.redColor,
          ),
        ],
      ),
    );
  }
}
