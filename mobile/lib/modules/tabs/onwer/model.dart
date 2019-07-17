import 'package:znk/utils/database/user.dart';

class OnwerModel {
  String icon;
  String title;
  String detail;

  static Future<List<OnwerModel>> generate() async {
    List<OnwerModel> models = [];
    UserModel user = await UserDB.dao.current;
    final photo = user?.user?.photo;
    var nickname = user?.user?.nickname;
    final account = user.user.account;
    if (nickname == account) {
      nickname = '昵称';
    }
    var model = OnwerModel()
      ..icon = photo ?? ''
      ..title = nickname ?? ''
      ..detail = account ?? '';
  }
}