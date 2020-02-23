import 'package:flutter/material.dart';
///UserModel 用户模型
class UserModel with ChangeNotifier {
  ///用户唯一id
  String userId;
  ///用户账号
  String account;
  ///用户昵称
  String username;
  ///用户头像
  String photo;
  ///手机号
  String phone;
  ///邮箱
  String email;
}