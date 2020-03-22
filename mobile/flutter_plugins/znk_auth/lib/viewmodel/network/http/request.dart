
import 'package:dio/dio.dart';

class NetRequest {
  Dio _dio = Dio();
  /// 单例
  NetRequest._();
  static final NetRequest _instance = NetRequest._();
  static NetRequest get shared => _instance;
  factory NetRequest() {
    return _instance;
  }

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
  static Future<void> requestPost(
    String path,
    {
     data,
     Map<String, dynamic> headers,
     Function(bool succ, Map<String, dynamic>) callback 
    }
  ) async {
    Response res = await NetRequest.shared._dio.post(
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