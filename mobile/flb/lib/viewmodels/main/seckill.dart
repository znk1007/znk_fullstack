import 'dart:async';

import 'package:flb/api/api.dart';
import 'package:flb/models/main/seckill.dart';
import 'package:flb/util/config/help.dart';
import 'package:flb/util/http/core/request.dart';
import 'package:flb/viewmodels/base.dart';
import 'package:flutter/material.dart';

class ZNKSeckillCountDownViewModel extends ZNKBaseViewModel {
  final ZNKApi api;
  ZNKSeckillCountDownViewModel({@required this.api}) : super(api: api);
  //天
  String _day = '0';
  String get day => _day;
  //小时
  String _hour = '00';
  String get hour => _hour;
  //分
  String _minute = '00';
  String get minute => _minute;
  //秒
  String _second = '00';
  String get second => _second;

  Timer _timer;

  Future<void> fetch(int seconds) async {
    if (_timer != null) {
      _timer.cancel();
      _timer = null;
    }
    _timer = Timer.periodic(Duration(seconds: 1), (timer) {
      int day = seconds ~/ (60 * 60 * 24); //换成天
      int hh = (seconds - (60 * 60 * 24 * day)) ~/
          3600; //总秒数-换算成天的秒数=剩余的秒数    剩余的秒数换算为小时
      int mm = (seconds - 60 * 60 * 24 * day - 3600 * hh) ~/
          60; //总秒数-换算成天的秒数-换算成小时的秒数=剩余的秒数    剩余的秒数换算为分
      int ss = seconds -
          60 * 60 * 24 * day -
          3600 * hh -
          60 * mm; //总秒数-换算成天的秒数-换算成小时的秒数-换算为分的秒数=剩余的秒数
      if (hh != 0) {
        _hour = hh >= 10 ? '$hh' : '0$hh';
      }
      if (mm != 0) {
        _minute = mm >= 10 ? '$mm' : '0$mm';
      }
      if (ss != 0) {
        _second = ss >= 10 ? '$ss' : '0$ss';
      }
      seconds--;
      if (seconds <= 0) {
        _timer.cancel();
        _timer = null;
      }
      notifyListeners();
    });
  }
}

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
        path: ZNKHelp.safeString(mData['path']),
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
          path: 'lib/resource/collection.jpg',
        ),
        ZNKSeckillItem(
          identifier: '2',
          title: '秒杀标题二',
          orgPrice: '￥618',
          newPrice: '￥289',
          path: 'lib/resource/collection.jpg',
        ),
        ZNKSeckillItem(
          identifier: '3',
          title: '秒杀标题三',
          orgPrice: '￥818',
          newPrice: '￥389',
          path: 'lib/resource/collection.jpg',
        ),
      ],
    );
    notifyListeners();
  }
}
