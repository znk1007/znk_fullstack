import 'package:flb/api/api.dart';
import 'package:flb/models/main/nav.dart';
import 'package:flb/util/config/help.dart';
import 'package:flb/util/http/core/request.dart';
import 'package:flb/viewmodels/base.dart';
import 'package:flutter/material.dart';

class ZNKNavViewModel extends ZNKBaseViewModel {
  final ZNKApi api;
  ZNKNavViewModel({@required this.api}) : super(api: api);

  //导航栏数据
  List<ZNKNav> _navs = [];
  List<ZNKNav> get navs => _navs;

  //拉取导航栏数据
  Future<void> fetch() async {
    if (this.api.navUrl.length == 0) {
      _defaultNavData();
      return;
    }
    ResponseResult result = await RequestHandler.get(this.api.navUrl);
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
      _defaultNavData();
      return;
    }
    List<ZNKNav> temp = [];
    for (var i = 0; i < data.length; i++) {
      Map<String, dynamic> dataMap = data[i];
      String id = ZNKHelp.safeString(dataMap['id']);
      String path = ZNKHelp.safeString(dataMap['path']);
      String title = ZNKHelp.safeString(dataMap['title']);
      if (id.length > 0 && path.length > 0 && title.length > 0) {
        ZNKNav m = ZNKNav(identifier: id, path: path, title: title);
        temp.add(m);
      }
    }
    _navs = temp;
    notifyListeners();
  }

  //默认数据
  void _defaultNavData() {
    if (_navs.length > 0) {
      return;
    }
    _navs = [
      ZNKNav(
          identifier: '1', path: 'lib/resource/collection.jpg', title: '标题标题一'),
      ZNKNav(
          identifier: '2', path: 'lib/resource/collection.jpg', title: '标题标题二'),
      ZNKNav(
          identifier: '3', path: 'lib/resource/collection.jpg', title: '标题标题三'),
      ZNKNav(
          identifier: '4', path: 'lib/resource/collection.jpg', title: '标题标题四'),
      ZNKNav(
          identifier: '5', path: 'lib/resource/collection.jpg', title: '标题标题五'),
      ZNKNav(
          identifier: '6', path: 'lib/resource/collection.jpg', title: '标题标题六'),
      ZNKNav(
          identifier: '7', path: 'lib/resource/collection.jpg', title: '标题标题七'),
      ZNKNav(
          identifier: '8', path: 'lib/resource/collection.jpg', title: '标题标题八'),
      ZNKNav(
          identifier: '9', path: 'lib/resource/collection.jpg', title: '标题标题九'),
      ZNKNav(
          identifier: '10',
          path: 'lib/resource/collection.jpg',
          title: '标题标题十'),
      ZNKNav(
          identifier: '11',
          path: 'lib/resource/collection.jpg',
          title: '标题标题十一'),
      ZNKNav(
          identifier: '12',
          path: 'lib/resource/collection.jpg',
          title: '标题标题十二'),
      ZNKNav(
          identifier: '13',
          path: 'lib/resource/collection.jpg',
          title: '标题标题十三'),
      ZNKNav(
          identifier: '14',
          path: 'lib/resource/collection.jpg',
          title: '标题标题十四'),
      ZNKNav(
          identifier: '15',
          path: 'lib/resource/collection.jpg',
          title: '标题标题十五'),
      ZNKNav(
          identifier: '16',
          path: 'lib/resource/collection.jpg',
          title: '标题标题十六'),
      ZNKNav(
          identifier: '17',
          path: 'lib/resource/collection.jpg',
          title: '标题标题十七'),
      ZNKNav(
          identifier: '18',
          path: 'lib/resource/collection.jpg',
          title: '标题标题十八'),
      ZNKNav(
          identifier: '19',
          path: 'lib/resource/collection.jpg',
          title: '标题标题十九'),
      ZNKNav(
          identifier: '20',
          path: 'lib/resource/collection.jpg',
          title: '标题标题二十'),
    ];
    notifyListeners();
  }
}
