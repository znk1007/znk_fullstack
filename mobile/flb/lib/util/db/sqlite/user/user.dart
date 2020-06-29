
import 'package:flb/util/db/protos/generated/user/user.pb.dart';

import '../sqlitedb.dart';

///UserModel 用户模型
class UserDB  {
  /* 数据表格名 */
  static final _dbName = 'user';

  /* 模型转换 */
  static Map<String, dynamic> _toMap(User user) {
    var map = new Map<String, dynamic>();
    map['userID'] = user.userID;
    map['account'] = user.account;
    map['nickname'] = user.nickname;
    map['photo'] = user.photo;
    map['phone'] = user.phone;
    map['email'] = user.email;
    map['createdAt'] = user.createdAt;
    map['updatedAt'] = DateTime.now().toString();
    map['status'] = user.status;
    return map;
  }

  /* 字典转模型 */
  static User _fromMap(Map<String, dynamic> map) {
    User user = new User();
    user.userID = map['userID'] ?? '';
    user.sessionID = map['sessionID'] ?? '';
    user.account = map['account'] ?? '';
    user.nickname = map['nickname'] ?? '';
    user.photo = map['photo'] ?? '';
    user.phone = map['phone'] ?? '';
    user.email = map['email'] ?? '';
    user.createdAt = map['createdAt'] ?? '';
    user.updatedAt = map['updatedAt'] ?? '';
    user.status = map['status'] ?? 0;
    return user;
  }

  /* 创建用户表 */
  static Future<void> createUserTable() async {
    await SqliteDB.shared.createTable('''
    $_dbName (
      userID TEXT not null primary key,
      sessionID TEXT not null,
      account TEXT not null,
      nickname TEXT not null,
      photo TEXT not null,
      phone TEXT not null, 
      email TEXT not null,
      createdAt TEXT not null,
      updatedAt TEXT not null,
      status INTEGER,
    )
    ''');
  }

  /* 插入或更新数据 */
  static Future<int> upsertUser(User user) async {
    int state = await SqliteDB.shared.upsert(
      _dbName, 
      _toMap(user)
    );
    return state;
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
  static Future<User> findUser(String userID) async {
    List<Map<String, dynamic>> userMaps = await SqliteDB.shared.find(
      _dbName,
      where: 'userID = ?',
      whereArgs: [userID],
    );
    Map<String, dynamic> userMap = userMaps.first;
    return userMap == null ? null : _fromMap(userMap);
  }
  /* 上次登录的用户 */
  static Future<User> lastLoginUser() async {
    List<Map<String, dynamic>> userMaps = await SqliteDB.shared.find(
      _dbName,
      orderBy: 'updatedAt DESC',
    );
    Map<String, dynamic> userMap = userMaps.first;
    return userMap == null ? null : _fromMap(userMap);
  }

  /* 当前登录用户 */
  static Future<User> currentUser() async {
    List<Map<String, dynamic>> userMaps = await SqliteDB.shared.find(
      _dbName,
      where: 'status = 1',
      orderBy: 'updatedAt DESC',
    );
    Map<String, dynamic> userMap = userMaps.first;
    return userMap == null ? null : _fromMap(userMap);
  }

}

