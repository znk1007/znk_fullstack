import 'package:flutter/material.dart';
import 'package:userAuth/usercenter/sqlitedb.dart';
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

  Map<String, dynamic> toMap() {
    var map = new Map<String, dynamic>();
    
  }

  /* 创建用户名 */
  Future<void> createUserTBL() async {
    await SqliteDB.shared.createTable('''
    user (
      userId text not null primary key,
      account text not null,
      username text not null,
      photo text not null,
      phone text not null, 
      email text not null,
    )
    ''');
  }

  Future<int> upsertUser() async {

  }

}

