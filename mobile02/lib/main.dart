import 'dart:io';

import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:mobile02/common/database/business_layer/user.dart';
import 'package:mobile02/utils/helper/random.dart';
import '3rd/plugins/device/device_helper.dart';
import '3rd/plugins/device/path_helper.dart' as path;
import '3rd/state_manager/src/consumer.dart';
import 'common/database/base/database.dart';
import 'model/user/user.dart';
import 'protos/generated/project/user.pb.dart';

void main() async {
  runApp(MyApp());
  await DBHelper.initDeleteDBFile(dbName: 'znk.db');
}

class MyApp extends StatelessWidget {

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
        // This is the theme of your application.
        //
        // Try running your application with "flutter run". You'll see the
        // application has a blue toolbar. Then, without quitting the app, try
        // changing the primarySwatch below to Colors.green and then invoke
        // "hot reload" (press "r" in the console where you ran "flutter run",
        // or simply save your changes to "hot reload" in a Flutter IDE).
        // Notice that the counter didn't reset back to zero; the application
        // is not restarted.
        primarySwatch: Colors.blue,
      ),
      home: MyHomePage(title: 'Flutter Demo Home Page'),
    );
  }
}

class MyHomePage extends StatefulWidget {
  MyHomePage({Key key, this.title}) : super(key: key);

  // This widget is the home page of your application. It is stateful, meaning
  // that it has a State object (defined below) that contains fields that affect
  // how it looks.

  // This class is the configuration for the state. It holds the values (in this
  // case the title) provided by the parent (in this case the App widget) and
  // used by the build method of the State. Fields in a Widget subclass are
  // always marked "final".

  final String title;

  @override
  _MyHomePageState createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {

  static final DeviceHelper deviceInfoPlugin = DeviceHelper();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Column(
        children: <Widget>[
          RaisedButton(
            child: Text('模拟登录'),
            onPressed: () {
              String userId = RandomManager.randomString();
              String account = '用户' + RandomManager.randomString(len: 3);
              var user = User.create()
                ..account = account
                ..userId = userId;
              UserDB.insert(user, 0);
            },
          ),
          Consumer<UserModel>(
          builder: (context, model, _) {
            if (model.isLogined == 1) {
              return Text('已登录');
            } else {
              return Text('退出登录');
            }
          }
        )
        ],
      ),
      floatingActionButton: FloatingActionButton(
        child: Icon(Icons.add),
        onPressed: () {

        },
      ),
    );
  }
}
