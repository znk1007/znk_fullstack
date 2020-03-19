import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:znk_auth/znk_auth.dart';
class HomePage extends StatefulWidget {
  HomePage({Key key}) : super(key: key);

  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  String _platformVersion = 'Unknown';
  @override
  void initState() {
    super.initState();
    initPlatformState();
  }

  // Platform messages are asynchronous, so we initialize in an async method.
  Future<void> initPlatformState() async {
    String platformVersion;
    // Platform messages may fail, so we use a try/catch PlatformException.
    try {
      platformVersion = await ZnkAuth.platformVersion;
    } on PlatformException {
      platformVersion = 'Failed to get platform version.';
    }
  /*
  The context used to push or pop routes from the Navigator must be that of a widget that is a descendant of a Navigator widget.
  */
    // If the widget was removed from the tree while the asynchronous platform
    // message was in flight, we want to discard the reply rather than calling
    // setState to update our non-existent appearance.
    if (!mounted) return;

    setState(() {
      _platformVersion = platformVersion;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          title: const Text('Plugin example app'),
        ),
        body: Center(
          child: Stack(
            children: [
              Container(
                child:Text('Running on: $_platformVersion\n'),
                padding:EdgeInsets.only(top: 10) 
              ),
              Container(
                padding: EdgeInsets.only(top:10),
                child: FlatButton(
                  child: Text('登录'),
                  onPressed: () => ZnkAuth.show(context, (succ, msg) {
                    print('show login succ: $succ');
                    print('login msg: $msg');
                  }),
                )
              )
            ]
          )
        ),
      );
  }
}
