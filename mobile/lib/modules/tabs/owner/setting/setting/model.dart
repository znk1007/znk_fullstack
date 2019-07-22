import 'package:flutter/foundation.dart';

enum SettingType {
  security,
  privacy,
  clean,
  version,
  logout,
}

class SettingModel {
  // 标题
  final String title;
  // 类型
  final SettingType type;
  SettingModel({@required this.title, @required this.type});
  // 生成数据模型
  static List<SettingModel> generate() {
    List<SettingModel> models = [];
    SettingModel model = SettingModel(title: '账号与安全', type: SettingType.security);
    models.add(model);
    model = SettingModel(title: '隐私', type: SettingType.privacy);
    models.add(model);
    model = SettingModel(title: '清理缓存', type: SettingType.clean);
    models.add(model);
    model = SettingModel(title: '当前版本', type: SettingType.version);
    models.add(model);
    model = SettingModel(title: '退出登录', type: SettingType.logout);
    models.add(model);
    return models;
  }
}