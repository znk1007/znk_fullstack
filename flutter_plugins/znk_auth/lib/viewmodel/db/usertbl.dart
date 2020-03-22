import 'package:flutter/material.dart';
import 'package:znk_auth/model/protos/generated/auth/user.pb.dart';
import 'sqlitedb.dart';
///UserModel 用户模型
class UserTBL with ChangeNotifier {

  static final _dbName = 'user';

  /* 模型转换 */
  static Map<String, dynamic> toMap(User user) {
    var map = new Map<String, dynamic>();
    map['userID'] = user.userID;
    map['account'] = user.account;
    map['nickname'] = user.nickname;
    map['photo'] = user.photo;
    map['phone'] = user.phone;
    map['email'] = user.email;
    map['createdAt'] = user.createdAt;
    map['updatedAt'] = DateTime.now().toString();
    return map;
  }

  /* 字典转模型 */
  static User fromMap(Map<String, dynamic> map) {
    User user = new User();
    user.userID = map['userID'] ?? '';
    user.account = map['account'] ?? '';
    user.nickname = map['nickname'] ?? '';
    user.photo = map['photo'] ?? '';
    user.phone = map['phone'] ?? '';
    user.email = map['email'] ?? '';
    user.createdAt = map['createdAt'] ?? '';
    user.updatedAt = map['updatedAt'] ?? '';
    user.isOnline = map['isOnline'] ?? 0;
    return user;
  }

  /* 创建用户名 */
  static Future<void> createUserTBL() async {
    await SqliteDB.shared.createTable('''
    $_dbName (
      userID text not null primary key,
      account text not null,
      nickname text not null,
      photo text not null,
      phone text not null, 
      email text not null,
      createdAt text not null,
      updatedAt text not null,
      isOnline integer,
    )
    ''');
  }

  /* 插入货更新数据 */
  static Future<int> upsertUser(User user) async {
    return await SqliteDB.shared.upsert(
      _dbName, 
      UserTBL.toMap(user)
    );
  }

  /* 删除指定用户 */
  static Future<int> deleteUser(String userID) async {
    return await SqliteDB.shared.delete(
      _dbName, 
      where: 'userID = ?', 
      whereArgs: [userID]
    );
  }
  /* 查找用户 */
  Future<User> findUser(String userID) async {
    List<Map<String, dynamic>> userMaps = await SqliteDB.shared.find(
      _dbName,
      where: 'userID = ?',
      whereArgs: [userID],
    );
    Map<String, dynamic> userMap = userMaps.first;
    if (userMap == null) {
      return null;
    }
    return UserTBL.fromMap(userMap);
  }
  /* 上次登录的用户 */
  static Future<User> lastLoginUser() async {
    List<Map<String, dynamic>> userMaps = await SqliteDB.shared.find(
      _dbName,
      orderBy: 'updatedAt DESC',
    );
    Map<String, dynamic> userMap = userMaps.first;
    if (userMap == null) {
      return null;
    }
    return UserTBL.fromMap(userMap);
  }

  /* 当前登录用户 */
  static Future<User> currentUser() async {
    List<Map<String, dynamic>> userMaps = await SqliteDB.shared.find(
      _dbName,
      where: 'isOnline = 1',
      orderBy: 'updatedAt DESC',
    );
    Map<String, dynamic> userMap = userMaps.first;
    if (userMap == null) {
      return null;
    }
    return UserTBL.fromMap(userMap);
  }

}

