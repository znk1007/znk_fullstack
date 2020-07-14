import 'package:shared_preferences/shared_preferences.dart';

class ZNKCache {
  static Future<SharedPreferences> _prefs = SharedPreferences.getInstance();
  //设置bool值
  static Future<void> setBool(String key, bool val) async {
    final SharedPreferences prefs = await _prefs;
    prefs.setBool(key, val);
  }

  //获取布尔值
  static Future<bool> getBool(key) async {
    final SharedPreferences prefs = await _prefs;
    return prefs.getBool(key) ?? false;
  }

  //设置字符串
  static Future<void> setValue(String key, String val) async {
    final SharedPreferences prefs = await _prefs;
    prefs.setString(key, val);
  }

  //获取字符串值
  static Future<String> getValue(String key) async {
    final SharedPreferences prefs = await _prefs;
    return prefs.getString(key) ?? '';
  }

  //移除指定key数据
  static Future<bool> remove(String key) async {
    final SharedPreferences prefs = await _prefs;
    return prefs.remove(key);
  }
}
