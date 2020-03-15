import 'package:flutter/material.dart';
import 'sqlitedb.dart';
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

  static final _dbName = 'user';

  /* 模型转换 */
  Map<String, dynamic> toMap() {
    var map = new Map<String, dynamic>();
    map['userId'] = userId;
    map['account'] = account;
    map['username'] = username;
    map['photo'] = photo;
    map['phone'] = phone;
    map['email'] = email;
    return map;
  }

  /* 字典转模型 */
  static UserModel fromMap(Map<String, dynamic> map) {
    UserModel userModel = new UserModel();
    userModel.userId = map['userId'] ?? '';
    userModel.account = map['account'] ?? '';
    userModel.username = map['username'] ?? '';
    userModel.photo = map['photo'] ?? '';
    userModel.phone = map['phone'] ?? '';
    userModel.email = map['email'] ?? '';
    return userModel;
  }

  /* 创建用户名 */
  Future<void> createUserTBL() async {
    await SqliteDB.shared.createTable('''
    $_dbName (
      userId text not null primary key,
      account text not null,
      username text not null,
      photo text not null,
      phone text not null, 
      email text not null,
    )
    ''');
  }

  /* 插入货更新数据 */
  Future<int> upsertUser(UserModel userModel) async {
    return await SqliteDB.shared.upsert(
      _dbName, 
      userModel.toMap()
    );
  }

  /* 删除指定用户 */
  Future<int> deleteUser(String userId) async {
    return await SqliteDB.shared.delete(
      _dbName, 
      where: 'userId = ?', 
      whereArgs: [userId]
    );
  }

  Future<UserModel> findUser(String userId) async {
    List<Map<String, dynamic>> users = await SqliteDB.shared.find(
      _dbName,
      where: 'userId = ?',
      whereArgs: [userId],
    );
    Map<String, dynamic> user = users.first;
    if (user == null) {
      return null;
    }
    return UserModel.fromMap(user);
  }

}

