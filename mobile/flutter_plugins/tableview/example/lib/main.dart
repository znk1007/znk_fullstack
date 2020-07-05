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
          numberOfSection: 2,
          separatorBuilder: (BuildContext context, ZNKIndexPath indexPath) {
            return Divider(height: 2);
          },
          viewForHeaderInSection: (context, section) => Container(
           child: Text('段头 $section'), 
          ),
          numberOfRowsInSection: (section) => 15,
          cellForRowAtIndexPath: (ctx, indexPath) => Container(
              child: Text('data ${indexPath.section} ${indexPath.row}'),
              // color: Colors.cyan,
          ),
          didSelectRowAtIndexPath: (context, indexPath) {
            print(
                'did select row at index path section: ${indexPath.section} row: ${indexPath.row}');
          },
        ),
      ),
    );
  }
}