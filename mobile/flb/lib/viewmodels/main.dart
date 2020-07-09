import 'package:flb/api/api.dart';
import 'package:flb/viewmodels/base.dart';
import 'package:flutter/material.dart';

class ZNKMainViewModel extends ZNKBaseViewModel {
  final ZNKApi api;
  ZNKMainViewModel({@required this.api}):super(api:api);
  //获取推荐数据
  Future<void> fetchRecommand() async {

  }
}