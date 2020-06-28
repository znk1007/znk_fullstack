import 'package:cached_network_image/cached_network_image.dart';
import 'package:flb/page/base/item.dart';
import 'package:flb/page/base/launch.dart';
import 'package:flb/page/base/tab.dart';
import 'package:flb/util/http/business/tab.dart';
import 'package:flb/util/http/core/request.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

void main() {
  runApp(
    MultiProvider(
      providers: [
        ChangeNotifierProvider(create: (_) => TabbarItems())
      ],
      child: MyApp(),
    ),
  );
}

class MyApp extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    fetchTabbarItems(context);
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
      home: context.watch<TabbarItems>().items.length > 0 ? TabPage() : LaunchPage(),
    );
  }

  //获取分栏类目数据
  void fetchTabbarItems(BuildContext context) async {
    ResponseResult res = await TabbarItemReq.fetch();
    if (res.statusCode != 0) {
      // Provider.of<TabbarItems>(context, listen: false).add([]);
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
          )
        )
      );
      items.add(item);
    }
    context.read<TabbarItems>().add(items);
    // Provider.of<TabbarItems>(context, listen: false).add(items);
  }
}
