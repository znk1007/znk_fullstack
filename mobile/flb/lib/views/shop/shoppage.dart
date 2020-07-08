import 'package:flb/util/random/random.dart';
import 'package:flutter/material.dart';

class ShopPage extends StatefulWidget {

  static String id = 'shop';

  ShopPage({Key key}) : super(key: key);

  @override
  _ShopPageState createState() => _ShopPageState();
}

class _ShopPageState extends State<ShopPage> {
  @override
  Widget build(BuildContext context) {
    return Container(
       child: Text('购物车'),
       color: RandomHandler.randomColor,
    );
  }
}