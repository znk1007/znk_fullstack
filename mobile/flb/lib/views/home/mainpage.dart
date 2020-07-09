import 'package:flb/models/style/style.dart';
import 'package:flb/models/main.dart';
import 'package:flb/pkg/banner/banner.dart';
import 'package:flb/pkg/screen/screen.dart';
import 'package:flb/pkg/search/search.dart';
import 'package:flb/viewmodels/main.dart';
import 'package:flb/views/base/base.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class ZNKMainPage extends StatelessWidget {
  static const String id = 'home';
  const ZNKMainPage({Key key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Stack(
      children: [
        ZNKBaseView<ZNKMainRecommand>(
          model: ZNKMainRecommand(api: Provider.of(context)),
          onReady: (model) async {
            model.fetchRecommand();
          },
          builder: (context, model, child) {
            double height = 40.0;
            return model.recommends.length > 0
                ? ZNKSearchView(
                    style: ZNKSearchStyle(
                      enabled: false,
                      backgroudColor: Colors.grey,
                      cornerRadius: height / 2.0,
                      margin: EdgeInsets.only(left: 14, top: ZNKScreen.safeTopArea),
                      width: ZNKScreen.screenWidth - 80,
                      height: height,
                    ),
                    child: ZNKBanner(
                      banners: model.recommends
                          .map((e) => Container(child: Text(e)))
                          .toList(),
                      showIndicator: false,
                      scrollDirection: Axis.vertical,
                      alignment: Alignment.centerLeft,
                      didSelected: (index) {
                        print('did selected: $index');
                      },
                    ),
                  )
                : Container();
          },
        ),
      ],
    );
  }
}

// class MainPage extends StatefulWidget {
//   static String id = 'home';

//   MainPage({Key key}) : super(key: key);

//   @override
//   _MainPageState createState() => _MainPageState();
// }

// class _MainPageState extends State<MainPage> {

//   @override
//   void initState() {
//     super.initState();
//   }

//   @override
//   void didChangeDependencies() {
//     super.didChangeDependencies();
//     print('didChangeDependencies');
//   }

//   @override
//   void didUpdateWidget(MainPage oldWidget) {
//     super.didUpdateWidget(oldWidget);
//     print('didUpdateWidget');
//   }

//   @override
//   Widget build(BuildContext context) {
//     return MultiProvider(
//       providers: [ChangeNotifierProvider(create: (_) => ZNKBannerModel())],
//       child: Consumer<ThemeStyle>(builder: (ctx, style, child) {
//         return Container(
//           child: Stack(
//             children: [
//               Container(
//                 child: Row(
//                   children: [
//                     Container(
//                       child: Consumer<ZNKBannerModel>(
//                           builder: (ctx, bannerModel, child) {
//                         return bannerModel.recommends.length > 0 ? ZNKBanner(
//                           banners: bannerModel.recommends
//                               .map((e) => Container(child: Text(e)))
//                               .toList(),
//                           margin: EdgeInsets.only(left: 50, top: 80),
//                           scrollDirection: Axis.horizontal,
//                           alignment: Alignment.centerLeft,
//                           didSelected: (index) {
//                             print('did selected: $index');
//                           },
//                         ) : Container();
//                       }),
//                     ),
//                   ],
//                 ),
//               ),
//             ],
//           ),
//         );
//       }),
//     );
//   }
// }
