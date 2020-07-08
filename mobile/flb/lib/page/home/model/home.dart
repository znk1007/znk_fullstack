import 'package:flutter/material.dart';

class ZNKBannerModel extends ChangeNotifier {
  List<String> recommends = [];//推荐列表
  bool test = true;
  //获取推荐数据
  Future <void> fetchRecommends() async {
    if (test) {
      recommends = ['防水地板1','防水地板2','防水地板2','防水地板2','防水地板2','防水地板6',];
    }
  }

}