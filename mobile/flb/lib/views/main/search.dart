import 'package:flb/pkg/banner/banner.dart';
import 'package:flb/pkg/screen/screen.dart';
import 'package:flb/pkg/search/search.dart';
import 'package:flb/viewmodels/main/recommend.dart';
import 'package:flb/views/base/base.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class ZNKSearchView extends StatelessWidget {
  final bool show;
  final Size searchSize;
  ZNKSearchView({
    Key key,
    @required this.show,
    @required this.searchSize,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ZNKBaseView<ZNKRecommendViewModel>(
      model: ZNKRecommendViewModel(api: Provider.of(context)),
      onReady: (model) async => model.fetch(),
      builder: (context, model, child) {
        return (show && model.recommends.length > 0)
            ? ZNKSearch(
                style: ZNKSearchStyle(
                  enabled: false,
                  backgroudColor: Colors.white,
                  cornerRadius: this.searchSize.height / 2.0,
                  margin: EdgeInsets.only(left: 14, top: ZNKScreen.safeTopArea),
                  width: this.searchSize.width,
                  height: this.searchSize.height,
                ),
                child: ZNKBanner(
                  size: Size(this.searchSize.width, this.searchSize.height),
                  decoration: BoxDecoration(
                      borderRadius: BorderRadius.all(
                          Radius.circular(this.searchSize.height / 2.0))),
                  banners: model.recommends
                      .map((e) => Container(
                            child: Text(
                              e,
                              style: TextStyle(color: Color(0xFF999999)),
                            ),
                            margin: EdgeInsets.only(left: 40),
                          ))
                      .toList(),
                  showIndicator: false,
                  scrollDirection: Axis.vertical,
                  alignment: Alignment.centerLeft,
                  didSelected: (index) {
                    print('did selected: $index');
                  },
                ),
              )
            : Container();
      },
    );
  }
}
