import 'package:cached_network_image/cached_network_image.dart';
import 'package:flb/models/main/magic.dart';
import 'package:flb/pkg/screen/screen.dart';
import 'package:flb/viewmodels/main/magic.dart';
import 'package:flb/views/base/base.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class ZNKMagicView extends StatelessWidget {
  final bool show;
  final double magicHeight;
  const ZNKMagicView({
    Key key,
    @required this.show,
    @required this.magicHeight,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ZNKBaseView<ZNKMagicViewModel>(
      model: ZNKMagicViewModel(
        api: Provider.of(context),
      ),
      onReady: (model) async => model.fetch(),
      builder: (context, model, child) => (this.show && model.magics.length > 0)
          ? Container(
              margin: EdgeInsets.only(left: 10, top: 10, right: 10),
              height: magicHeight,
              child: ListView.builder(
                scrollDirection: Axis.horizontal,
                itemCount: model.magics.length,
                itemBuilder: (BuildContext context, int index) {
                  ZNKMagic magic = model.magics[index];
                  return ClipRRect(
                    borderRadius: BorderRadius.circular(6),
                    child: Container(
                      margin: EdgeInsets.only(left: index > 0 ? 3.5 : 0),
                      width: (ZNKScreen.screenWidth -
                              (model.magics.length - 1) * 3.5 -
                              ZNKScreen.setWidth(15)) /
                          3.0,
                      height: magicHeight,
                      child: (magic.path.startsWith('http://') ||
                              magic.path.startsWith('https://'))
                          ? CachedNetworkImage(
                              imageUrl: magic.path,
                              fit: BoxFit.cover,
                            )
                          : Image.asset(
                              magic.path,
                              fit: BoxFit.cover,
                            ),
                    ),
                  );
                },
              ),
            )
          : Container(),
    );
  }
}
