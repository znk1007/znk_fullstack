import 'package:flutter/material.dart';

class SearchStyle {
  
}

class SearchView extends StatefulWidget {
  //搜索样式
  final SearchStyle style;
  SearchView({Key key,this.style}) : super(key: key);

  @override
  _SearchViewState createState() => _SearchViewState();
}

class _SearchViewState extends State<SearchView> {
  @override
  Widget build(BuildContext context) {
    return Container(
       child: child,
    );
  }
}