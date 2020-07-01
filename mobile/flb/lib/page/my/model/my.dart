import 'package:flb/util/db/protos/generated/user/user.pb.dart';
import 'package:flutter/widgets.dart';

class MyModel {
  //积分
  String integral;
  //红包
  String bestowal;
  //列表
  List<MyList> lists;
}

class MyEquality extends ChangeNotifier {}

class MyList {
  //唯一标识
  String identifier;
  //icon地址
  String iconPath;
  //标题
  String title;
}

class MyModelHandler extends ChangeNotifier {
  MyEquality _equality;

  Future<void> fetchEqualityData(User user) async {}

  //已登录数据
  MyModel _myLoginedModel;
  //未登录数据
  MyModel _myUnLoginedModel;
  //读取我的页面数据
  MyModel fetchMyList(bool isLogined) {
    if (isLogined) {
      if (_myLoginedModel != null) {}
      return _myLoginedModel;
    } else {
      return _myUnLoginedModel;
    }
  }
}
