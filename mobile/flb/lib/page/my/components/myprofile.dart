import 'package:flb/model/user/user.dart';
import 'package:flb/util/db/protos/generated/user/user.pbgrpc.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class MyProfile extends StatelessWidget {
  const MyProfile({Key key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    context.read<UserModel>().current;
    return Container(
      child: Text('data')
    );
  }

  /* 请求用户数据 */
  Future<User> _fetchUserData(BuildContext context) async {
    return await context.read<UserModel>().current;
  }
}