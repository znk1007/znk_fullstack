import 'package:flutter/widgets.dart';

class MyModel {
  //积分
  String integral;
  //红包
  String bestowal;
  //列表
  List<MyList> lists;
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
  MyModel _myModel;

  void fetchMyData() {
    
  }
}
