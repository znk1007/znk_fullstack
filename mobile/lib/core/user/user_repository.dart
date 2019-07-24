import 'dart:async';
import 'dart:io';

import 'package:flutter/material.dart';
import 'package:znk/core/requests/login_request.dart';
import 'package:znk/core/requests/logout_request.dart';
import 'package:znk/core/requests/register_request.dart';
import 'package:znk/core/requests/update_online.dart';
import 'package:znk/core/user/check_userId.dart';
import 'package:znk/protos/generated/project/user.pb.dart';
import 'package:znk/utils/database/user.dart';


enum UserRepositoryState {
  none,
  checkUserIdFailed,
  unregisted,
  registFailed,
  unlogined,
  loginFailed,
  unkown,
  paramsEmpty,
  logoutFailed,
}

class UserRepositoryResult extends Error {
  final UserRepositoryState type;
  final String description;
  UserRepositoryResult({this.type = UserRepositoryState.none, this.description = ''});

  @override
  String toString() {
    return this.description.isEmpty ? '' : 'error msg: ${this.description}';
  }
}

class UserRepository {
  // 是否已登录
  Future<bool> isSignedIn() async {
    UserModel current = await UserDB.dao.current;
    return current != null && current.isLogined;
  }

  // 注册
  Future<bool>signUp(BuildContext ctx, {@required String account, @required String password}) async {
    Register r = new Register(account: account, password: password);
    final res = await r.regist(ctx);
    final succ = res != null && res.status == 1;
    print('sign up res: $res');
    if (succ) {
      final user = User()
        ..account = res.account
        ..userId = res.userId;
      UserDB.dao.upsert(user, false);
    }
    return succ;
  }
  Future<UserRepositoryResult>signIn(BuildContext ctx, {@required String account, @required String password}) async {
    final user = await UserDB.dao.findByAccount(account);
    String userId = '';
    if (user == null) {
      CheckUserId c = CheckUserId(account: account);
      final cRes = await c.check(ctx);
      if (cRes == null) {
        return UserRepositoryResult(type: UserRepositoryState.unkown, description: 'error occur when check userId');
      }
      if (cRes.status != 1) {
        return UserRepositoryResult(type: UserRepositoryState.checkUserIdFailed, description: 'error occur when check userId: ${cRes.message}');
      }
      userId = cRes.userId;
    } else {
      userId = user.user.userId;
    }
    if (userId.isEmpty) {
      return UserRepositoryResult(type: UserRepositoryState.paramsEmpty, description: 'userId cannot be empty');
    }
    Login login = Login(account: account, userId: userId, password: password);
    final res = await login.login(ctx);
    if (res.status != 1) {
      return UserRepositoryResult(type: UserRepositoryState.loginFailed, description: 'login failed: ${res.message}');
    }
    await UserDB.dao.upsert(res.user, true);
    return  UserRepositoryResult();
  }

  // 退出登录
  Future<UserRepositoryResult>signOut(BuildContext ctx) async {
    final curUser = await UserDB.dao.current;
    print('cur user: $curUser');
    if (curUser == null || curUser.user == null) {
      return UserRepositoryResult(type: UserRepositoryState.unlogined, description: 'user unlogined');
    }
    final userId = curUser.user.userId;
    final sessionId = curUser.user.sessionId;
    if (userId.isEmpty || sessionId.isEmpty) {
      return UserRepositoryResult(type: UserRepositoryState.paramsEmpty, description: 'params cannot be empty');
    }
    Logout l = Logout(userId: userId, sessionId: sessionId);
    await l.logout(ctx);
    UserDB.dao.updateLoginState(false, userId);
    return UserRepositoryResult(type: UserRepositoryState.none, description: 'logout success');
  }
  // 获取用户id
  Future<String>getUserId() async {
    UserModel current = await UserDB.dao.current;
    if (current == null) {
      return '';
    }
    return current.user.userId;
  }
  
  // updateOnline 更新在线状态
  Future<bool> updateOnline(BuildContext ctx, bool online) async {
    UserModel current = await UserDB.dao.current;
    if (current == null) {
      return false;
    }
    UpdateOnline updateOnline = UpdateOnline(
      userId: current.user.userId,
      account: current.user.account,
      sessionId: current.user.sessionId,
      online: online,
    );
    final res = await updateOnline.update(ctx);
    return res.status == 1;
  }

}



