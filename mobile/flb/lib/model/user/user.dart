import 'package:flb/util/db/protos/generated/user/user.pb.dart';
import 'package:flb/util/db/sqlite/user/user.dart';
import 'package:flutter/material.dart';

class UserModel extends ChangeNotifier {
  //用户信息
  User _user;

  //当前用户
  Future<User> get current async {
    if (_user != null) {
      return _user;
    }
    _user = await UserDB.currentUser();
    return _user;
  }
  //更新用户数据
  Future<void> upsert(User user) async {
    int stat = await UserDB.upsertUser(user);
    if (stat == 1) {
      notifyListeners();
    }
  }
  //delete 输出用户数据
  Future<void> delete(User user) async {
    int stat = await UserDB.deleteUser(user.userID);
    if (stat == 1) {
      notifyListeners();
    }
  }

}