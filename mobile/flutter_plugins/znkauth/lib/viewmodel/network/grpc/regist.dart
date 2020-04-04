import 'package:flutter/material.dart';
import 'package:grpc/grpc.dart';
import 'package:znkauth/viewmodel/network/grpc/client.dart';

class RegistClient {
  //账号
  String _account;
  //密码
  String _password;
  //初始化
  RegistClient({
    @required String account,
    @required String password,
  }): assert(account.length != 0, password.length != 0) {
    this._account = account;
    this._password = password;
  }
  //run 执行请求
  Future<RegistRes> run() async {
    if (this._account.length == 0 || this._password.length == 0) {
      return null;
    }
    ClientChannel channel = await ZnkAuthRpc.shared.run();
    
    return null;
  }

}