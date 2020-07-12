import 'package:flb/api/api.dart';
import 'package:flb/models/main/prod.dart';
import 'package:flb/util/config/help.dart';
import 'package:flb/util/http/core/request.dart';
import 'package:flb/viewmodels/base.dart';
import 'package:flutter/material.dart';

class ZNKProdListViewModel extends ZNKBaseViewModel {
  final ZNKApi api;
  ZNKProdListViewModel({@required this.api}) : super(api: api);
  //产品列表集合
  List<ZNKProd> _prods = [];
  List<ZNKProd> get prods => _prods;

  Future<void> fetch() async {
    if (this.api.prodListUrl.length == 0) {
      _defaultProdData();
      return;
    }
    ResponseResult result = await RequestHandler.get(this.api.prodListUrl);
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
      _defaultProdData();
      return;
    }
    List<ZNKProd> prods = [];
    for (var i = 0; i < data.length; i++) {
      Map<String, dynamic> mainMap = data[i];
      List<ZNKProdItem> items = [];
      String mainTitle = ZNKHelp.safeString(mainMap['title']);
      List<Map<String, dynamic>> itemsMap = mainMap['items'] ?? [];
      for (var j = 0; j < itemsMap.length; j++) {
        Map<String, dynamic> itemMap = itemsMap[j];
        String path = ZNKHelp.safeString(itemMap['path']);
        String subTitle = ZNKHelp.safeString(itemMap['title']);
        String detail = ZNKHelp.safeString(itemMap['detail']);
        List<String> tags = itemMap['tags'] ?? [];
        String coinType = itemMap['coinType'] ?? '￥';
        String unit = ZNKHelp.safeString(itemMap['unit']);
        String orgPrice = ZNKHelp.safeString(itemMap['orgPrice']);
        String newPrice = ZNKHelp.safeString(itemMap['newPrice']);
        String solt = ZNKHelp.safeString(itemMap['solt']);
        String stock = ZNKHelp.safeString(itemMap['stock']);

        List<ZNKProdItem> subItems = [];
        List<Map<String, dynamic>> subItemsMap = itemMap['items'] ?? [];
        for (var j = 0; j < itemsMap.length; j++) {
          Map<String, dynamic> itemMap = itemsMap[j];
          String path = ZNKHelp.safeString(itemMap['path']);
          String subTitle = ZNKHelp.safeString(itemMap['title']);
          String detail = ZNKHelp.safeString(itemMap['detail']);
          List<String> tags = itemMap['tags'] ?? [];
          String coinType = itemMap['coinType'] ?? '￥';
          String unit = ZNKHelp.safeString(itemMap['unit']);
          String orgPrice = ZNKHelp.safeString(itemMap['orgPrice']);
          String newPrice = ZNKHelp.safeString(itemMap['newPrice']);
          String solt = ZNKHelp.safeString(itemMap['solt']);
          String stock = ZNKHelp.safeString(itemMap['stock']);
          ZNKProdItem subProdItem = ZNKProdItem(
            title: subTitle,
            path: path,
            detail: detail,
            tags: tags,
            coinType: coinType,
            unit: unit,
            orgPrice: orgPrice,
            newPrice: newPrice,
            solt: solt,
            stock: stock,
          );
          subItems.add(subProdItem);
        }
        ZNKProdItem prodItem = ZNKProdItem(
          title: subTitle,
          path: path,
          detail: detail,
          tags: tags,
          coinType: coinType,
          unit: unit,
          orgPrice: orgPrice,
          newPrice: newPrice,
          solt: solt,
          stock: stock,
          prods: subItems,
        );
        items.add(prodItem);
      }
      prods.add(
        ZNKProd(
          title: mainTitle,
          items: items,
        ),
      );
    }
    notifyListeners();
  }

  //默认产品数据
  void _defaultProdData() {
    if (_prods.length > 0) {
      return;
    }
    _prods = [
      ZNKProd(
        title: '热卖产品',
        items: [
          ZNKProdItem(
            title: 'PVC复合型木塑防水运动地板01',
            detail: '采用悬浮式安装，无缝设计，无色差，锁扣安装。免漆工艺，国内首次01',
            tags: ['强化复合', '防水'],
            unit: '平方',
            coinType: '￥',
            orgPrice: '90.00',
            newPrice: '60.00',
            solt: '34287',
            stock: '320000',
          ),
          ZNKProdItem(
            title: 'WPC-1523',
            detail: '采用悬浮式安装，无缝设计，无色差，锁扣安装。免漆工艺，国内首次01',
            tags: ['强化复合', '防水'],
            unit: '平方',
            coinType: '￥',
            orgPrice: '90.00',
            newPrice: '60.00',
            solt: '34287',
            stock: '320000',
          ),
          ZNKProdItem(
            title: '特价好货',
            detail: '',
            tags: [],
            unit: '',
            coinType: '',
            orgPrice: '',
            newPrice: '',
            solt: '',
            stock: '',
            prods: [
              ZNKProdItem(
                title: 'WPC-1681',
                detail: '采用悬浮式安装，无缝设计，无色差，锁扣安装。免漆工艺，国内首次01',
                tags: ['强化复合', '防水'],
                unit: '平方',
                coinType: '￥',
                orgPrice: '90.00',
                newPrice: '60.00',
                solt: '34287',
                stock: '320000',
              ),
              ZNKProdItem(
                title: 'WPC-1681',
                detail: '采用悬浮式安装，无缝设计，无色差，锁扣安装。免漆工艺，国内首次01',
                tags: ['强化复合', '防水'],
                unit: '平方',
                coinType: '￥',
                orgPrice: '90.00',
                newPrice: '60.00',
                solt: '34287',
                stock: '320000',
              ),
              ZNKProdItem(
                title: 'WPC-1681',
                detail: '采用悬浮式安装，无缝设计，无色差，锁扣安装。免漆工艺，国内首次01',
                tags: ['强化复合', '防水'],
                unit: '平方',
                coinType: '￥',
                orgPrice: '90.00',
                newPrice: '60.00',
                solt: '34287',
                stock: '320000',
              ),
            ],
          ),
          ZNKProdItem(
            title: 'WPC-1525',
            detail: '采用悬浮式安装，无缝设计，无色差，锁扣安装。免漆工艺，国内首次01',
            tags: ['强化复合', '防水'],
            unit: '平方',
            coinType: '￥',
            orgPrice: '90.00',
            newPrice: '60.00',
            solt: '34287',
            stock: '320000',
          ),
          ZNKProdItem(
            title: 'WPC-1528',
            detail: '采用悬浮式安装，无缝设计，无色差，锁扣安装。免漆工艺，国内首次01',
            tags: ['强化复合', '防水'],
            unit: '平方',
            coinType: '￥',
            orgPrice: '90.00',
            newPrice: '60.00',
            solt: '34287',
            stock: '320000',
          ),
        ],
      ),
    ];
    notifyListeners();
  }
}

/*
//大标题
  final String title;
  //详情
  final String detail;
  //图片路径
  final String path;
  //标签集合
  final List<String> tags;
  //原价
  final String orgPrice;
  //折扣价
  final String newPrice;
  //已卖
  final String solt;
  //库存
  final String stock;
  //商品列表
  final List<ZNKProd> prods;
*/
