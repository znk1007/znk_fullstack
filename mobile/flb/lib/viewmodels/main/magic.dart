import 'package:flb/api/api.dart';
import 'package:flb/models/main/magic.dart';
import 'package:flb/util/config/help.dart';
import 'package:flb/util/http/core/request.dart';
import 'package:flb/viewmodels/base.dart';
import 'package:flutter/material.dart';

class ZNKMagicViewModel extends ZNKBaseViewModel {
  final ZNKApi api;
  ZNKMagicViewModel({@required this.api}) : super(api: api);

  //魔方栏数据
  List<ZNKMagic> _magics = [];
  List<ZNKMagic> get magics => _magics;
  //获取魔方栏数据
  Future<void> fetch() async {
    if (this.api.magicUrl.length == 0) {
      _defaultMagicData();
      return;
    }
    ResponseResult result = await RequestHandler.get(this.api.magicUrl);
    result.code = -1;
    List<Map<String, dynamic>> data = [];
    if (result.statusCode == 200 && result.data != null) {
      String code = result.data['code'];
      result.code = int.parse(code);
      data = result.data['data'];
      if (data.length <= 0) {
        result.code = -1;
      }
    }
    if (result.code != 0) {
      _defaultMagicData();
      return;
    }
    List<ZNKMagic> datas = [];
    for (var i = 0; i < data.length; i++) {
      Map<String, String> mData = data[i];
      ZNKMagic magic = ZNKMagic(
          identifier: ZNKHelp.safeString(mData['id']),
          path: ZNKHelp.safeString(mData['path']));
      datas.add(magic);
    }
    _magics = datas;
    notifyListeners();
  }

  //默认魔方栏数据
  void _defaultMagicData() {
    if (_magics.length != 0) {
      return;
    }
    _magics = [
      ZNKMagic(identifier: '1', path: 'lib/resource/sample.jpg'),
      ZNKMagic(identifier: '2', path: 'lib/resource/sample.jpg'),
      ZNKMagic(identifier: '3', path: 'lib/resource/sample.jpg'),
      ZNKMagic(identifier: '4', path: 'lib/resource/sample.jpg'),
      ZNKMagic(identifier: '5', path: 'lib/resource/sample.jpg'),
    ];
    notifyListeners();
  }
}
