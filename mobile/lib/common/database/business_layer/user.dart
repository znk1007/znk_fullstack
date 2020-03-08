import 'package:flutter/material.dart';
import 'package:mobile02/common/database/base/database.dart';
import 'package:mobile02/model/user/user.dart';
import 'package:mobile02/protos/generated/project/user.pb.dart';

class UserDB {
  /// 单例
  static UserDB _instance;
  static UserDB get dao {
    if (_instance == null) {
      _instance = UserDB._();
    }
    return _instance;
  }
  /// 数据库句柄
  DBHelper _helper;

  /*
  companyId INTEGER,
    FOREIGN KEY (companyId) REFERENCES Company(id) 
    ON DELETE CASCADE
  */

  UserDB._(){
    String tableName = 'user';
    _helper = DBHelper(
      tableName: tableName,
      createSql: '''
            create table if not exists 
            $tableName(
              userId text primary key,
              sessionId text,
              account text,
              nickname text,
              phone text,
              email text,
              photo text,
              createdAt text,
              updatedAt text,
              isLogined integer
            ) '''
    );
  }
  /* 插入用户数据 */
  static Future insert(User user, int isLogined) async {
    DBHelper h = UserDB.dao._helper;
    if (h == null) {
      return;
    }
    UserModel userModel = UserModel.fromUser(user, isLogined);
    try {
      await h.insert(userModel.toJson());
      if (isLogined == 1) {
        userModel.notifyLoginState();
      }
    } catch (e) {
      print('insert user error: $e');
    }
  }

  /// 更新用户登录状态
  static Future updateUserLoginState(int isLogined) async {
    DBHelper h = UserDB.dao._helper;
    if (h == null) {
      return;
    }
    UserModel m = await currentUser;
    if (m == null) {
      return;
    }
    try {
      UserModel newM = m.updateUserLoginState(isLogined);
      h.update(newM.toJson());
    } catch (e) {
      print('update user login state e: $e');
    }
  }

  /// 当前用户
  static Future<UserModel> get currentUser async {
    DBHelper h = UserDB.dao._helper;
    if (h == null) {
      return null;
    }
    List<Map<String, dynamic>> userJsons = await h.query(
      orderBy: 'updatedAt ASC', 
    );
    if (userJsons.isEmpty) {
      return null;
    }
    Map<String, dynamic> userJson = userJsons.last;
    return UserModel.fromMap(userJson);
  }


}