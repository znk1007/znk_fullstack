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

  bool _loadedItem = false;

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    _fetchTabbarItems(context);
    List<TabbarItem> items = context.watch<TabbarItems>().items;

    return MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: (items.length > 0) ? TabPage(items: items,) : LaunchPage(),
    );
  }

  //获取分栏类目数据
  void _fetchTabbarItems(BuildContext context) async {
    if (_loadedItem) {
      return;
    }
    _loadedItem = true;
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
          )
        )
      );
      items.add(item);
    }
    context.read<TabbarItems>().add(items);
    // Provider.of<TabbarItems>(context, listen: false).add(items);
  }
}
