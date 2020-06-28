
import 'package:dio/dio.dart';

class ResponseResult {
  int code;//状态码
  String message;//结果描述
  Map data;//数据
}

class RequestHandler {
  Dio _dio = Dio();
  /// 单例
  RequestHandler._();
  static final RequestHandler _instance = RequestHandler._();
  static RequestHandler get shared => _instance;
  factory RequestHandler() {
    return _instance;
  }

  /// GET 请求
  static Future<ResponseResult> requestGet (
    String path,
    {
      Map<String, dynamic> queryParams,
      Map<String, dynamic> headers,
      Function(bool succ, Map<String, dynamic>) callback,
    }
  ) async {
    Response res = await RequestHandler.shared._dio.get(
      path,
      queryParameters: queryParams,
      options: Options(headers: headers)
    );
    ResponseResult result = ResponseResult();
    result.code = res.statusCode;
    result.message = res.statusMessage;
    result.data = res.data;
    return result;
  }
  /* post 请求 */
  static Future<ResponseResult> requestPost(
    String path,
    {
     data,
     Map<String, dynamic> headers,
     Function(bool succ, Map<String, dynamic>) callback 
    }
  ) async {
    Response res = await RequestHandler.shared._dio.post(
      path,
      data: data,
      options: Options(headers: headers)
    );
    ResponseResult result = ResponseResult();
    result.code = res.statusCode;
    result.message = res.statusMessage;
    result.data = res.data;
    return result;
  }
}