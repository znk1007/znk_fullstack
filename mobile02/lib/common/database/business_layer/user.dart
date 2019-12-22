import 'package:flutter/material.dart';
import 'package:mobile02/common/database/base/database.dart';

class UserDB {
  /// 单例
  static UserDB _instance;
  static UserDB get dao {
    if (_instance == null) {
      _instance = UserDB._();
    }
    return _instance;
  }

  DBHelper _helper;

  /*
  companyId INTEGER,
    FOREIGN KEY (companyId) REFERENCES Company(id) 
    ON DELETE CASCADE
  */

  UserDB._(){
    String tableName = 'user';
    _helper = DBHelper(
      tableName: tableName,
      createSql: '''
            create table if not exists 
            $tableName(
              userId text primary key,
              sessionId text,
              account text,
              nickname text,
              phone text,
              email text,
              photo text,
              createdAt text,
              updatedAt text,
              isLogined integer
            ) '''
    );
  }
}