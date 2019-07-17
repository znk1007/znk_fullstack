
import 'dart:async';
import 'dart:io';
import 'package:flutter/foundation.dart';
import 'package:pedantic/pedantic.dart';
import 'package:sembast/sembast.dart';
import 'package:sembast/sembast_io.dart';
import 'package:path_provider/path_provider.dart' as path;


class SembastDB {

  String _dbName;
  String _storeName;
  SembastDB(String dbName, String storeName) {
    _dbName = dbName;
    if (!_dbName.endsWith('.db')) {
      _dbName = dbName + '.db';
    }
    _storeName = storeName;
    if (!_storeName.startsWith('_')) {
      _storeName = '_' + _storeName;
    }
  }
 
  // 数据库
  Database _db;
  // Database this._db;
  StoreRef _store;
  // 连接数据库
  Future _connectDB() async {
    await this._openDB();
    this._setStore();
  }
  
  // 打开数据库并获取数据库对象
  Future _openDB() async {
    if (this._db != null) {
      return;
    }
    Directory appDocDir = await path.getApplicationDocumentsDirectory();
    var dbFilePath = [appDocDir.path, 'znk', this._dbName].join('/');
    this._db = await databaseFactoryIo.openDatabase(dbFilePath);
  }

  // 关闭数据库
  Future closeDB() async {
    await this._db.close();
    this._db = null;
    
  }

  // 设置Collection
  Future _setStore() async {
    if (this._db == null || this._storeName.startsWith('_') == false || this._store != null) {
      return;
    }
    this._store = StoreRef<String, dynamic>(this._storeName);
  }
  // 统计store总数
  Future<int> count({String field, pattern, bool anyInList = false, String sortFiled, bool ascending = true, int limit = 0}) async {
    await this._connectDB();
    if (this._db == null || this._store == null) {
      return null;
    }
    Future<int> c;
    if (field != null && field != '' && pattern != null && pattern != '') {
      Filter ft = Filter.matches(field, pattern);
      c = this._store.count(this._db, filter: ft); 
    } else {
      c = this._store.count(this._db);
    }
    return c;
  }
 
  // 保存数据
  Future<K> save<K, V>(K recordName, V value) async {
    await this._connectDB();
    if (this._db == null || recordName == '' || recordName == null || this._store == null) {
      return null;
    }
     return await this._store.record(recordName).put(this._db, value, merge: true);
  }

  // 获取指定记录的所有数据
  Future<V> fetch<V>(String recordName) async {
    await this._connectDB();
    if (this._db == null || recordName == '' || recordName == null || this._store == null) {
      return null;
    }
    return await this._store.record(recordName).get(this._db);
  }


  // 查询字段符合正则的数据
  Future<List<RecordSnapshot<String, dynamic>>> findMatches({
    String field, pattern, 
    bool anyInList = false, 
    String sortFiled, 
    bool ascending = true, 
    int limit
  }) async {
    await this._connectDB();
    if (this._db == null || this._store == null) {
      return null;
    }
    SembastFinder fd = SembastFinder();
    if (field != null && field != '' && pattern != null && pattern != '') {
      fd.filter = Filter.matches(field, pattern);
    }
    if (sortFiled != null && sortFiled != '') {
      fd.sortOrder = SortOrder(sortFiled, ascending, true);
    }
    if (limit != null && limit != 0) {
      fd.limit = limit;
    }
    return await this._store.find(this._db,finder: fd);
  }

  // 查找一条符合条件的数据
  Future<RecordSnapshot> findFirstEquals({
    @required String field, 
    @required dynamic value,
    bool anyInList = false, 
  }) async {
    await this._connectDB();
    if (this._db == null || this._store == null) {
      return null;
    }
    SembastFinder fd = SembastFinder();
    if (field != null && field != '' && value != null && value != '') {
      fd.filter = Filter.equals(field, value, anyInList: anyInList);
    }
    return await this._store.findFirst(this._db, finder: fd);
  }


  // 查询字段符合正则的数据
  Future<List<RecordSnapshot<String, dynamic>>> findMatchesRegExp({
    String field, 
    RegExp re, 
    bool anyInList = false, 
    String sortFiled, 
    bool ascending = true, 
    int limit
  }) async {
    await this._connectDB();
    if (this._db == null || this._store == null) {
      return null;
    }
    SembastFinder fd = SembastFinder();
    if (field != null && field != '' && re != null) {
      fd.filter = Filter.matchesRegExp(field, re);
    }
    if (sortFiled != null && sortFiled != '') {
      fd.sortOrder = SortOrder(sortFiled, ascending, true);
    }
    if (limit != null && limit != 0) {
      fd.limit = limit;
    }
    return await this._store.find(this._db,finder: fd);
  }

