import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:znk_auth/controller/auth.dart';
import 'package:znk_auth/model/delegate/auth.dart';
import 'package:znk_auth/model/protos/generated/auth/user.pb.dart';

export 'model/delegate/auth.dart';
export 'model/protos/generated/auth/user.pb.dart';

class ZnkAuth {
  /* 配置是否OK */
  static bool _isOK = false;

  static bool _isTest = true;

  static ZnkAuthConfig _config;
  /* 通道 */
  static const MethodChannel _channel =
      const MethodChannel('znk_auth');
  /* 获取平台信息 */
  static Future<String> get platformVersion async {
    final String version = await _channel.invokeMethod('getPlatformVersion');
    return version;
  }
  /* 基本配置 */
  static void configuration(ZnkAuthConfig config) {
    _config = config;
    _isOK = true;
  }
  /* 路由名称 */
  static String get znkRouteName => '/znkauth';
  /* 路由 */
  static Map<String, Widget Function(BuildContext)> znkRoute(BuildContext context) => {znkRouteName: (context) => AuthPage(config: _config)};
  /* show 显示验证页面 */
  static void show(BuildContext context, Callback callback) {
    if (!_isOK || _config == null) {
      if (callback != null) {
        callback(false, '缺少配置参数');
      }
      if (!_isTest) {
        return;
      }
    }
    callback(true, '缺少配置参数');
    Navigator.push(
      context, 
      MaterialPageRoute(
        builder: (context) => AuthPage(config: _config),
      ),
    );
  }
  /* 退出登录 */
  static void logout(User user) {
    
  }
}
