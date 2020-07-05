import 'package:flb/model/style/mystyle.dart';
import 'package:flb/model/user/user.dart';
import 'package:flb/page/my/components/mylist.dart';
import 'package:flb/page/my/components/myprofile.dart';
import 'package:flb/page/my/model/my.dart';
import 'package:flb/pkg/table/table.dart';
import 'package:flb/util/random/color.dart';
import 'package:flb/util/screen/screen.dart';
import 'package:flutter/material.dart';
import 'package:flutter_easyrefresh/easy_refresh.dart';
import 'package:provider/provider.dart';

class MyPage extends StatelessWidget {
  static String id = 'my';

  final EasyRefreshController _controller = EasyRefreshController();

  MyPage({Key key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return MultiProvider(
      providers: [
        ChangeNotifierProvider(create: (ctx) => MyModelHandler()),
        ChangeNotifierProvider(create: (ctx) => MyPageStyle()),
      ],
      child: Container(
        height: Screen.screenHeight - 48,
        child: EasyRefresh(
          controller: _controller,
          child: Consumer2<UserModel, MyModelHandler>(builder: (ctx, u, m, c) {
            List<MyModel> models = m.fetchMyList(u.isLogined);
            return Container(
              height: Screen.screenHeight - 48,
              child: ZNKTable(
                numberOfSection: models.length,
                numberOfRowsInSection: (section) {
                  List<MyList> lists = models[section].lists;
                  return lists.length;
                },
                cellForRowAtIndexPath: (ctx, indexPath){
                  return Text('data section ${indexPath.section} row ${indexPath.row}');
                }),
            );
          })
      ),
      ),
    );
  }
}

/*
Container(
              child: 
            )
*/