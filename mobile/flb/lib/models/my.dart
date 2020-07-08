import 'package:flb/models/url.dart';
import 'package:flb/models/user.dart';
import 'package:flutter/widgets.dart';

enum MyItemType {
  extend, //推广
  cashcoupon, //代金券
  order, //订单
  collection, //收藏
  integ, //积分商城
  addr, //地址
  msg, //消息
  setting, //设置
  about, //关于
}

enum MyListType {
  normal, //普通
  system, //系统
}

class MyItem {
  //icon地址
  String iconPath = '';
  //标题
  String title;
  //列表类型
  MyItemType type;
}

class MyList extends ChangeNotifier {
  //段类型
  MyListType type;
  //列表
  List<MyItem> items = [];

  //已登录数据
  List<MyList> _loginedList = [];
  //未登录数据
  List<MyList> _unloginedList = [];

  //读取我的页面数据
  List<MyList> fetch(bool isLogined) {
    if (isLogined) {
      _unloginedList = [];
      if (_loginedList.length == 0) {
        List<MyList> models = [];
        List<MyItem> temp = [];

        MyList model = MyList();
        model.type = MyListType.normal;

        MyItem sub = MyItem();
        sub.title = '推广';
        sub.type = MyItemType.extend;
        temp.add(sub);

        sub = MyItem();
        sub.title = '代金券';
        sub.type = MyItemType.cashcoupon;
        temp.add(sub);

        sub = MyItem();
        sub.title = '我的订单';
        sub.type = MyItemType.order;
        temp.add(sub);

        sub = MyItem();
        sub.title = '收藏';
        sub.type = MyItemType.collection;
        temp.add(sub);

        sub = MyItem();
        sub.title = '积分商城';
        sub.type = MyItemType.integ;
        temp.add(sub);

        sub = MyItem();
        sub.title = '我的地址';
        sub.type = MyItemType.addr;
        temp.add(sub);

        sub = MyItem();
        sub.title = '我的消息';
        sub.type = MyItemType.extend;
        temp.add(sub);

        model.items = temp;
        models.add(model);

        model = MyList();
        temp = [];
        sub = MyItem();
        sub.title = '关于我们';
        sub.type = MyItemType.about;
        temp.add(sub);

        sub = MyItem();
        sub.title = '设置';
        sub.type = MyItemType.setting;
        temp.add(sub);

        model.items = temp;
        models.add(model);
        _loginedList = models;
      }
      return _loginedList;
    } else {
      _loginedList = [];
      if (_unloginedList.length == 0) {
        List<MyList> models = [];
        MyList model = MyList();
        List<MyItem> temp = [];
        MyItem sub = MyItem();
        sub.title = '关于我们';
        sub.type = MyItemType.about;
        temp.add(sub);
        sub = MyItem();
        sub.title = '设置';
        sub.type = MyItemType.setting;
        temp.add(sub);
        model.items = temp;
        models.add(model);
        _unloginedList = models;
      }
      return _unloginedList;
    }
  }
}

class MyCompany extends ChangeNotifier {
  //企业名
  String name;
  //企业编码
  String code;
  //积分
  String integ;
  //红包
  String redPack;

  //企业信息
  MyCompany _innerInfo;
  //企业信息
  MyCompany get info => _innerInfo;
  //拉取企业信息
  Future<void> fetchCompanyInfo(UserModel userModel) async {
    String sessionID = userModel.currentUser.sessionID;
    String compInfoUrl = CompanyInfoURL.compInfo;
    if (userModel.isLogined && sessionID.length > 0 && compInfoUrl.length > 0) {
      return;
    }
    MyCompany comp = MyCompany();
    comp.name = 'xxxx有限公司';
    comp.code = '10000';
    comp.integ = '0';
    comp.redPack = '0';
    _innerInfo = comp;
  }
}
