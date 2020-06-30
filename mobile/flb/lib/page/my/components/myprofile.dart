import 'package:flb/model/user/user.dart';
import 'package:flb/util/db/protos/generated/user/user.pbgrpc.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class MyProfile extends StatelessWidget {
  const MyProfile({Key key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    bool isLogined = context.watch<UserModel>().isLogined;
    return Container(
      child: isLogined ? _loginedWidget(context) : _unLoginedWidget(context),
    );
  }

  Widget _loginedWidget(BuildContext context) {}

  Widget _unLoginedWidget(BuildContext context) {}
}
