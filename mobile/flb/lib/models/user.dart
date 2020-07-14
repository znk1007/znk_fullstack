import 'package:flb/util/db/preferences/cache.dart';
import 'package:flb/util/db/sqlite/user/user.dart';
import 'package:flutter/material.dart';

//保存当前userID
const String _userIDKey = 'userLoginStatusKey';

class User {
  //会话ID
  String sessionID;
  //用户ID
  String userID;
  //账号
  String account;
  //真实姓名
  String realName;
  //手机号
  String phone;
  //头像
  String photo;
  //地址
  String address;
  //公司名称
  String compName;
  //邀请码
  String inviteCode;
  //身份类型 1:公司,2个人
  int identifyType;
  //注册时间
  String createdAt;
  //登录时间
  String updatedAt;
}

class UserModel extends ChangeNotifier {
  //用户ID
  String _userID = '';
  bool get logined => _userID.length > 0;

  /* 当前用户 */
  User get currentUser => _user;
  //用户信息
  User _user;

  //加载当前用户信息
  Future<void> loadUserData() async {
    if (_user != null) {
      return _user;
    }
    _userID = await ZNKCache.getValue(_userIDKey);
    _user = await UserDB.findUser(_userID);
  }

  //更新用户数据
  Future<void> upsert(User user) async {
    int stat = await UserDB.upsertUser(user);
    if (stat == 1) {
      //登录
      await ZNKCache.setValue(_userIDKey, user.userID);
      notifyListeners();
    }
  }

  //delete 输出用户数据
  Future<void> delete(User user) async {
    int stat = await UserDB.deleteUser(user.userID);
    if (stat == 1) {
      await ZNKCache.remove(_userIDKey);
      _user = null;
      _userID = '';
      notifyListeners();
    }
  }
}
