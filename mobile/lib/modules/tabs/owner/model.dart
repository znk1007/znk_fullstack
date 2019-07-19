import 'package:flutter/widgets.dart';
import 'package:znk/images/manager.dart';
import 'package:znk/utils/database/user.dart';

enum OwnerType {
  person,
  fileStore,
  collection,
  setting
}

class OwnerModel {
  Image icon;
  String title;
  String detail;
  OwnerType type;
  // 生成数据
  Future<List<OwnerModel>> generate() async {
    List<OwnerModel> models = [];
    UserModel user = await UserDB.dao.current;
    final photo = user?.user?.photo ?? '';
    Image header = (photo.startsWith('http') || photo.startsWith('https')) ? 
      Image.network(photo) : 
      Image.asset(photo);
    var nickname = user?.user?.nickname;
    final account = user.user.account;
    if (nickname == account) {
      nickname = '昵称';
    }
    var model = OwnerModel()
      ..icon = header
      ..title = nickname ?? ''
      ..detail = account ?? ''
      ..type = OwnerType.person;
    models.add(model);

    model = OwnerModel()
      ..icon = OwnerAsset.fileStore
      ..title = '我的网盘'
      ..detail = ''
      ..type = OwnerType.fileStore;
    models.add(model);
    model = OwnerModel()
      ..icon = OwnerAsset.collection
      ..title = '我的收藏'
      ..detail = ''
      ..type = OwnerType.collection;
    models.add(model);

    model = OwnerModel()
      ..icon = OwnerAsset.setting
      ..title = '设置'
      ..detail = ''
      ..type = OwnerType.setting;
    models.add(model);
    return models;
  }
}