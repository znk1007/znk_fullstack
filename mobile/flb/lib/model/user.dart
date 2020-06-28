import 'package:hive/hive.dart';

@HiveType(typeId: 0)
class User extends HiveObject {
  //用户ID
  @HiveField(0)
  String userID;
  //账号
  @HiveField(1)
  String account;
  //昵称
  @HiveField(2)
  String nickname;
  //头像
  @HiveField(3)
  String photo;
  //手机号
  @HiveField(4)
  String phone;
  //邮箱
  @HiveField(5)
  String email;
  //创建日期
  @HiveField(6)
  String createdAt;
  //更新日期
  @HiveField(7)
  String updatedAt;
}