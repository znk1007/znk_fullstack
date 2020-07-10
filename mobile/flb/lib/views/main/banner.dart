import 'package:cached_network_image/cached_network_image.dart';
import 'package:flb/pkg/banner/banner.dart';
import 'package:flb/pkg/screen/screen.dart';
import 'package:flb/viewmodels/main/banner.dart';
import 'package:flb/views/base/base.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class ZNKBannerView extends StatelessWidget {
  final double height;
  const ZNKBannerView({Key key, @required this.height}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ZNKBaseView<ZNKBannerViewModel>(
      model: ZNKBannerViewModel(api: Provider.of(context)),
      onReady: (model) async {
        model.fetch();
      },
      builder: (context, model, child) {
        return ZNKBanner(
          indicatorTrackColor: Color(0xFFD73B1E), //ox后两位表示透明度
          indicatorTintColor: Colors.white,
          size: Size(ZNKScreen.screenWidth, this.height),
          banners: model.banners
              .map((e) => (e.path.startsWith('http://') ||
                      e.path.startsWith('https://'))
                  ? CachedNetworkImage(imageUrl: e.path)
                  : Image.asset(e.path,
                      fit: BoxFit.fill,
                      width: ZNKScreen.screenWidth,
                      height: this.height))
              .toList(),
          scrollDirection: Axis.horizontal,
          alignment: Alignment.centerLeft,
          didSelected: (index) {
            print('did selected: $index');
          },
        );
      },
    );
  }
}
