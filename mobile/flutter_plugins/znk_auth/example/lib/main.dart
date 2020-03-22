import 'package:flutter/material.dart';
import 'dart:async';

import 'package:flutter/services.dart';
import 'package:znk_auth/znk_auth.dart';
import 'package:znk_auth_example/home.dart';

void main() => runApp(MyApp());

class MyApp extends StatefulWidget {
  @override
  _MyAppState createState() => _MyAppState();
}

class _MyAppState extends State<MyApp> {
  

  @override
  void initState() {
    super.initState();
  }

  

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      // routes: ,
      home: HomePage(),
      routes: ZnkAuth.znkRoute(context),
    );
  }
}

