import 'package:cached_network_image/cached_network_image.dart';
import 'package:flb/models/style/style.dart';
import 'package:flb/models/user.dart';
import 'package:flb/provider/provider.dart';
import 'package:flb/views/tabbar/item.dart';
import 'package:flb/views/base/launch.dart';
import 'package:flb/views/tabbar/tabbar.dart';
import 'package:flb/util/db/protos/generated/user/user.pb.dart';
import 'package:flb/util/db/sqlite/sqlitedb.dart';
import 'package:flb/util/db/sqlite/user/user.dart';
import 'package:flb/util/http/tab/tab.dart';
import 'package:flb/util/http/core/request.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'dart:async';

void main() async {
  WidgetsFlutterBinding.ensureInitialized(); //fixed binary message binding
  SqliteDB.shared.setdbName('user.db');
  UserDB.createUserTable();
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    //拉取分栏数据
    _fetchTabbarItems(context);
    return MultiProvider(
      providers: znkProviders,
      child: MaterialApp(
        debugShowCheckedModeBanner: false,
        title: '货满仓',
        theme: ThemeData(
          primarySwatch: Provider.of<ThemeStyle>(context).primarySwatch,
        ),
        home: Consumer<TabbarItems>(
            builder: (ctx1, t, w) => (t.items.length > 0)
                ? FutureBuilder(
                    future: _loadUserData(context),
                    // initialData: InitialData,
                    builder: (BuildContext ctx2, AsyncSnapshot<User> snapshot) {
                      return ZNKTabbar(items: t.items);
                    },
                  )
                : LaunchPage()),
      ),
    );
  }

  //获取分栏类目数据
  Future<void> _fetchTabbarItems(BuildContext context) async {
    ResponseResult res = await TabbarItemReq.fetch();
    if (res.statusCode != 0) {
      context.read<TabbarItems>().add([]);
      return;
    }
    List<Map> body = res.data['body'];
    List<TabbarItem> items = [];
    for (var i = 0; i < body.length; i++) {
      Map<String, dynamic> itemMap = body[i];
      TabbarItem item = TabbarItem(
          identifier: itemMap["identifier"] ? "${itemMap['identifier']}" : "$i",
          index: i,
          item: BottomNavigationBarItem(
              icon: CachedNetworkImage(
            imageUrl: itemMap['icon'],
            placeholder: (context, url) => Icon(Icons.pages),
          )));
      items.add(item);
    }
    context.read<TabbarItems>().add(items);
  }

  //加载用户数据
  Future<User> _loadUserData(BuildContext context) async {
    UserModel userModel = context.read<UserModel>();
    await userModel.loadUserData();
    return userModel.currentUser;
  }
}
