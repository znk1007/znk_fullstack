import 'package:flutter/widgets.dart';
import 'package:znk/protos/generated/project/updateonline.pb.dart';
import 'package:znk/protos/generated/project/updateonline.pbgrpc.dart';
import 'package:znk/utils/base/device.dart';
import 'package:znk/utils/security/security.dart' as ssl;
class UpdateOnline {
  String userId;
  String account;
  String sessionId;
  bool online;
  UpdateOnline({@required this.userId, @required this.account, @required this.sessionId, @required this.online}):
    assert(userId != null && userId.isNotEmpty && account != null && account.isNotEmpty && sessionId != null && sessionId.isNotEmpty);
  // 更新在线状态
  Future<UpdateOnlineResponse> update(BuildContext ctx) async {
    final sec = ssl.Security(
      useTls: true,
      certFile: Device.pemPath,  
    );
    final channel = await sec.configurateChannel(ctx);
    final options = sec.configurateCallOptions();
    final client = UpdateOnlineClient(channel);
    final req = UpdateOnlineRequest()
      ..account = this.account
      ..userId = this.userId
      ..device = Device.systemName
      ..sessionId = this.sessionId
      ..online = this.online;
      UpdateOnlineResponse res;
      try {
        res = await client.update(req, options: options);
      } catch (e) {
      }
      await channel.shutdown();
      return res;
  }
}