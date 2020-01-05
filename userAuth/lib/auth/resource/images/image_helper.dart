import 'dart:io';

import 'package:flutter/material.dart';

class ImageHelper {
  /// 加载图片
  static Image load(String src, {double scale, double width, double height}) {
    if (src.startsWith('http://') || 
      src.startsWith('https://') || 
      src.startsWith('ftp://')) {
      return Image.network(src, scale: scale, width: width, height: height);
    } else if (src.startsWith('file://')) {
      return Image.file(File(src), scale: scale, width: width, height: height);
    } else if (src.contains('/') == false) {
      if (Platform.isIOS) {
        return Image.asset('lib/auth/resource/images/iOS/$src', scale: scale, width: width, height: height);
      }
      return Image.asset('lib/auth/resource/images/android/$src', scale: scale, width: width, height: height);
    }
    return Image.asset(src, scale: scale, width: width, height: height);
  }
}