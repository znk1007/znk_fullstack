import 'package:flutter/material.dart';

class ImageHelper {
  /// 获取图片资源
  static Image getImage(String src, {bool isPNG = true}) {
    if (src.startsWith('http://') || src.startsWith('https://')) {
      return Image.network(src);
    }
    bool hasSubfix = src.toLowerCase().endsWith('jpg') || src.toLowerCase().endsWith('jpeg') || src.toLowerCase().endsWith('png');
    if (hasSubfix) {
      return Image.asset('lib/resource/images/iOS/$src');
    } else {
      if (isPNG) {
        return Image.asset('lib/resource/images/iOS/$src.png');
      } else {
        return Image.asset('lib/resource/images/iOS/$src.jpg');
      }
    }
  }
}