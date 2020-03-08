
import 'package:dio/dio.dart';

class NetRequest {
  Dio _dio = Dio();
  /// 单例
  static NetRequest _instance;
  static get shared {
    if (_instance == null) {
      _instance = NetRequest._();
    }
    return _instance;
  }
  NetRequest._();

  /// GET 请求
  static Future<void> requestGet (
    String path,
    {
      Map<String, dynamic> queryParams,
      Map<String, dynamic> headers,
      Function(bool succ, Map<String, dynamic>) callback,
    }
  ) async {
    Response res = await NetRequest.shared._dio.get(
      path,
      queryParameters: queryParams,
      options: Options(headers: headers)
    );
    if (res.statusCode != 200) {
      callback ?? callback(false, null);
      return;
    }
    if (res.data is Map) {
      callback ?? callback(true, res.data);
    } else {
      callback ?? callback(false, null);
    }
  }
  /* post 请求 */
  Future<void> requestPost(
    String path,
    {
     data,
     Map<String, dynamic> headers,
     Function(bool succ, Map<String, dynamic>) callback 
    }
  ) async {
    Response res = NetRequest.shared._dio.post(
      path,
      data: data,
      options: Options(headers: headers)
    );
    if (res.statusCode != 200) {
      callback ?? callback(false, null);
      return;
    }
    if (res.data is Map) {
      callback ?? callback(true, res.data);
    } else {
      callback ?? callback(false, null);
    }
  }
}