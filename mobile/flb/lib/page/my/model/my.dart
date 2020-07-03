import 'package:flb/model/user/user.dart';
import 'package:flb/util/db/protos/generated/user/user.pb.dart';
import 'package:flutter/widgets.dart';

enum MyListType {
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

enum MyModelType {
  normal, //普通
  system, //系统
}

class MyModel {
  //段类型
  MyModelType type;
  //列表
  List<MyList> lists;
}

class MyList {
  //icon地址
  String iconPath = '';
  //标题
  String title;
  //列表类型
  MyListType type;
}

class MyCompany {
  //企业名
  String name;
  //企业编码
  String code;
  //积分
  String integ;
  //红包
  String redPack;
}

class MyModelHandler extends ChangeNotifier {
  //企业信息
  MyCompany _company;

  MyCompany get company => _company;
  //拉取企业信息
  Future<void> fetchCompanyInfo(UserModel userModel) async {
    String sessionID = userModel.currentUser.sessionID;
    if (userModel.isLogined && sessionID.length > 0) {
      return
    }
  }

  /* 拉取收益数据 */
  Future<void> fetchEqualityData(UserModel userModel) async {
    List<MyEquality> eqs = [];
    MyEquality eq = MyEquality();
    eq.number = userModel.isLogined ? '88' : '0';
    eq.title = '积分';
    eq.offsetRadio = 0;
    eq.widthRadio = 1 / 2.0;
    eqs.add(eq);

    eq = MyEquality();
    eq.number = userModel.isLogined ? '88' : '0';
    eq.title = '红包';
    eq.offsetRadio = 1 / 2.0;
    eq.widthRadio = 1 / 2.0;
    eqs.add(eq);

    _equalitys = eqs;
  }

  //已登录数据
  List<MyModel> _loginedList;
  //未登录数据
  List<MyModel> _unloginedList;

  //读取我的页面数据
  List<MyModel> fetchMyList(bool isLogined) {
    if (isLogined) {
      _unloginedList = [];
      if (_loginedList.length == 0) {
        List<MyModel> models = [];
        MyModel model = MyModel();
        model.type = MyModelType.normal;
        List<MyList> temp = [];
        MyList sub = MyList();
        sub.title = '推广';
        sub.type = MyListType.extend;
        temp.add(sub);
        sub = MyList();
        sub.title = '代金券';
        sub.type = MyListType.cashcoupon;
        temp.add(sub);
        sub = MyList();
        sub.title = '我的订单';
        sub.type = MyListType.order;
        temp.add(sub);
        sub = MyList();
        sub.title = '收藏';
        sub.type = MyListType.collection;
        temp.add(sub);
        sub = MyList();
        sub.title = '积分商城';
        sub.type = MyListType.integ;
        temp.add(sub);
        sub = MyList();
        sub.title = '我的地址';
        sub.type = MyListType.addr;
        temp.add(sub);
        sub = MyList();
        sub.title = '我的消息';
        sub.type = MyListType.extend;
        temp.add(sub);
        sub = MyList();
        model.lists = temp;
        models.add(model);

        model = MyModel();
        temp = [];
        sub = MyList();
        sub.title = '关于我们';
        sub.type = MyListType.about;
        temp.add(sub);
        sub = MyList();
        sub.title = '设置';
        sub.type = MyListType.setting;
        temp.add(sub);
        models.add(model);
        _loginedList = models;
      }
      return _loginedList;
    } else {
      _loginedList = [];
      if (_unloginedList.length == 0) {
        List<MyModel> models = [];
        MyModel model = MyModel();
        List<MyList> temp = [];
        MyList sub = MyList();
        sub.title = '关于我们';
        sub.type = MyListType.about;
        temp.add(sub);
        sub = MyList();
        sub.title = '设置';
        sub.type = MyListType.setting;
        temp.add(sub);
        models.add(model);
        _unloginedList = models;
      }
      return _unloginedList;
    }
  }
}
