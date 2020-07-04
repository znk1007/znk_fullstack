import 'package:flutter/material.dart';
import 'package:tableview/tableview.dart';
void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      title: 'Flutter Demo',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: Scaffold(
        appBar: AppBar(
          title: Text('Grouped List View Example'),
        ),
        body: ZNKTable(
          numberOfSection: 50,
        ),
      ),
    );
  }
}


// void main() {
//   runApp(new MaterialApp(
//     title: 'Container demo',
//     home: new MyScrv(),
//   ));
// }

// class MyScrv extends StatefulWidget {
//   @override
//   State<StatefulWidget> createState() {
//     // TODO: implement createState

//     return new MyScrvState();
//   }
// }

// class MyScrvState extends State<MyScrv> {
//   List<String> _list = new List();
//   List<Color> myColors = new List();

//   @override
//   void initState() {
//     _list.add("政府");
//     _list.add("部门11");
//     _list.add("部门22");
//     myColors.add(Colors.red);
//     myColors.add(Colors.lightBlue);
//     myColors.add(Colors.lightBlue);
//   }

//   @override
//   Widget build(BuildContext context) {
//     // TODO: implement build
//     return new CustomScrollView(physics: ScrollPhysics(), slivers: <Widget>[
//       const SliverAppBar(
//         pinned: true,
//         expandedHeight: 250.0,
//         flexibleSpace: const FlexibleSpaceBar(
//           title: const Text('demo'),
//         ),
//       ),
//       new SliverGrid(
//         gridDelegate: new SliverGridDelegateWithMaxCrossAxisExtent(
//           ///设置item的最大像素宽度  比如 130
//           maxCrossAxisExtent: 150,

//           ///沿主轴的每个子节点之间的逻辑像素数。 默认垂直方向的子child间距  这里的是主轴方向 当你改变 scrollDirection: Axis.vertical,就是改变了主轴发方向
//           mainAxisSpacing: 10.0,

//           ///沿横轴的每个子节点之间的逻辑像素数。默认水平方向的子child间距
//           crossAxisSpacing: 10.0,

//           ///每个孩子的横轴与主轴范围的比率。 child的宽高比  常用来处理child的适配
//           childAspectRatio: 1.0,
//         ),
//         delegate: new SliverChildBuilderDelegate(
//           (BuildContext context, int index) {
//             return new Container(
//               //设置child居中显示
//               alignment: Alignment.center,
//               child: Text(
//                 _list[index],
//                 style: TextStyle(
//                     color: Colors.white,
//                     fontSize: 18.0,
//                     fontWeight: FontWeight.w500,
//                     //去掉黄色下划线
//                     decoration: TextDecoration.none),
//               ),
//               decoration: new BoxDecoration(
//                 color: myColors[index],
//               ),
//             );
//           },
//           childCount: _list.length,
//         ),
//       ),
//       new SliverFixedExtentList(
//         itemExtent: 50.0,
//         delegate: new SliverChildBuilderDelegate(
//           (BuildContext context, int index) {
//             return new Container(
//               alignment: Alignment.center,
//               color: Colors.lightBlue[100 * (index % 9)],
//               child: new Text('list item $index'),
//             );
//           },
//         ),
//       ),
//     ]);
//   }
// }
