import 'package:flutter/material.dart';
import 'package:znk/modules/tabs/owner/setting/setting/setting_page.dart';

class Routes {
  static Map<String, WidgetBuilder> generate() {
    final routes = Map<String, WidgetBuilder>();
    routes[SettingPage.routeName] = (BuildContext ctx) => SettingPage();
    return routes;
  }
}