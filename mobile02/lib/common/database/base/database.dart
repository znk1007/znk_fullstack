import 'dart:async';
import 'dart:io';

import 'package:flutter/material.dart';
import 'package:sqflite/sqflite.dart';
import 'package:path/path.dart';

/// 数据库路径
String _dbPath;

class DBHelper {

  /// 创建数据库文件或删除数据库
  static Future initDeleteDBFile({String dbName, bool clean = false}) async {
    final String databasePath = await getDatabasesPath();
    _dbPath = join(databasePath, dbName);
    print('db path: $_dbPath');
    if (await Directory(dirname(_dbPath)).exists() && clean) {
      await deleteDatabase(_dbPath);
    } else {
      try {
        await Directory(dirname(_dbPath)).create(recursive: true);
      } catch (e) {
        print('create db path e: $e');
      }
    }
  }

  /// 数据库对象
  Database _db;
  /// 表名
  String _tableName;
  /// 是否清理数据库
  bool _clean = false;
  /// 数据库版本
  int _version = 1;
  /// 建表语句
  String _createSql = '';
  /// 配置数据库语句
  String _configureSql = '';
  /// 升级数据库语句
  String _upgrageSql = '';
  /// 降级数据库语句
  String _downgradeSql = '';

  String _openSql = '';
  /// 初始化
  DBHelper({
    @required String tableName,
    String configureSql, 
    String createSql = '',
    String upgradeSql = '',
    String downgradeSql = '',
    String openSql = '', 
    int version = 1, 
    bool clean = false,
  }) {
    _clean = clean;
    _tableName = tableName;
    _configureSql = configureSql;
    _createSql = createSql;
    _upgrageSql = upgradeSql;
    _downgradeSql = downgradeSql;
    _openSql = openSql;
  }
  
  /// 打开数据库
  Future<bool> _openDB() async {
    if (this._db != null || this._createSql.isEmpty) {
      return false;
    }
    if (_dbPath.isEmpty) {
      await DBHelper.initDeleteDBFile(dbName: 'znk');
    }
    this._db = await openDatabase(
      _dbPath, 
      version: _version, 
      onConfigure: _configureSql.isEmpty ? null : _onConfigure, 
      onCreate: _createSql.isEmpty ? null : _onCreate, 
      onUpgrade: _upgrageSql.isEmpty ? null : _onUpgrade, 
      onDowngrade: _downgradeSql.isEmpty ? null : _onDowngrade,
      onOpen: _openSql.isEmpty ? null : _onOpen,
    );
    return true;
  }

  /// 配置数据库语句
  FutureOr<void> _onConfigure(Database db) async {
    db.execute(_configureSql);
  }

  /// 建表语句
  FutureOr<void> _onCreate(Database db, int version) async {
    db.execute(_createSql);
  }

  /// 升级数据库语句
  FutureOr<void> _onUpgrade(Database db, int oldVersion, int newVersion) async {
    db.execute(_upgrageSql);
  }

  /// 降级数据库语句
  FutureOr<void> _onDowngrade(Database db, int oldVersion, int newVersion) async {
    db.execute(_downgradeSql);
  }

  /// 打开数据库语句
  FutureOr<void> _onOpen(Database db) async {
    db.execute(_openSql);
  }

  /// 关闭数据库
  Future _closeDB() async {
    if (this._db == null) {
      return;
    }
    await this._db.close();
    this._db = null;
  }

  /// 插入表数据
  Future insert(Map<String, dynamic> values) async {
    try {
      if (_tableName.isEmpty) {
        return;
      }
      await this._openDB();
      await this._db.insert(_tableName, values, conflictAlgorithm: ConflictAlgorithm.rollback);
    } catch (e) {
      print('insert err: $e');
    } finally {
      await this._closeDB();
    }
  }
  
  /// 删除表数据
  Future delete({String where, List<dynamic> whereArgs}) async {
    try {
      if (_tableName.isEmpty) {
        return;
      }
      await this._openDB();
      await this._db.delete(_tableName, where: where, whereArgs: whereArgs);
    } catch (e) {
      print('delete err: $e');
    } finally {
      await this._closeDB();
    }
  } 

  /// 更新表数据
  Future update(Map<String, dynamic> values) async {
    try {
      if (_tableName.isEmpty) {
        return;
      }
      await this._openDB();
      await this._db.update(_tableName, values);
    } catch (e) {
      print('delete err: $e');
    } finally {
      await this._closeDB();
    }
  }

  /// 查询数据
  Future<List<Map<String, dynamic>>> query(
      {bool distinct,
      List<String> columns,
      String where,
      List<dynamic> whereArgs,
      String groupBy,
      String having,
      String orderBy,
      int limit,
      int offset}) async {
        try {
          if (_tableName.isEmpty) {
            return null;
          }
          await this._openDB();
          return await this._db.query(
            _tableName, 
            distinct: distinct, 
            columns: columns, 
            where: where, 
            whereArgs: whereArgs, 
            groupBy: groupBy, 
            having: having, 
            orderBy: orderBy, 
            limit: limit, 
            offset: offset,
          );
        } catch (e) {
          print('delete err: $e');
          return null;
        } finally {
          await this._closeDB();
        }
    }
  
}