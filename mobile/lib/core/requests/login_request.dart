import 'package:flutter/material.dart';
import 'package:znk/protos/generated/project/login.pbgrpc.dart';
import 'package:znk/utils/base/device.dart';
import 'package:znk/utils/security/security.dart' as ssl;
class Login {
  final String account;
  final String password;
  final String userId;
  Login({@required this.account, @required this.userId, @required this.password}):
  assert(account != null && account.isNotEmpty && userId != null && userId.isNotEmpty && password != null && password.isNotEmpty);

  Future<LoginResponse> login(BuildContext ctx) async {
    var psw = ssl.Security.aesEncode(this.password);
    final sec = ssl.Security(
      useTls: true,
      certFile: Device.pemPath,
    );
    final channel = await sec.configurateChannel(ctx);
    final callOptions = sec.configurateCallOptions();
    final client = LoginClient(channel);
    final req = LoginRequest()
      ..account = this.account
      ..userId = this.userId
      ..password = psw
      ..device = Device.systemName;
    
    LoginResponse res;
    try {
      res = await client.login(req, options: callOptions);
    } catch (e) {
      print('login err: $e');
    }
    await channel.shutdown();
    return res;
  }
}