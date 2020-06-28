import 'package:flb/util/config/url/url.dart';
import 'package:flb/util/http/core/request.dart';

class TabbarItemReq {
  //fetch 获取分栏类目
  static Future<void> fetch() async {
    ResponseResult res = await RequestHandler.get(URLString.tabbarItem);
    if (res.code != 200) {
      
    }
  }
}