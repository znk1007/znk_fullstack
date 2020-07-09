import 'package:flb/api/api.dart';
import 'package:flb/util/http/core/request.dart';
import 'package:flb/viewmodels/base.dart';
import 'package:flutter/material.dart';

class ZNKMainRecommand extends ZNKBaseViewModel {
  final ZNKApi api;
  ZNKMainRecommand({@required this.api}) : super(api: api);
  //推荐数据
  List<String> _recommends = [];
  List<String> get recommends => _recommends;

  //获取推荐数据
  Future<void> fetch() async {
    if (this.api.recommandUrl.length == 0) {
      _recommends = ['防水地板', '集成墙板', '墙布墙漆', '家居软装', '吊顶天花', '五金配件'];
      notifyListeners();
      return;
    }
    ResponseResult result = await RequestHandler.get(this.api.tabbarUrl);
    result.code = -1;
    List<String> data = [];
    if (result.statusCode == 200 && result.data != null) {
      String code = result.data['code'];
      result.code = int.parse(code);
      data = result.data['data'];
      if (data.length == 0) {
        result.code = -1;
      }
    }
    if (result.statusCode != 0) {
      _recommends = ['防水地板', '集成墙板', '墙布墙漆', '家居软装', '吊顶天花', '五金配件'];
      notifyListeners();
      return;
    }
    _recommends = data;
    notifyListeners();
  }
}