import 'package:flb/api/api.dart';
import 'package:flb/models/main/seckill.dart';
import 'package:flb/util/config/help.dart';
import 'package:flb/util/http/core/request.dart';
import 'package:flb/viewmodels/base.dart';
import 'package:flutter/material.dart';

class ZNKSeckillViewModel extends ZNKBaseViewModel {
  final ZNKApi api;
  ZNKSeckillViewModel({@required this.api}) : super(api: api);

  //秒杀
  ZNKSeckill _seckill;
  ZNKSeckill get seckill => _seckill;
  Future<void> fetch() async {
    if (this.api.seckillUrl.length == 0) {
      _defaultSeckillData();
      return;
    }
    ResponseResult result = await RequestHandler.get(this.api.seckillUrl);
    result.code = -1;
    Map<String, dynamic> data;
    List<Map<String, dynamic>> items = [];
    if (result.statusCode == 200 && result.data != null) {
      String code = result.data['code'];
      result.code = int.parse(code);
      data = result.data['data'];
      if (data == null) {
        _defaultSeckillData();
        return;
      }
      items = data['items'];
      if (items.length <= 0) {
        result.code = -1;
      }
    }
    if (result.code != 0) {
      _defaultSeckillData();
      return;
    }
    String title = data['title'] ?? '限时秒杀';
    String time = ZNKHelp.safeString(data['time']);
    List<ZNKSeckillItem> tempItems = [];
    for (var i = 0; i < items.length; i++) {
      Map<String, String> mData = items[i];
      ZNKSeckillItem item = ZNKSeckillItem(
        identifier: ZNKHelp.safeString(mData['id']),
        title: ZNKHelp.safeString(mData['title']),
        orgPrice: ZNKHelp.safeString(mData['orgPrice']),
        newPrice: ZNKHelp.safeString(mData['newPrice']),
      );
      tempItems.add(item);
    }
    _seckill = ZNKSeckill(title: title, time: time, items: tempItems);
    notifyListeners();
  }

  void _defaultSeckillData() {
    _seckill = ZNKSeckill(
      title: '限时秒杀',
      time: '240',
      items: [
        ZNKSeckillItem(
          identifier: '1',
          title: '秒杀标题一',
          orgPrice: '￥418',
          newPrice: '￥189',
        ),
        ZNKSeckillItem(
          identifier: '2',
          title: '秒杀标题二',
          orgPrice: '￥618',
          newPrice: '￥289',
        ),
        ZNKSeckillItem(
          identifier: '3',
          title: '秒杀标题三',
          orgPrice: '￥818',
          newPrice: '￥389',
        ),
      ],
    );
    notifyListeners();
  }
}
