import 'package:flb/pkg/banner/banner.dart';
import 'package:flb/pkg/screen/screen.dart';
import 'package:flb/pkg/search/search.dart';
import 'package:flb/viewmodels/main/msgnum.dart';
import 'package:flb/viewmodels/main/recommend.dart';
import 'package:flb/views/base/base.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class ZNKSearchView extends StatelessWidget {
  final Size searchSize = Size(ZNKScreen.screenWidth - 70, 31.0);
  ZNKSearchView({Key key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Row(
            children: [
              ZNKBaseView<ZNKMainRecommand>(
                model: ZNKMainRecommand(api: Provider.of(context)),
                onReady: (model) async {
                  model.fetch();
                },
                builder: (context, model, child) {
                  return model.recommends.length > 0
                      ? ZNKSearch(
                          style: ZNKSearchStyle(
                            enabled: false,
                            backgroudColor: Colors.white,
                            cornerRadius: searchSize.height / 2.0,
                            margin: EdgeInsets.only(
                                left: 14, top: ZNKScreen.safeTopArea),
                            width: searchSize.width,
                            height: searchSize.height,
                          ),
                          child: ZNKBanner(
                            size: Size(searchSize.width, searchSize.height),
                            decoration: BoxDecoration(
                                borderRadius: BorderRadius.all(
                                    Radius.circular(searchSize.height / 2.0))),
                            banners: model.recommends
                                .map((e) => Container(
                                      child: Text(
                                        e,
                                        style:
                                            TextStyle(color: Color(0xFF999999)),
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
              ),
              ZNKBaseView<ZNKMsgViewModel>(
                model: ZNKMsgViewModel(api: Provider.of(context)),
                onReady: (model) async {
                  model.fetch();
                },
                builder: (context, model, child) {
                  double msgSize = 12;
                  double msgIconSize = 23;
                  return Stack(
                    children: [
                      Container(
                        child: Icon(Icons.message),
                        width: msgIconSize,
                        height: msgIconSize,
                        margin: EdgeInsets.only(
                            left: 16,
                            top: (searchSize.height +
                                    ZNKScreen.safeTopArea -
                                    msgIconSize / 2.0) /
                                2.0),
                      ),
                      Container(
                        width: msgSize,
                        height: msgSize,
                        margin: EdgeInsets.only(
                            left: msgIconSize + msgSize,
                            top: (searchSize.height) / 2.0),
                        decoration: BoxDecoration(
                          color: Colors.red,
                          borderRadius: BorderRadius.circular(msgSize / 2.0),
                        ),
                        child: Text(model.msgNum,
                            textAlign: TextAlign.center,
                            style:
                                TextStyle(fontSize: 10, color: Colors.white)),
                      ),
                    ],
                  );
                },
              )
            ],
          );
  }
}