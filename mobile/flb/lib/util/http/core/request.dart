
import 'package:dio/dio.dart';

class ResponseResult {
  int statusCode;//状态码
  int code;//结果状态码
  String message;//结果描述
  Map data;//数据
}

class RequestHandler {
  Dio _dio = Dio();
  /// 单例
  RequestHandler._();
  static final RequestHandler _instance = RequestHandler._();
  factory RequestHandler() {
    return _instance;
  }

  /// GET 请求
  static Future<ResponseResult> get(
    String path,
    {
      Map<String, dynamic> queryParams,
      Map<String, dynamic> headers,
      Function(bool succ, Map<String, dynamic>) callback,
    }
  ) async {
    if (path.length == 0) {
      ResponseResult result = ResponseResult();
      result.statusCode = 404;
      result.message = '网络异常';
      result.data = null;
      return result;
    }
    try {
      Response res = await RequestHandler._instance._dio.get(
        path,
        queryParameters: queryParams,
        options: Options(headers: headers)
      );
      ResponseResult result = ResponseResult();
      result.statusCode = res.statusCode;
      result.message = res.statusMessage;
      result.data = res.data;
      return result;
    } catch (e) {
      ResponseResult result = ResponseResult();
      result.statusCode = 404;
      result.message = '网络异常';
      result.data = null;
      return result;
    }
  }

  /* post 请求 */
  static Future<ResponseResult> post(
    String path,
    {
     data,
     Map<String, dynamic> headers,
     Function(bool succ, Map<String, dynamic>) callback 
    }
  ) async {
    try {
      Response res = await RequestHandler._instance._dio.post(
        path,
        data: data,
        options: Options(headers: headers)
      );
      ResponseResult result = ResponseResult();
      result.statusCode = res.statusCode;
      result.message = res.statusMessage;
      result.data = res.data;
      return result;
    } catch (e) {
      ResponseResult result = ResponseResult();
      result.statusCode = 404;
      result.message = '网络异常';
      result.data = null;
      return result;
    }
  }
}