import 'package:flb/api/api.dart';
import 'package:flb/models/main/main.dart';
import 'package:flb/util/config/help.dart';
import 'package:flb/util/http/core/request.dart';
import 'package:flb/viewmodels/base.dart';
import 'package:flutter/material.dart';

class ZNKMainViewModel extends ZNKBaseViewModel {
  final ZNKApi api;
  ZNKMainViewModel({@required this.api}) : super(api: api);

  //首页布局配置
  List<ZNKMainModel> _mainModels;
  //是否显示模块
  bool showModule(ZNKMainModule module) {
    if (_mainModels != null) {
      ZNKMainModel model =
          _mainModels.firstWhere((element) => element.module == module);
      return model.show;
    }
    return false;
  }

  //获取首页布局配置
  Future<void> fetchMainLayoutConfig() async {
    if (this.api.mainPageConfigUrl.length == 0) {
      _defaultMainLayoutData();
      return;
    }
    ResponseResult result =
        await RequestHandler.get(this.api.mainPageConfigUrl);
    result.code = -1;
    List<Map<String, dynamic>> layouts = [];
    if (result.statusCode == 200 && result.data != null) {
      String code = result.data['code'];
      result.code = int.parse(code);
      layouts = result.data['layout'];
      if (layouts.length == 0) {
        result.code = -1;
      }
    }
    if (result.code != 0) {
      _defaultMainLayoutData();
      return;
    }
    List<ZNKMainModel> models = [];
    for (var i = 0; i < layouts.length; i++) {
      Map<String, String> layout = layouts[i];
      ZNKMainModule innerModule = ZNKMainModule.search;
      int module = int.parse(ZNKHelp.safeString(layout['module']));
      switch (module) {
        case 1:
          innerModule = ZNKMainModule.search;
          break;
        case 2:
          innerModule = ZNKMainModule.msessage;
          break;
        case 3:
          innerModule = ZNKMainModule.slide;
          break;
        case 4:
          innerModule = ZNKMainModule.nav;
          break;
        case 5:
          innerModule = ZNKMainModule.magic;
          break;
        case 6:
          innerModule = ZNKMainModule.notify;
          break;
        case 7:
          innerModule = ZNKMainModule.seckill;
          break;
        case 8:
          innerModule = ZNKMainModule.ads;
          break;
        case 9:
          innerModule = ZNKMainModule.prod;
          break;
        default:
      }
      bool show = ZNKHelp.safeString(layout['module']) == '1';
      models.add(ZNKMainModel(module: innerModule, show: show));
    }
    notifyListeners();
  }

  void _defaultMainLayoutData() {
    if (_mainModels != null) {
      return;
    }
    _mainModels = [
      ZNKMainModel(module: ZNKMainModule.search, show: true),
      ZNKMainModel(module: ZNKMainModule.msessage, show: true),
      ZNKMainModel(module: ZNKMainModule.slide, show: true),
      ZNKMainModel(module: ZNKMainModule.nav, show: true),
      ZNKMainModel(module: ZNKMainModule.notify, show: false),
      ZNKMainModel(module: ZNKMainModule.seckill, show: true),
      ZNKMainModel(module: ZNKMainModule.magic, show: true),
      ZNKMainModel(module: ZNKMainModule.ads, show: false),
      ZNKMainModel(module: ZNKMainModule.prod, show: true),
    ];
    notifyListeners();
  }
}
