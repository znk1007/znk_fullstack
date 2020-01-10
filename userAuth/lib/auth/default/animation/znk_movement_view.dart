
import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';

class MovementWidget extends StatefulWidget {
  MovementWidget({Key key}) : super(key: key);

  @override
  _MovementWidgetState createState() => _MovementWidgetState();
}

class _MovementWidgetState extends State<MovementWidget> {
  @override
  Widget build(BuildContext context) {
    OffsetPair.fromEventPosition(event)
    Transform.scale()
    return Container(
       child: PageView.builder(
         
       ),
    );
  }
}