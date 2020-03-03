
import 'dart:io';

import 'package:flutter/material.dart';
import 'package:sqflite/sqflite.dart';
import 'package:sqflite/sqlite_api.dart';
import 'package:path/path.dart';
import 'package:synchronized/synchronized.dart';

class SqliteDB {
  ///数据库地址
  final String name;

  ///初始化
  SqliteDB(this.name) {
    
  }
  ///数据库句柄
  Database _database;
  /// 锁
  final _lock = new Lock();

  /// 获取数据库对象
  Future<Database> _getDB() async {
    if (_database == null) {
      await _lock.synchronized(() async {
        if (_database == null) {
          String path = await _initDataBasePath();
          _database = await openDatabase(path);
        }
      });
    }
    return _database;
  }

  ///初始化数据库地址
  Future<String> _initDataBasePath() async {
    final path = await _getFilePath(this.name);
    bool isDB =  await _isDatabase(path);
    if (!isDB) {
      await _createFile(path);
    }
    return path;
  }

  ///获取数据库文件路径
  Future<String> _getFilePath(String name) async {
    final databasePath = await getDatabasesPath();
    return join(databasePath, name);
  }

  /// 创建文件
  Future<void> _createFile(String dbPath) async {
    final bool isExists = await Directory(dirname(dbPath)).exists();
    if (isExists) {
      return;
    }
    try {
      await Directory(dbPath).create(recursive: true);
    } catch (_) {
    }
  }

  ///判断是否是数据库路径
  Future<bool> _isDatabase(String path) async {
    Database db;
    bool isDatabase = false;
    try {
      db = await openReadOnlyDatabase(path);
      int version = await db.getVersion();
      if (version != null) {
        isDatabase = true;
      }
    } catch (_) {
    } finally {
      await db?.close();
    }
    return isDatabase;
  }

  Future<void> _onOpen(Database db) async {
    print('db version ${await db.getVersion()}');
  }

  /* 建表 */
  Future<void> createTable(String tableSql) async {
    Database db = await this._getDB();
    await db.execute("CREATE TABLE IF NOT EXISTS ${tableSql}");
  }

  /* 插入数据 */
  Future<int> insert(String table, Map<String, dynamic> values) async {
    Database db = await this._getDB();
    await db.insert(table, values);
  }

  /* 删除数据 */
  Future<int> delete(String table, {String where, List<dynamic> whereArgs}) async {
    Database db = await this._getDB();
    return await db.delete(table, where: where, whereArgs: whereArgs);
  }

  Future<int> update(String table, Map<String, dynamic> values,
      {String where,
      List<dynamic> whereArgs,
      ConflictAlgorithm conflictAlgorithm}) async {
    Database db = await this._getDB();
    return await db.update(table, values, where: where, whereArgs: whereArgs, conflictAlgorithm: conflictAlgorithm);
  }
  /* 查询数据 */
  Future<List<Map<String, dynamic>>> find(String table,
      {bool distinct,
      List<String> columns,
      String where,
      List<dynamic> whereArgs,
      String groupBy,
      String having,
      String orderBy,
      int limit,
      int offset}) async {
    Database db = await this._getDB();
    return await db.query(
      table, 
      distinct: distinct, 
      columns: columns, 
      where: where, 
      whereArgs: whereArgs,
      groupBy: groupBy,
      having: having,
      orderBy: orderBy,
      limit: limit, 
      offset: offset
    );
  }

  /// 关闭数据库
  void close() {
    _database?.close();
    _database = null;
  }
}