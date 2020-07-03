import 'dart:js';

import 'package:cached_network_image/cached_network_image.dart';
import 'package:flb/model/style/style.dart';
import 'package:flb/model/user/user.dart';
import 'package:flb/page/base/item.dart';
import 'package:flb/page/base/launch.dart';
import 'package:flb/page/base/tab.dart';
import 'package:flb/util/db/protos/generated/user/user.pb.dart';
import 'package:flb/util/db/sqlite/sqlitedb.dart';
import 'package:flb/util/db/sqlite/user/user.dart';
import 'package:flb/util/http/tab/tab.dart';
import 'package:flb/util/http/core/request.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized(); //fixed binary message binding
  SqliteDB.shared.setdbName('user.db');
  UserDB.createUserTable();
  runApp(
    MultiProvider(
      providers: [
        ChangeNotifierProvider(create: (_) => TabbarItems()),
        ChangeNotifierProvider(create: (_) => UserModel()),
        ChangeNotifierProvider(create: (_) => ThemeStyle()),
      ],
      child: MyApp(),
    ),
  );
}

class MyApp extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    //拉取分栏数据
    _fetchTabbarItems(context);

    return MaterialApp(
      debugShowCheckedModeBanner: false,
      title: '货满仓',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: Consumer<TabbarItems>(
          builder: (ctx, t, w) =>
              (t.items.length > 0) ? TabPage() : LaunchPage()),
    );
  }

  //获取分栏类目数据
  void _fetchTabbarItems(BuildContext context) async {
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
}
