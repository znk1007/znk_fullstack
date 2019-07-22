import 'package:flutter/material.dart';
import 'package:znk/protos/generated/project/register.pbgrpc.dart';
import 'package:znk/utils/base/device.dart';
import 'package:znk/utils/database/user.dart';
import 'package:znk/utils/security/security.dart' as ssl;
/// 注册类
class Register {
  String _account;
  String _password;
  Register({String account, String password}) {
    this._account = account;
    this._password = ssl.Security.aesEncode(password);
  }

  ///注册方法
  Future<RegistResponse> regist(BuildContext ctx) async {
    final sec = ssl.Security(useTls: true, certFile: Device.pemPath);
    final channel = await sec.configurateChannel(ctx);
    final callOptions = sec.configurateCallOptions();
    final client = RegisterClient(channel);
    final req = RegistRequest()
      ..account = this._account
      ..password = this._password
      ..device = Device.systemName;
    RegistResponse res;
    try {
      res = await client.regist(req, options: callOptions);
    } catch (e) {
      print('regist err: $e');
    }
    await channel.shutdown();
    return res;
  }
}