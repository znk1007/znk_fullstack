import 'package:flutter/material.dart';
import 'package:mobile02/protos/generated/project/user.pb.dart';

class UserModel extends ChangeNotifier {
  /// 用户数据
  User user;
  /// 是否已登录
  int isLogined;
  /// 记录登录时间
  String updatedAt;
  factory UserModel.fromUser(User user, int isLogined) {
    return UserModel(
      user: user, 
      isLogined: isLogined, 
      updatedAt: DateTime.now().toLocal().millisecondsSinceEpoch.toString()
    );
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

  /* 更新登录状态 */
  UserModel updateUserLoginState(
    int isLogined,
  ) {
    return _copy(
      isLogined: isLogined,
    );
  }

  /// 拷贝对象
  UserModel _copy({
    String account, 
    String nickname,
    String phone, 
    String email, 
    String photo, 
    int isLogined, 
  }) {
    if (this.user == null) {
      throw 'user cannot be null';
    }
    User u = this.user;
    u.account = account ?? u.account;
    u.nickname = nickname ?? u.nickname;
    u.phone = phone ?? u.phone;
    u.email = email ?? u.email;
    u.photo = photo ?? u.photo;
    return UserModel(
       user: u,
       updatedAt: DateTime.now().toLocal().millisecondsSinceEpoch.toString(), 
       isLogined: isLogined == null ? this.isLogined : isLogined, 
    );
  }


  /// 通知改变登录状态
  void notifyLoginState() {
    notifyListeners();
  }
  
}