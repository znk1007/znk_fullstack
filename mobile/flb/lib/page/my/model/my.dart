import 'package:flb/model/user/user.dart';
import 'package:flb/util/db/protos/generated/user/user.pb.dart';
import 'package:flutter/widgets.dart';

class MyModel {
  //列表
  List<MyList> lists;
}

class MyEquality extends ChangeNotifier {
  //权益数
  String number;
  //权益标题
  String title;
}

class MyList {
  //唯一标识
  String identifier;
  //icon地址
  String iconPath;
  //标题
  String title;
}

class MyModelHandler extends ChangeNotifier {
  List<MyEquality> _equalitys;

  //权益数据
  List<MyEquality> get equalitys => _equalitys;

  Future<void> fetchEqualityData(UserModel userModel) async {
    List<MyEquality> eqs = [];
    MyEquality eq = MyEquality();
    eq.number = userModel.isLogined ? '88' : '0';
    eq.title = '积分';
    eqs.add(eq);

    eq = MyEquality();
    eq.number = userModel.isLogined ? '88' : '0';
    eq.title = '红包';
    eqs.add(eq);

    _equalitys = eqs;
  }

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
