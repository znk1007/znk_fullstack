import 'package:cached_network_image/cached_network_image.dart';
import 'package:flb/models/style/style.dart';
import 'package:flb/models/user.dart';
import 'package:flb/provider/provider.dart';
import 'package:flb/router/paths.dart';
import 'package:flb/router/router.dart';
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
  runApp(MultiProvider(
    providers: znkProviders,
    child: MyApp(),
  ));
}

class MyApp extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    //拉取分栏数据
    // _fetchTabbarItems(context);
    ThemeStyle style = context.watch<ThemeStyle>();
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      title: '货满仓',
      theme: ThemeData(
        primarySwatch: style.primarySwatch,
        visualDensity: style.visualDensity,
      ),
      initialRoute: ZNKRoutePaths.tabbar,
      onGenerateRoute: ZNKRouter.generateRoute,
    );
  }

  

  //加载用户数据
  Future<User> _loadUserData(BuildContext context) async {
    UserModel userModel = context.read<UserModel>();
    await userModel.loadUserData();
    return userModel.currentUser;
  }
}
