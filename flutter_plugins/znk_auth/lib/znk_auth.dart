import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:znk_auth/controller/auth.dart';
import 'package:znk_auth/model/delegate/auth.dart';

import 'model/delegate/url.dart';

export 'model/delegate/auth.dart';

class ZnkAuth {
  /* 配置是否OK */
  static bool _isOK = false;

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
  static Map<String, Widget Function(BuildContext)> znkRoute(BuildContext context) => {znkRouteName: (context) => AuthPage()};
  /* push 到验证页面 */
  static void push(BuildContext context) {
    Navigator.push(
      context, 
      MaterialPageRoute(
        builder: (context) => AuthPage(),
      ),
    );
  }
}
