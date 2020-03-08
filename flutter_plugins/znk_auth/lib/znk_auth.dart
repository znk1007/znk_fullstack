import 'dart:async';

import 'package:flutter/services.dart';

class ZnkAuth {
  static const MethodChannel _channel =
      const MethodChannel('znk_auth');

  static Future<String> get platformVersion async {
    final String version = await _channel.invokeMethod('getPlatformVersion');
    return version;
  }
}