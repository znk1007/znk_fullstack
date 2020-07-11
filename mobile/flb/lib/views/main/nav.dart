import 'package:flb/pkg/grid/grid.dart';
import 'package:flb/viewmodels/main/nav.dart';
import 'package:flb/views/base/base.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class ZNKNavView extends StatelessWidget {
  final bool show;
  final double navHeight;
  const ZNKNavView({
    Key key,
    @required this.show,
    @required this.navHeight,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ZNKBaseView<ZNKNavViewModel>(
      model: ZNKNavViewModel(api: Provider.of(context)),
      onReady: (model) async => model.fetch(),
      builder: (context, model, child) => (this.show && model.navs.length > 0)
          ? ZNKGrid(
              items: model.navs
                  .map(
                    (e) => ZNKGridItem(
                      identifier: e.identifier,
                      title: e.title,
                      path: e.path,
                    ),
                  )
                  .toList(),
            )
          : Container(),
    );
  }
}
