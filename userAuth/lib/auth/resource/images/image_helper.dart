import 'dart:io';

import 'package:flutter/material.dart';

class ImageHelper {
  /// 加载图片
  static Image load(String src) {
    if (src.startsWith('http://') || 
      src.startsWith('https://') || 
      src.startsWith('ftp://')) {
      return Image.network(src);
    } else if (src.startsWith('file://')) {
      return Image.file(File(src));
    } else if (src.contains('/')) {
      if (Platform.isIOS) {
        return Image.asset('lib/auth/resource/images/iOS/$src');
      }
      return Image.asset('lib/auth/resource/images/android/$src');
    }
    return Image.asset(src);
  }
}