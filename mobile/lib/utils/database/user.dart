

import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:znk/protos/generated/project/user.pb.dart';
import 'package:znk/utils/database/base/sembastdb.dart';



class UserModel {
  User user;
  bool isLogined;
  String updatedAt;
  factory UserModel.fromUser(User user, bool isLogined) {
    return UserModel(user: user, isLogined: isLogined, updatedAt: DateTime.now().toLocal().millisecondsSinceEpoch.toString());
  }

  UserModel({@required this.user, this.isLogined, this.updatedAt});
  // 转JSON
  Map<String, dynamic> toJson() => {
    'userId': user.userId,
    'sessionId': user.sessionId,
    'account': user.account,
    'nickname': user.nickname,
    'phone': user.phone,
    'email': user.email,
    'photo': user.photo,
    'createdAt': user.createdAt,
    'updatedAt': updatedAt,
    'isLogined': isLogined,
  };

  // Map转模型
  UserModel.fromMap(Map<String, dynamic> json){
    if (user == null) {
      user = User();
    }
    user.userId = json['userId'];
    user.sessionId = json['sessionId'];
    user.account = json['account'];
    user.nickname = json['nickname'];
    user.phone = json['phone'];
    user.email = json['email'];
    user.photo = json['photo'];
    user.createdAt = json['createdAt'] ?? '';
    updatedAt = json['updatedAt'];
    isLogined = json['isLogined'];
  }
}

class UserDB {
  // 单例
  static UserDB _instance;
  static UserDB get dao {
    if (_instance == null) {
      _instance = UserDB._();
    }
    return _instance;
  }

  SembastDB _dbClient;

  UserDB._() {
    _init();
  }

  void _init() async {
    this._dbClient = SembastDB('users.db', '_user');
  }

  // 关闭客户端
  Future close() async {
    await this._dbClient.closeDB();
  }

  // 插入/更新用户
  Future<void> upsert(User user, bool isLogined) async {
    UserModel model = UserModel.fromUser(user, isLogined);
    print('isLogined $isLogined');
    
    if (model.updatedAt == '') {
      model.updatedAt = DateTime.now().toLocal().millisecondsSinceEpoch.toString();
    }
    bool exists = await this._dbClient.isRecordExists(user.userId);
    if (model.user.userId == '' || exists == true) {
      final val = await this._dbClient.updateRecord(user.userId, model.toJson());
      print('update val: $val');
      return;
    }
    await this._dbClient.save<String, Map<String, dynamic>>(user.userId, model.toJson());
  }
  // 删除用户
  Future<void> delete(String userId) async {
    await this._dbClient.deleteRecord(userId);
  }
  // 查找指定用户
  Future<UserModel> find(String userId) async {
    var userJson = await this._dbClient.fetch<Map<String, dynamic>>(userId);
    final u = UserModel.fromMap(userJson);
    return u;
  }
  // findByAccount 根据账号查用户
  Future<UserModel> findByAccount(String account) async {
    var snapshot = await this._dbClient.findFirstEquals(field: 'account', value: account);
    if (snapshot == null) {
      return null;
    }
    return UserModel.fromMap(snapshot.value);
  }

  // 查找所有用户
  Future<List<UserModel>> findAll() async {
    var snapshots = await this._dbClient.findMatches();
    if (snapshots == null) {
      return null;
    }
    List<UserModel> users = [];
    try {
      snapshots.forEach((record){
        final u = UserModel.fromMap(record.value);
        users.add(u);
      });    
    } catch (e) {
      print('find all err: $e');
    }
    return users;
  }
  // 指定用户是否已登录
  Future<bool> isUserLogined(String userId) async {
    var user = await this.find(userId);
    return user.isLogined;
  }

  // 当前用户是否已登录
  Future<bool> get isCurrentLogined async {
    var curr = await this.current;
    return await this.isUserLogined(curr.user.userId);
  }

  // 更新登录状态
  Future updateLoginState(bool isLogined, String userId) async {
    var userModel = await this.find(userId);
    userModel.isLogined = isLogined;
    userModel.updatedAt = DateTime.now().millisecondsSinceEpoch.toString();
    await this._dbClient.updateRecord(userModel.user.userId, userModel.toJson());
  }

  // 当前用户
  Future<UserModel> get current async {
    var snapshot = await this._dbClient.firstRecord<String, dynamic>(sortFiled: 'updatedAt', ascending: false);
    if (snapshot == null) {
      return null;
    }
    return UserModel.fromMap(snapshot.value);
  }
  

}
