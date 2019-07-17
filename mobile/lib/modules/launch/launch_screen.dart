import 'package:flutter/material.dart';
import 'package:package_info/package_info.dart';
import 'package:znk/utils/base/device.dart';

class LaunchScreen extends StatefulWidget {
  LaunchScreen({Key key}) : super(key: key);

  _LaunchScreenState createState() => _LaunchScreenState();

}

class _LaunchScreenState extends State<LaunchScreen> {

  String _version = '0.0.1';
  
  @override
  Widget build(BuildContext context) {
    PackageInfo.fromPlatform().then((info) {
      setState(() {
        this._version = info.version;
      });
    });
    return Scaffold(
      body: Container(
      child: Column(
        children: <Widget>[
          Container(
            height: Device.iOSRelativeWidth(67),
            margin: EdgeInsets.fromLTRB(Device.iOSRelativeWidth(90), Device.iOSRelativeWidth(201), Device.iOSRelativeWidth(93), 0),
            decoration: BoxDecoration(
              image: DecorationImage(
                image: AssetImage('lib/images/iOS/launch/launchImg.png'),
              )
            ),
          ),
          Container(
            height: Device.iOSRelativeWidth(26),
            margin: EdgeInsets.fromLTRB(Device.iOSRelativeWidth(45), Device.iOSRelativeHeight(390), 38, 0),
            decoration: BoxDecoration(
              image: DecorationImage(
                image: AssetImage('lib/images/iOS/launch/launchVersionBg.png')
              )
            ),
            child: Align(
              alignment: FractionalOffset(0.5, 0.5),
              child: Text(
                'v${this._version}' ?? '',
                textAlign: TextAlign.center,
                textScaleFactor: 0.8,
              ),
            ),
          ),
          Container(
            child: Text(
                'Copyright 2018-2030 Xiangyi Mobile ALL Rights Reserved',
                textAlign: TextAlign.center,
                textScaleFactor: 0.8,
                style: TextStyle(
                  color: Colors.grey,
                ),
                
              ),
            margin: EdgeInsets.fromLTRB(Device.iOSRelativeWidth(60), Device.iOSRelativeWidth(20), Device.iOSRelativeWidth(60), 0),
          )
        ],
      ),
    ),
    );
  }
}



