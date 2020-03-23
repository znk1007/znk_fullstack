import 'dart:async';

import 'package:flutter/services.dart';

class Znkauth {
  static const MethodChannel _channel =
      const MethodChannel('znkauth');

  static Future<String> get platformVersion async {
    final String version = await _channel.invokeMethod('getPlatformVersion');
    return version;
  }
}
