import 'package:flb/api/api.dart';
import 'package:flb/models/main/collection.dart';
import 'package:flb/util/config/help.dart';
import 'package:flb/util/http/core/request.dart';
import 'package:flb/viewmodels/base.dart';
import 'package:flutter/material.dart';

class ZNKCollectionViewModel extends ZNKBaseViewModel {
  final ZNKApi api;
  ZNKCollectionViewModel({@required this.api}) : super(api: api);

//广告数据
  List<ZNKCollection> _collections = [];
  List<ZNKCollection> get collections => _collections;

  //拉取广告数据
  Future<void> fetch() async {
    if (this.api.collectionUrl.length == 0) {
      _defaultData();
      return;
    }
    ResponseResult result = await RequestHandler.get(this.api.tabbarUrl);
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
      _defaultData();
      return;
    }
    List<ZNKCollection> temp = [];
    for (var i = 0; i < data.length; i++) {
      Map<String, dynamic> dataMap = data[i];
      String id = ZNKHelp.safeString(dataMap['id']);
      String path = ZNKHelp.safeString(dataMap['path']);
      String title = ZNKHelp.safeString(dataMap['title']);
      if (id.length > 0 && path.length > 0 && title.length > 0) {
        ZNKCollection m =
            ZNKCollection(identifier: id, path: path, title: title);
        temp.add(m);
      }
    }
    _collections = temp;
    notifyListeners();
  }

  //默认数据
  void _defaultData() {
    if (_collections.length > 0) {
      return;
    }
    _collections = [
      ZNKCollection(
          identifier: '1', path: 'lib/resource/collection.jpg', title: '标题一'),
      ZNKCollection(
          identifier: '2', path: 'lib/resource/collection.jpg', title: '标题二'),
      ZNKCollection(
          identifier: '3', path: 'lib/resource/collection.jpg', title: '标题三'),
      ZNKCollection(
          identifier: '4', path: 'lib/resource/collection.jpg', title: '标题四'),
      ZNKCollection(
          identifier: '5', path: 'lib/resource/collection.jpg', title: '标题五'),
      ZNKCollection(
          identifier: '6', path: 'lib/resource/collection.jpg', title: '标题六'),
      ZNKCollection(
          identifier: '7', path: 'lib/resource/collection.jpg', title: '标题七'),
      ZNKCollection(
          identifier: '8', path: 'lib/resource/collection.jpg', title: '标题八'),
      ZNKCollection(
          identifier: '9', path: 'lib/resource/collection.jpg', title: '标题九'),
      ZNKCollection(
          identifier: '10', path: 'lib/resource/collection.jpg', title: '标题十'),
      ZNKCollection(
          identifier: '11', path: 'lib/resource/collection.jpg', title: '标题十一'),
      ZNKCollection(
          identifier: '12', path: 'lib/resource/collection.jpg', title: '标题十二'),
      ZNKCollection(
          identifier: '13', path: 'lib/resource/collection.jpg', title: '标题十三'),
      ZNKCollection(
          identifier: '14', path: 'lib/resource/collection.jpg', title: '标题十四'),
      ZNKCollection(
          identifier: '15', path: 'lib/resource/collection.jpg', title: '标题十五'),
      ZNKCollection(
          identifier: '16', path: 'lib/resource/collection.jpg', title: '标题十六'),
      ZNKCollection(
          identifier: '17', path: 'lib/resource/collection.jpg', title: '标题十七'),
      ZNKCollection(
          identifier: '18', path: 'lib/resource/collection.jpg', title: '标题十八'),
      ZNKCollection(
          identifier: '19', path: 'lib/resource/collection.jpg', title: '标题十九'),
      ZNKCollection(
          identifier: '20', path: 'lib/resource/collection.jpg', title: '标题二十'),
    ];
    notifyListeners();
  }
}
