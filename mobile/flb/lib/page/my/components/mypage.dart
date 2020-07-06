import 'package:flb/model/style/mystyle.dart';
import 'package:flb/model/style/style.dart';
import 'package:flb/model/user/user.dart';
import 'package:flb/page/my/components/mylist.dart';
import 'package:flb/page/my/components/myprofile.dart';
import 'package:flb/page/my/model/my.dart';
import 'package:flb/pkg/screen/screen.dart';
import 'package:flb/pkg/table/table.dart';
import 'package:flb/util/random/color.dart';
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
      child: Consumer<ThemeStyle>(
        builder: (context, s, child) {
          return Container(
            height: Screen.screenHeight - s.tabbarHeight,
            child: EasyRefresh(
                controller: _controller,
                child: Consumer3<UserModel, MyModelHandler, MyPageStyle>(
                    builder: (ctx, u, m, ms, c) {
                  List<MyModel> models = m.fetchMyList(u.isLogined);
                  return Container(
                    height: Screen.screenHeight - s.tabbarHeight,
                    child: ZNKTable(
                        scrollable: false,
                        headerSliverBuilder: (context, innerBoxIsScrolled) =>
                            <Widget>[
                              SliverList(
                                  delegate:
                                      SliverChildBuilderDelegate((ctx, index) {
                                return Container(
                                  child: MyProfileView(),
                                );
                              }, childCount: 1)),
                            ],
                        numberOfSection: models.length,
                        numberOfRowsInSection: (section) {
                          List<MyList> lists = models[section].lists;
                          return lists.length;
                        },
                        cellForRowAtIndexPath: (ctx, indexPath) {
                          return Text(
                              'data section ${indexPath.section} row ${indexPath.row}');
                        }),
                  );
                })),
          );
        },
      ),
    );
  }
}

/*
Container(
              child: 
            )
*/
