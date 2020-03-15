import 'dart:math' as math;

import 'package:flutter/material.dart';

const _preList = [
  "130","131","132","133","134","135","136","137","138","139",
  "147",
  "150","151","152","153","155","156","157","158","159",
  "166",
  "171","176","177",
  "186","187","188",
  "198"
];

const _alphaNum = 'qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890';

const _colors = [
  Colors.black,
  Colors.blue,
  Colors.cyan,
  Colors.orange,
  Colors.red,
  Colors.green,
  Colors.grey,
  Colors.brown,
  Colors.purple,
  Colors.pink,
];

int _min = 0;
int _max = 10;

class RandomManager {
  // 随机生成手机号
  static String randomPhone() {
    String prefix = _preList[math.Random.secure().nextInt(_preList.length)];
    String subfix = '';
    for (var i = 0; i < 8; i++) {
      subfix = subfix + (_min + (math.Random.secure().nextInt(_max - _min))).toString();
    }
    return prefix + subfix;
  }
  /// 随机颜色
  static Color randomColor() {
    return _colors[math.Random.secure().nextInt(_colors.length)];
  }
  
  /* 随机字符串 */
  static String randomString({int len = 18}) {
    String left = '';
    for (var i = 0; i < len; i++) {
      left = left + _alphaNum[math.Random.secure().nextInt(_alphaNum.length)];
    }
    return left;
  }
}