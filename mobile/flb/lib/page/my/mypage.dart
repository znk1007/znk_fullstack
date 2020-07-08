import 'package:flb/model/style/mystyle.dart';
import 'package:flb/model/style/style.dart';
import 'package:flb/model/user/user.dart';
import 'package:flb/page/my/myprofile.dart';
import 'package:flb/page/my/model/my.dart';
import 'package:flb/pkg/screen/screen.dart';
import 'package:flb/pkg/table/table.dart';
import 'package:flutter/material.dart';
import 'package:flutter_easyrefresh/easy_refresh.dart';
import 'package:provider/provider.dart';

class StickyProfileDelegate extends SliverPersistentHeaderDelegate {
  final Widget child;
  final double max;
  StickyProfileDelegate({@required this.child, this.max});
  @override
  Widget build(
          BuildContext context, double shrinkOffset, bool overlapsContent) =>
      this.child;

  @override
  double get maxExtent => this.max;

  @override
  double get minExtent => this.max;

  @override
  bool shouldRebuild(SliverPersistentHeaderDelegate oldDelegate) => true;
}

class MyPage extends StatefulWidget {
  MyPage({Key key}) : super(key: key);

  static String id = 'my';

  final EasyRefreshController _controller = EasyRefreshController();

  @override
  _MyPageState createState() => _MyPageState();
}

class _MyPageState extends State<MyPage> {
  @override
  void dispose() {
    widget._controller.dispose();
    super.dispose();
  }

  @override
  void didUpdateWidget(MyPage oldWidget) {
    super.didUpdateWidget(oldWidget);
  }

  @override
  void didChangeDependencies() {
    print('didChangeDependencies ${this.context}');
    super.didChangeDependencies();
  }

  @override
  Widget build(BuildContext context) {
    return MultiProvider(
      providers: [
        ChangeNotifierProvider(create: (ctx) => MyCompany()),
        ChangeNotifierProvider(create: (ctx) => MyList()),
        ChangeNotifierProvider(create: (ctx) => MyPageStyle()),
      ],
      child: Consumer<ThemeStyle>(
        builder: (context, s, child) {
          return Container(
            height: Screen.screenHeight - s.tabbarHeight,
            child: Consumer5<UserModel, MyList, MyCompany, MyPageStyle,
                ThemeStyle>(
              builder: (context, userModel, list, company, myStyle, themeStyle,
                  child) {
                List<MyList> models = list.fetch(userModel.isLogined);
                return Container(
                  height: Screen.screenHeight - themeStyle.tabbarHeight,
                  child: ZNKTable(
                      scrollable: false,
                      headerSliverBuilder: (context, innerBoxIsScrolled) =>
                          <Widget>[
                            SliverPersistentHeader(
                                pinned: true,
                                delegate: StickyProfileDelegate(
                                    child: MyProfileView(
                                      userModel: userModel,
                                      style: myStyle,
                                      company: company,
                                    ),
                                    max: myStyle.profileBgHeight)),
                          ],
                      numberOfSection: models.length,
                      numberOfRowsInSection: (section) {
                        List<MyItem> lists = models[section].items;
                        return lists.length;
                      },
                      viewForHeaderInSection: (context, section) =>
                          Container(height: 10, color: Colors.grey[200]),
                      cellForRowAtIndexPath: (ctx, indexPath) {
                        List<MyItem> items = models[indexPath.section].items;
                        MyItem item = items[indexPath.row];
                        double iconTop = Screen.setWidth(5);
                        double iconS = myStyle.rowHeight - 2 * iconTop;
                        return GestureDetector(
                            onTap: () {
                              List<MyItem> items =
                                  models[indexPath.section].items;
                              MyItem item = items[indexPath.row];
                              print(
                                  'select index path: $indexPath title: ${item.title}');
                            },
                            child: Row(
                              children: [
                                Container(
                                  margin: EdgeInsets.only(left: 10, top: 0),
                                  width: iconS,
                                  height: iconS,
                                  child: (item.iconPath != null &&
                                          item.iconPath.length > 0)
                                      ? Image.asset(item.iconPath)
                                      : Icon(Icons.add_alert),
                                ),
                                Expanded(
                                  flex: 1,
                                  child: Container(
                                    margin: EdgeInsets.only(left: 8),
                                    child: Text(item.title ?? ''),
                                  ),
                                ),
                                Container(
                                  child: Icon(Icons.keyboard_arrow_right),
                                ),
                              ],
                            ));
                      }),
                );
              },
            ),
          );
        },
      ),
    );
  }
}

/*
Column(children: [
                  Consumer<MyCompany>(builder: (ctx, company, child) {
                    return MyProfileView(
                      userModel: userModel,
                      style: myStyle,
                      company: company,
                    );
                  }),
                  Container(
                      height: Screen.screenHeight -
                          myStyle.profileBgHeight -
                          themeStyle.tabbarHeight,
                      child: EasyRefresh(
                        child: Consumer<MyModel>(
                          builder: (context, m, child) {
                            List<MyModel> models =
                                m.fetchMyList(userModel.isLogined);
                            return Container(
                                height: Screen.screenHeight -
                                    myStyle.profileBgHeight -
                                    themeStyle.tabbarHeight,
                                child: ZNKTable(
                                    numberOfSection: models.length,
                                    numberOfRowsInSection: (section) {
                                      List<MyList> lists =
                                          models[section].lists;
                                      return lists.length;
                                    },
                                    heightForRowAtIndexPath:
                                        (context, indexPath) =>
                                            myStyle.rowHeight,
                                    viewForHeaderInSection:
                                        (context, section) => Container(
                                            height: 10,
                                            color: Colors.grey[300]),
                                    cellForRowAtIndexPath: (ctx, indexPath) {
                                      List<MyList> lists =
                                          models[indexPath.section].lists;
                                      MyList list = lists[indexPath.row];
                                      double iconTop = Screen.setWidth(5);
                                      double iconS =
                                          myStyle.rowHeight - 2 * iconTop;
                                      return GestureDetector(
                                          child: Row(
                                        children: [
                                          Container(
                                            margin: EdgeInsets.only(
                                                left: 10, top: 0),
                                            width: iconS,
                                            height: iconS,
                                            child: (list.iconPath != null &&
                                                    list.iconPath.length > 0)
                                                ? Image.asset(list.iconPath)
                                                : Icon(Icons.add_alert),
                                          ),
                                          Expanded(
                                            flex: 1,
                                            child: Container(
                                              margin: EdgeInsets.only(left: 8),
                                              child: Text(list.title ?? ''),
                                            ),
                                          ),
                                          Container(
                                            child: Icon(
                                                Icons.keyboard_arrow_right),
                                          ),
                                        ],
                                      ));
                                    }));
                          },
                        ),
                      )),
                ])
*/
