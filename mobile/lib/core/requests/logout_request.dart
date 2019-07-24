import 'package:flutter/widgets.dart';
import 'package:znk/protos/generated/project/logout.pb.dart';
import 'package:znk/protos/generated/project/logout.pbgrpc.dart';
import 'package:znk/utils/base/device.dart';
import 'package:znk/utils/security/security.dart' as ssl;
class Logout {
  final String userId;
  final String sessionId;
  Logout({
    @required this.userId,
    @required this.sessionId
  }): assert(userId != null && userId.isNotEmpty && sessionId != null && sessionId.isNotEmpty);
  Future<LogoutResponse> logout(BuildContext ctx) async {
    final sec = ssl.Security(
      useTls: true,
      certFile: Device.pemPath,
    );
    final channel = await sec.configurateChannel(ctx);
    final ops = sec.configurateCallOptions();
    final client = LogoutServiceClient(channel);
    LogoutResponse res;
    try {
      res = await client.logout(
        LogoutRequest()
          ..userId = this.userId
          ..sessionId = this.sessionId
          ..device = Device.systemName,
        options: ops,
      );
    } catch (e) {
    }
    await channel.shutdown();
    return res;
  }
}