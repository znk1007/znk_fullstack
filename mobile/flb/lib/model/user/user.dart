import 'package:flb/util/db/protos/generated/user/user.pb.dart';
import 'package:flb/util/db/sqlite/user/user.dart';
import 'package:flutter/material.dart';

class UserModel extends ChangeNotifier {
  /* 当前用户 */
  User get currentUser => _user;

  //用户信息
  User _user;
  //是否已登录
  bool get isLogined => (_user != null && _user.status == 1);
  //是否测试
  bool _test = false;

  //加载当前用户信息
  Future<void> loadUserData() async {
    if (_test) {
      _user = User();
      _user.status = 1;
    }
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
      //登录
      notifyListeners();
    }
    if (user.status == 0) {
      _user = null;
    }
  }

  //delete 输出用户数据
  Future<void> delete(User user) async {
    int stat = await UserDB.deleteUser(user.userID);
    if (stat == 1) {
      notifyListeners();
    }
    if (user.status == 0) {
      _user = null;
    }
  }
}
