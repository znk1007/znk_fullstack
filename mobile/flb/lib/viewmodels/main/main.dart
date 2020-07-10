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
      ZNKMainModel model = _mainModels.firstWhere((element) => element.module == module);
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
          innerModule = ZNKMainModule.notify;
          break;
        case 6:
          innerModule = ZNKMainModule.seckill;
          break;
        case 7:
          innerModule = ZNKMainModule.magic;
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
      ZNKMainModel(module: ZNKMainModule.magic, show: false),
      ZNKMainModel(module: ZNKMainModule.ads, show: false),
      ZNKMainModel(module: ZNKMainModule.prod, show: true),
    ];
    notifyListeners();
  }

  //广告数据
  List<ZNKBannerModel> _banners = [];
  List<ZNKBannerModel> get banners => _banners;

  //拉取广告数据
  Future<void> fetchBanner() async {
    if (this.api.msgNumUrl.length == 0) {
      _defaultBannerData();
      return;
    }
    ResponseResult result = await RequestHandler.get(this.api.bannerUrl);
    result.code = -1;
    List<Map> data = [];
    if (result.statusCode == 200 && result.data != null) {
      String code = result.data['code'];
      result.code = int.parse(code);
      data = result.data['data'];
      if (data.length == 0) {
        result.code = -1;
      }
    }
    if (result.statusCode != 0) {
      _defaultBannerData();
      return;
    }
    List<ZNKBannerModel> temp = [];
    for (var i = 0; i < data.length; i++) {
      Map<String, dynamic> dataMap = data[i];
      String id = ZNKHelp.safeString(dataMap['id']);
      String path = ZNKHelp.safeString(dataMap['path']);
      if (id.length > 0 && path.length > 0) {
        ZNKBannerModel m = ZNKBannerModel(identifier: id, path: path);
        temp.add(m);
      }
    }
    _banners = temp;
    notifyListeners();
  }

  //默认数据
  void _defaultBannerData() {
    _banners = [
      ZNKBannerModel(identifier: '1', path: 'lib/resource/test_img_01.jpg'),
      ZNKBannerModel(identifier: '2', path: 'lib/resource/test_img_02.jpg'),
      ZNKBannerModel(identifier: '3', path: 'lib/resource/test_img_03.jpg'),
    ];
    notifyListeners();
  }

  //导航栏数据
  List<ZNKCollection> _navs = [];
  List<ZNKCollection> get navs => _navs;

  //拉取导航栏数据
  Future<void> fetchNav() async {
    if (this.api.navUrl.length == 0) {
      _defaultNavData();
      return;
    }
    ResponseResult result = await RequestHandler.get(this.api.navUrl);
    result.code = -1;
    List<Map> data = [];
    if (result.statusCode == 200 && result.data != null) {
      String code = result.data['code'];
      result.code = int.parse(code);
      data = result.data['data'];
      if (data.length == 0) {
        result.code = -1;
      }
    }
    if (result.statusCode != 0) {
      _defaultNavData();
      return;
    }
    List<ZNKCollection> temp = [];
    for (var i = 0; i < data.length; i++) {
      Map<String, dynamic> dataMap = data[i];
      String id = ZNKHelp.safeString(dataMap['id']);
      String path = ZNKHelp.safeString(dataMap['path']);
      String title = ZNKHelp.safeString(dataMap['title']);
      if (id.length > 0 && path.length > 0 && title.length > 0) {
        ZNKCollection m =
            ZNKCollection(identifier: id, path: path, title: title);
        temp.add(m);
      }
    }
    _navs = temp;
  }

  //默认数据
  void _defaultNavData() {
    if (_navs.length > 0) {
      return;
    }
    _navs = [
      ZNKCollection(
          identifier: '1', path: 'lib/resource/collection.jpg', title: '标题一'),
      ZNKCollection(
          identifier: '2', path: 'lib/resource/collection.jpg', title: '标题二'),
      ZNKCollection(
          identifier: '3', path: 'lib/resource/collection.jpg', title: '标题三'),
      ZNKCollection(
          identifier: '4', path: 'lib/resource/collection.jpg', title: '标题四'),
      ZNKCollection(
          identifier: '5', path: 'lib/resource/collection.jpg', title: '标题五'),
      ZNKCollection(
          identifier: '6', path: 'lib/resource/collection.jpg', title: '标题六'),
      ZNKCollection(
          identifier: '7', path: 'lib/resource/collection.jpg', title: '标题七'),
      ZNKCollection(
          identifier: '8', path: 'lib/resource/collection.jpg', title: '标题八'),
      ZNKCollection(
          identifier: '9', path: 'lib/resource/collection.jpg', title: '标题九'),
      ZNKCollection(
          identifier: '10', path: 'lib/resource/collection.jpg', title: '标题十'),
      ZNKCollection(
          identifier: '11', path: 'lib/resource/collection.jpg', title: '标题十一'),
      ZNKCollection(
          identifier: '12', path: 'lib/resource/collection.jpg', title: '标题十二'),
      ZNKCollection(
          identifier: '13', path: 'lib/resource/collection.jpg', title: '标题十三'),
      ZNKCollection(
          identifier: '14', path: 'lib/resource/collection.jpg', title: '标题十四'),
      ZNKCollection(
          identifier: '15', path: 'lib/resource/collection.jpg', title: '标题十五'),
      ZNKCollection(
          identifier: '16', path: 'lib/resource/collection.jpg', title: '标题十六'),
      ZNKCollection(
          identifier: '17', path: 'lib/resource/collection.jpg', title: '标题十七'),
      ZNKCollection(
          identifier: '18', path: 'lib/resource/collection.jpg', title: '标题十八'),
      ZNKCollection(
          identifier: '19', path: 'lib/resource/collection.jpg', title: '标题十九'),
      ZNKCollection(
          identifier: '20', path: 'lib/resource/collection.jpg', title: '标题二十'),
    ];
    notifyListeners();
  }

  //消息数量
  String _msgNum = '';
  String get msgNum => _msgNum;

  //获取消息数量
  Future<void> fetchMsgNum() async {
    if (this.api.msgNumUrl.length == 0) {
      _msgNum = '3';
      notifyListeners();
      return;
    }
    ResponseResult result = await RequestHandler.get(this.api.msgNumUrl);
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

  //推荐数据
  List<String> _recommends = [];
  List<String> get recommends => _recommends;

  //获取推荐数据
  Future<void> fetchRecommends() async {
    if (this.api.recommandUrl.length == 0) {
      _recommends = ['防水地板', '集成墙板', '墙布墙漆', '家居软装', '吊顶天花', '五金配件'];
      notifyListeners();
      return;
    }
    ResponseResult result = await RequestHandler.get(this.api.recommandUrl);
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
