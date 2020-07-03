//TableViewStyle 表格样式
import 'package:flutter/material.dart';

enum TableViewStyle {
  plain, //平铺
  grouped, //分组
}

//IndexPath 地址索引
class IndexPath {
  final int section; //分段
  final int row; //行
  IndexPath({this.section, this.row});
}

class TableView extends StatefulWidget {
  TableView({Key key}) {
    
  }

  @override
  _TableViewState createState() => _TableViewState();
}

class _TableViewState extends State<TableView> {
  @override
  Widget build(BuildContext context) {
    return Container(
       child: child,
    );
  }
}