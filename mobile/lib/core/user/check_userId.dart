
import 'package:flutter/material.dart';
import 'package:znk/protos/generated/project/checkuserId.pbgrpc.dart';
import 'package:znk/utils/base/device.dart';
import 'package:znk/utils/security/security.dart' as ssl;
class CheckUserId {
  final String account;
  CheckUserId({@required this.account}):
    assert(account.isNotEmpty);

  Future<CheckUserIdResponse> check(BuildContext ctx) async {
    final sec = ssl.Security(
      useTls: true,
      certFile: Device.pemPath,
    );
    final channel = await sec.configurateChannel(ctx);
    final callOptions = sec.configurateCallOptions();
    final client = CheckUserIdClient(channel);
    final req = CheckUserIdRequest()
      ..account = this.account
      ..device = Device.systemName;
    CheckUserIdResponse res;
    try {
      res = await client.check(req, options: callOptions);
    } catch (e) {
      print('check userId err: $e');
    }
    await channel.shutdown();
    return res;
  }
}