import 'package:flb/model/user.dart';
import 'sqlitedb.dart';
///UserModel 用户模型
class UserDB {
  /* 数据表格名 */
  static final _dbName = 'user';

  /* 模型转换 */
  Map<String, dynamic> _toMap(User user) {
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
  User _fromMap(Map<String, dynamic> map) {
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
    return user;
  }

  /* 创建用户名 */
  Future<void> createUser() async {
    await SqliteDB.shared.createTable('''
    $_dbName (
      userID text not null primary key,
      sessionID text not null,
      account text not null,
      nickname text not null,
      photo text not null,
      phone text not null, 
      email text not null,
      createdAt text not null,
      updatedAt text not null,
    )
    ''');
  }

  /* 插入或更新数据 */
  Future<int> upsertUser(User user) async {
    int state = await SqliteDB.shared.upsert(
      _dbName, 
      _toMap(user)
    );
    return state;
  }

  /* 删除指定用户 */
  Future<int> deleteUser(String userID) async {
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
    return userMap == null ? null : _fromMap(userMap);
  }
  /* 上次登录的用户 */
  Future<User> lastLoginUser() async {
    List<Map<String, dynamic>> userMaps = await SqliteDB.shared.find(
      _dbName,
      orderBy: 'updatedAt DESC',
    );
    Map<String, dynamic> userMap = userMaps.first;
    return userMap == null ? null : _fromMap(userMap);
  }

  /* 当前登录用户 */
  Future<User> currentUser() async {
    List<Map<String, dynamic>> userMaps = await SqliteDB.shared.find(
      _dbName,
      orderBy: 'updatedAt DESC',
    );
    Map<String, dynamic> userMap = userMaps.first;
    return userMap == null ? null : _fromMap(userMap);
  }

}

