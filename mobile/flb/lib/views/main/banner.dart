import 'package:cached_network_image/cached_network_image.dart';
import 'package:flb/pkg/banner/banner.dart';
import 'package:flb/pkg/screen/screen.dart';
import 'package:flb/viewmodels/main/banner.dart';
import 'package:flb/views/base/base.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class ZNKBannerView extends StatelessWidget {
  final bool show;
  final double bannerHeight;
  const ZNKBannerView({
    Key key,
    @required this.show,
    @required this.bannerHeight,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ZNKBaseView<ZNKBannerViewModel>(
      model: ZNKBannerViewModel(api: Provider.of(context)),
      onReady: (model) async => model.fetch(),
      builder: (context, model, child) => this.show
          ? ZNKBanner(
              indicatorTrackColor: Color(0xFFD73B1E), //ox后两位表示透明度
              indicatorTintColor: Colors.white,
              size: Size(ZNKScreen.screenWidth, bannerHeight),
              banners: model.banners
                  .map((e) => (e.path.startsWith('http://') ||
                          e.path.startsWith('https://'))
                      ? CachedNetworkImage(
                          imageUrl: e.path,
                          fit: BoxFit.fill,
                        )
                      : Image.asset(e.path,
                          fit: BoxFit.fill,
                          width: ZNKScreen.screenWidth,
                          height: bannerHeight))
                  .toList(),
              scrollDirection: Axis.horizontal,
              alignment: Alignment.centerLeft,
              didSelected: (index) {
                print('did selected: $index');
              },
            )
          : Container(),
    );
  }
}
