import 'package:flb/api/api.dart';
import 'package:flb/models/main/banner.dart';
import 'package:flb/util/config/help.dart';
import 'package:flb/util/http/core/request.dart';
import 'package:flb/viewmodels/base.dart';
import 'package:flutter/material.dart';

class ZNKBannerViewModel extends ZNKBaseViewModel {
  final ZNKApi api;
  ZNKBannerViewModel({@required this.api}) : super(api: api);

  //广告数据
  List<ZNKBannerModel> _banners = [];
  List<ZNKBannerModel> get banners => _banners;

  //拉取广告数据
  Future<void> fetch() async {
    if (this.api.msgNumUrl.length == 0) {
      _defaultBannerData();
      return;
    }
    ResponseResult result = await RequestHandler.get(this.api.bannerUrl);
    result.code = -1;
    List<Map> data = [];
    if (result.statusCode == 200 && result.data != null) {
      String code = result.data['code'];
      result.code = int.parse(code);
      data = result.data['data'];
      if (data.length == 0) {
        result.code = -1;
      }
    }
    if (result.statusCode != 0) {
      _defaultBannerData();
      return;
    }
    List<ZNKBannerModel> temp = [];
    for (var i = 0; i < data.length; i++) {
      Map<String, dynamic> dataMap = data[i];
      String id = ZNKHelp.safeString(dataMap['id']);
      String path = ZNKHelp.safeString(dataMap['path']);
      if (id.length > 0 && path.length > 0) {
        ZNKBannerModel m = ZNKBannerModel(identifier: id, path: path);
        temp.add(m);
      }
    }
    _banners = temp;
    notifyListeners();
  }

  //默认数据
  void _defaultBannerData() {
    _banners = [
      ZNKBannerModel(identifier: '1', path: 'lib/resource/test_img_01.jpg'),
      ZNKBannerModel(identifier: '2', path: 'lib/resource/test_img_02.jpg'),
      ZNKBannerModel(identifier: '3', path: 'lib/resource/test_img_03.jpg'),
    ];
    notifyListeners();
  }
}
