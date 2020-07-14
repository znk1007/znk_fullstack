import 'package:flb/api/api.dart';
import 'package:flb/viewmodels/base.dart';
import 'package:flutter/material.dart';

class ZNKLoginViewModel extends ZNKBaseViewModel {
  final ZNKApi api;
  ZNKLoginViewModel({@required this.api}) : super(api: api);
  //登录
  Future<void> login({
    @required String account,
    String password,
    String verifyCode,
  }) async {
    if (this.api.loginUrl.length <= 0) {
      return;
    }
  }
}
