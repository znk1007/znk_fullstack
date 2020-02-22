
import 'package:flutter/material.dart';

enum ThemeType {
  dark,
  light
}

class ThemeModel with ChangeNotifier {
  ThemeData themeData;
  ThemeType curType;

  ThemeModel(ThemeType type) {
    curType = type;
    if (type == ThemeType.dark) {
      themeData = ThemeData.dark();
    } else {
      themeData = ThemeData.light();
    }
  }

  void reverse() {
    if (curType == ThemeType.dark) {
      themeData = ThemeData.light();
      curType = ThemeType.light;
    } else {
      themeData = ThemeData.dark();
      curType = ThemeType.dark;
    }
    notifyListeners();
  }
}