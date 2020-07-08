import 'package:flb/util/config/url/url.dart';
import 'package:flb/util/http/core/request.dart';

class TabbarItemReq {
  //fetch 获取分栏类目
  static Future<ResponseResult> fetch() async {
    ResponseResult result = await RequestHandler.get(URLString.tabbarItem);
    result.code = -1;
    if (result.statusCode == 200 && result.data != null) {
      String code = result.data['code'];
      result.code = int.parse(code);
      List<Map> body = result.data['body'];
      if (body.length == 0) {
        result.code = -1;
      }
    }
    return result;
  }
}