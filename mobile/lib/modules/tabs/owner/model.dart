import 'package:flutter/widgets.dart';
import 'package:znk/images/manager.dart';
import 'package:znk/utils/database/user.dart';

class OwnerModel {
  Image icon;
  String title;
  String detail;
  // 生成数据
  Future<List<List<OwnerModel>>> generate() async {
    List<List<OwnerModel>> models = [];
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
    List<OwnerModel> temps = [];
    var model = OwnerModel()
      ..icon = header
      ..title = nickname ?? ''
      ..detail = account ?? '';
    temps.add(model);
    models.add(temps);

    temps = [];
    model = OwnerModel()
      ..icon = OwnerAsset.fileStore
      ..title = '我的网盘'
      ..detail = '';
    temps.add(model);
    model = OwnerModel()
      ..icon = OwnerAsset.collection
      ..title = '我的收藏'
      ..detail = '';
    temps.add(model);
    models.add(temps);

    temps = [];
    model = OwnerModel()
      ..icon = OwnerAsset.setting
      ..title = '设置'
      ..detail = '';
    temps.add(model);
    models.add(temps);
    return models;
  }
}