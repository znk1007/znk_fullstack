import 'package:flb/api/api.dart';
import 'package:flb/models/main/combine.dart';
import 'package:flb/util/config/help.dart';
import 'package:flb/util/http/core/request.dart';
import 'package:flb/viewmodels/base.dart';
import 'package:flutter/material.dart';

class ZNKCombineViewModel extends ZNKBaseViewModel {
  final ZNKApi api;
  ZNKCombineViewModel({@required this.api}) : super(api: api);
  //火拼爆品
  ZNKCombine _combine;
  ZNKCombine get combine => _combine;
  Future<void> fetchCombine() async {
    if (this.api.combineUrl.length == 0) {
      _defaultCombineData();
      return;
    }
    ResponseResult result = await RequestHandler.get(this.api.combineUrl);
    result.code = -1;
    Map<String, dynamic> data;
    List<Map<String, dynamic>> items = [];
    if (result.statusCode == 200 && result.data != null) {
      String code = result.data['code'];
      result.code = int.parse(code);
      data = result.data['data'];
      if (data == null) {
        _defaultCombineData();
        return;
      }
      items = data['items'];
      if (items.length <= 0) {
        result.code = -1;
      }
    }
    if (result.code != 0) {
      _defaultCombineData();
      return;
    }
    String title = data['title'] ?? '火拼爆品';
    List<ZNKCombineItem> tempItems = [];
    for (var i = 0; i < items.length; i++) {
      Map<String, String> mData = items[i];
      ZNKCombineItem item = ZNKCombineItem(
        identifier: ZNKHelp.safeString(mData['id']),
        title: ZNKHelp.safeString(mData['title']),
        price: ZNKHelp.safeString(mData['price']),
      );
      tempItems.add(item);
    }
    _combine = ZNKCombine(title: title, items: tempItems);
    notifyListeners();
  }

  void _defaultCombineData() {
    _combine = ZNKCombine(
      title: '火拼爆品',
      items: [
        ZNKCombineItem(
          identifier: '1',
          title: '爆品标题一',
          price: '￥189',
        ),
        ZNKCombineItem(
          identifier: '2',
          title: '爆品标题二',
          price: '￥289',
        ),
        ZNKCombineItem(
          identifier: '3',
          title: '爆品标题三',
          price: '￥389',
        ),
      ],
    );
    notifyListeners();
  }
}
