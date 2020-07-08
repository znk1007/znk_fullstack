import 'dart:math' as math;

import 'package:flutter/material.dart';

int _min = 0;
  int _max = 10;
  List _prelist = [
    "130",
    "131",
    "132",
    "133",
    "134",
    "135",
    "136",
    "137",
    "138",
    "139",
    "147",
    "150",
    "151",
    "152",
    "153",
    "155",
    "156",
    "157",
    "158",
    "159",
    "186",
    "187",
    "188",
    "176",
    "177",
  ];

class RandomHandler {
  //随机色1
  static Color color({
    int r = 255,
    int g = 255,
    int b = 255,
    int a = 255,
  }) {
    if (r == 0 || g == 0 || b == 0) {
      return Colors.black;
    }
    if (r == 0) {
      return Colors.white;
    }
    return Color.fromARGB(
      a,
      r != 255 ? r : math.Random.secure().nextInt(r),
      g != 255 ? g : math.Random.secure().nextInt(g),
      b != 255 ? b : math.Random.secure().nextInt(b),
    );
  }

  //随机色2
  static get randomColor =>
      Color((math.Random().nextDouble() * 0xFFFFFF).toInt() << 0)
          .withOpacity(1.0);
  //随机数
  static int randomNum({int max = 99}) {
    return math.Random.secure().nextInt(max);
  }
  //随机手机号
  static String randomPhone() {
    String left = _prelist[math.Random.secure().nextInt(_prelist.length)];
    String right = '';
    for (var i = 0; i < 8; i++) {
      right = right + (_min + (math.Random.secure().nextInt(_max - _min))).toString();
    }
    return left + right;
  }
}
