import 'package:flutter/material.dart';
import 'package:znk/utils/base/device.dart';

class CustomColors {
  // navigatorBackColor 导航返回按钮颜色
  static Color get navigatorBackColor {
    return Color.fromARGB(255, 62, 166, 254);
  }
  // 背景颜色
  static Color get backgroundColor {
    return Color.fromARGB(255, 249, 249, 249);
  }
  // 分割线颜色
  static Color get separatorColor {
    return Colors.grey[100];
  }

}

class CustomMeasure {
  // 右箭头大小
  static Size get arrowSize {
    return Size(Device.relativeWidth(15), Device.relativeWidth(15));
  } 
}