  // 查询关键值的数据
  Future<K> findKey<K>(
    K key, 
    {
      bool anyInList = false, 
      String sortFiled, 
      bool ascending = true, 
      int limit
    }
  ) async {
    await this._connectDB();
    if (this._db == null || this._store == null) {
      return null;
    }
    SembastFinder fd = SembastFinder();
    fd.filter = Filter.byKey(key);
    if (sortFiled != null && sortFiled != '') {
      fd.sortOrder = SortOrder(sortFiled, ascending, true);
    }
    if (limit != 0) {
      fd.limit = limit;
    }
    return this._store.findKey(this._db,finder: fd);
  }
  // isRecordExists 记录是否存在
  Future<bool> isRecordExists(String recordName) async {
    await this._connectDB();
    if (this._db == null || this._store == null) {
      return false;
    }
    return await this._store.record(recordName).exists(this._db);
  }

  // 更新指定记录内容
  Future<V> updateRecord<K, V> (K recordName, V value) async {
    await this._connectDB();
    if (recordName == null || recordName == '') {
      return null;
    }
    await this._connectDB();
    if (this._db == null || recordName == null ) {
      return null;
    }
    return await this._store.record(recordName).update(this._db, value);
  }
  // 删除指定记录
  Future deleteRecord(String recordName) async {
    if (recordName == null || recordName == '') {
      return null;
    }
    await this._connectDB();
    if (this._db == null) {
      return null;
    }
    return await this._store.record(recordName).delete(this._db);
  }

  // 删除符合正则要求数据，如无设置feild，pattern，将删除所有数据
  Future<int> deleteMatches({String field, pattern}) async {
    if (pattern != '' && pattern != null && field != '' && field != null) {
      return null;
    }
    await this._connectDB();
    if (this._db == null || field == null) {
      return null;
    }
    SembastFinder fd = SembastFinder();
    Filter ft = Filter.matches(field, pattern);
    fd.filter = ft;
    return await this._store.delete(this._db, finder: fd);
  }

  // 查询一条记录
  Future<RecordSnapshot<K, V>> firstRecord<K, V>(
    {
      String field, pattern, 
      bool anyInList = false, 
      String sortFiled, 
      bool ascending = true
    }
  ) async {
  
    await this._connectDB();
    if (this._db == null || this._store == null) {
      return null;
    }
    SembastFinder fd = SembastFinder();
    if (pattern != '' && pattern != null && field != '' && field != null) {
      fd.filter = Filter.matches(field, pattern);
    }
    if (sortFiled != null && sortFiled != '') {
      fd.sortOrder = SortOrder(sortFiled, ascending, true);
    }
    return await this._store.findFirst(this._db, finder: fd);
  }
  // watchRecord 监听指定记录变化
  Future<StreamSubscription<RecordSnapshot<K, V>>> watchRecord<K, V>(
    String recordName, 
    void changes(RecordSnapshot<K, V> snapshot), 
    {
      void done(), 
      Function onErr, 
      bool cancelOnErr
    }
  ) async {
    if (recordName == '') {
      return null;
    }
    await this._connectDB();
    if (this._db == null || this._store == null) {
      return null;
    }
    return this._store.record(recordName).onSnapshot(this._db).listen(changes, onError: onErr, onDone: done, cancelOnError: cancelOnErr);
  }

  // 取消记录订阅
  void cancelRecordSub<K, V>(StreamSubscription<RecordSnapshot<K, V>> sub) {
    unawaited(sub.cancel());
  }
  // 监听指定文档变化
  Future<StreamSubscription<List<RecordSnapshot<K, V>>>> watchStore<K, V>(void changes(List<RecordSnapshot<K, V>> snapshot), 
    {String field, pattern, 
    bool anyInList = false, 
    String sortFiled, 
    bool ascending = true, 
    int limit, 
    void done(), 
    Function onErr, 
    bool cancelOnErr}) async {
    await this._connectDB();
    if (this._db == null || field == null) {
      return null;
    }
    SembastFinder fd = SembastFinder();
    if (pattern != '' && pattern != null && field != '' && field != null) {
      fd.filter = Filter.matches(field, pattern);
    } 
    if (sortFiled != null && sortFiled != '') {
      fd.sortOrder = SortOrder(sortFiled, ascending, true);
    }
    if (limit != null && limit != 0) {
      fd.limit = limit;
    }
    QueryRef<K, V> q = this._store.query(finder: fd);
    return q.onSnapshots(this._db).listen(changes, onError: onErr, onDone: done, cancelOnError: cancelOnErr);
  }

  // 取消store订阅
  void cancelStoreSub<K, V>(StreamSubscription<List<RecordSnapshot<K, V>>> sub) {
    unawaited(sub.cancel());
  }
  
}