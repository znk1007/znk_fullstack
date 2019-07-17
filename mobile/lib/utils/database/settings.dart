import 'package:znk/utils/database/base/sembastdb.dart';

String _recordPswKey = 'recordPswKey';

class Settings {
  // 单例模块
  static Settings _instance;
  static Settings get dao {
    if (_instance == null) {
      _instance = Settings._();
    }
    return _instance;
  } 
  // 数据库客户端
  SembastDB _db;

  Settings._() {
    _db = SembastDB('settings', '_settings');
  }

  // 设置记住密码
  Future setRecordPsw(bool val) async {
    return await _db.save(_recordPswKey, val);
  }
  // 是否记住密码
  Future<bool> get recordPsw async {
    var val = await _db.fetch(_recordPswKey);
    if (val is bool) {
      return val;
    }
    return false;
  }
  
}