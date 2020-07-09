import 'package:flb/router/paths.dart';
import 'package:flb/views/auth/login.dart';
import 'package:flb/views/base/launch.dart';
import 'package:flb/views/home.dart';
import 'package:flb/views/tabbar/tabbar.dart';
import 'package:flutter/material.dart';

class ZNKRouter {
  //配置路由
  static Route<dynamic> generateRoute(RouteSettings settings) {
    switch (settings.name) {
      case ZNKRoutePaths.login:
        return MaterialPageRoute(builder: (_) => ZNKLoginView());
        break;
      case ZNKRoutePaths.home:
      return MaterialPageRoute(builder: (_) => ZNKHomePage());
        break;
      default:
    }
  }
}
