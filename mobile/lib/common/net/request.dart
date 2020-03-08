import 'dart:convert';

import 'package:dio/dio.dart';


class NetRequest  {
  /// 请求类
  Dio _dio = Dio();

  /// singleton
  static NetRequest _instance;
  static NetRequest _shared() {
    if (_instance == null) {
      _instance = NetRequest._();
    }
    return _instance;
  }
  NetRequest._();

  /// GET 请求
  static Future<void> requestGet(
    String path, 
    {
      Map<String, dynamic> queryParameters, 
      Map<String, dynamic> headers,
      Function(bool succ, Map<String, dynamic>) callback, 
    }
  ) async {
    Response res = await NetRequest._shared()._dio.get(
      path, 
      queryParameters: queryParameters, 
      options: Options(headers: headers)
    );
    if (res.statusCode == 200) {
      if (res.data is Map) {
        callback ?? callback(true, res.data);
      } else {
        callback ?? callback(false, null);
      }
    } else {
      callback ?? callback(false, null);
    }
  }
  /// post 请求
  static Future<void> requestPost(
    String path, 
    {
      data,
      Map<String, dynamic> headers,
      Function(bool succ, Map<String, dynamic>) callback, 
    }
  ) async {
    Response res = await NetRequest._shared()._dio.post(
      path, 
      data: data, 
      options: Options(headers: headers)
    );
    if (res.statusCode == 200) {
      if (res.data is Map) {
        callback ?? callback(true, res.data);
      } else {
        callback ?? callback(false, null);
      }
    } else {
      callback ?? callback(false, null);
    }
  }

}