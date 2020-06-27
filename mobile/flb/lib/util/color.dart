import 'dart:math' as math;

import 'package:flutter/material.dart';

class RandomHandler {
  //随机色1
  static Color color({int r = 255,int g = 255,int b = 255,int a = 255,}) {
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
  static get randomColor => Color((math.Random().nextDouble() * 0xFFFFFF).toInt()<<0).withOpacity(1.0);
}