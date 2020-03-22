import 'package:znk_auth/model/protos/generated/auth/user.pb.dart';

import 'url.dart';
abstract class ZnkAuthConfig extends ZnkUrlConfig {
  /* 请求地址域名 */
  String get baseUrl;
  /* 用户信息，登录状态回调 */
  Function(User user, bool offline) authenticate;
}