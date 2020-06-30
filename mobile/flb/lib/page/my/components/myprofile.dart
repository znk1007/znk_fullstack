import 'package:flb/model/user/user.dart';
import 'package:flb/util/db/protos/generated/user/user.pbgrpc.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class MyProfile extends StatelessWidget {
  const MyProfile({Key key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      child: FutureBuilder(
        future: _fetchUserData(context),
        builder: (BuildContext context, AsyncSnapshot<User> snapshot) {
          print("user: $snapshot");
          return Container(
            child: Text('我的'),
          );
        },
      ),
    );
  }

  /* 请求用户数据 */
  Future<User> _fetchUserData(BuildContext context) async {
    return await context.read<UserModel>().current;
  }
}
