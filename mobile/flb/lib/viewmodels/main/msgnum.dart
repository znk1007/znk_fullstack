import 'package:flb/api/api.dart';
import 'package:flb/util/http/core/request.dart';
import 'package:flb/viewmodels/base.dart';
import 'package:flutter/material.dart';

class ZNKMsgViewModel extends ZNKBaseViewModel {
  final ZNKApi api;
  ZNKMsgViewModel({@required this.api}) : super(api: api);

  //消息数量
  String _msgNum = '';
  String get msgNum => _msgNum;

  //获取消息数量
  Future<void> fetch() async {
    if (this.api.msgNumUrl.length == 0) {
      _msgNum = '3';
      notifyListeners();
      return;
    }
    ResponseResult result = await RequestHandler.get(this.api.tabbarUrl);
    result.code = -1;
    if (result.statusCode == 200 && result.data != null) {
      String code = result.data['code'];
      result.code = int.parse(code);
      String number = result.data['number'];
      if (number != null) {
        _msgNum = number;
      }
    }
    notifyListeners();
  }
}