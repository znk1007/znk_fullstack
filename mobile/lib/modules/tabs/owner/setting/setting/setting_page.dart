import 'package:flutter/material.dart';
import 'package:znk/core/user/index.dart';
import 'package:znk/images/manager.dart';
import 'package:znk/modules/tabs/owner/setting/setting/model.dart';
import 'package:znk/utils/base/custom_theme.dart';
import 'package:znk/utils/base/device.dart';

class SettingPage extends StatelessWidget {
  static const String routeName = "/setting";

  List<SettingModel> _models;

  UserRepository _userRepository;

  SettingPage({@required UserRepository userRepository}) {
    _userRepository = userRepository;
    _models = SettingModel.generate();
  }

  @override
  Widget build(BuildContext context) {
    return new Scaffold(
      appBar: new AppBar(
        elevation: 0,
        iconTheme: IconThemeData(
          color: CustomColors.navigatorBackColor,
          size: 0.5,
        ),
        title: Text(
          '设置',
          style: TextStyle(
            color: Colors.black,
          ),
        ),
        centerTitle: true,
        backgroundColor: Colors.white,
      ),
      body: ListView.separated(
        itemCount: _models.length,
        itemBuilder: (BuildContext ctx, int idx) {
          return _SettimgItem(
            model: _models[idx],
            onItemPressed: (SettingModel m) {
              print('model type: ${m.type}');
              switch (m.type) {
                case SettingType.logout:
                  {
                    _userRepository.signOut().then((val) {

                    });
                  }
                  break;
                default:
              }
            },
          );
        },
        separatorBuilder: (BuildContext ctx, int section) {
          double sepHeight = 0;
          switch (section) {
            case 1:
              sepHeight = 15;
              break;
            case 3:
              sepHeight = 20;
              break;
            default:
              sepHeight = 1;
          }
          return Container(
            height: sepHeight,
            color: CustomColors.separatorColor,
          );
        },
      ),
    );
  }
}

class _SettimgItem extends StatelessWidget {
  final SettingModel model;
  final Function(SettingModel) onItemPressed;
  const _SettimgItem({Key key, @required this.model, @required this.onItemPressed}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    Alignment align = Alignment.centerLeft;
    double titleWidth = Device.relativeWidth(100);
    double marginLeft = Device.relativeWidth(20);
    bool hasArrow = (model.type == SettingType.security || model.type == SettingType.privacy);
    if (model.type == SettingType.logout) {
      align = Alignment.center;
      titleWidth = Device.width;
      marginLeft = 0;
    }
    return InkWell(
      child: Container(
        height: Device.relativeHeight(53),
        color: Colors.white,
        child: Row(
          children: <Widget>[
            Container(
              margin: EdgeInsets.only(left: marginLeft),
              alignment: align,
              width: titleWidth,
              height: Device.relativeHeight(53),
              child: Text(
                model.title,
              ),
            ),
            !hasArrow ? 
            model.type == SettingType.version ?
            Container(
              width: Device.width - Device.relativeWidth(100) - 2 * marginLeft,
              height: Device.relativeHeight(53),
              alignment: Alignment.centerRight,
              child: Text(
                'v${Device.version}',
                style: TextStyle(
                  color: Colors.grey[400],
                ),
              ),
            ) :
            Container() :
            Container(
              width: CustomMeasure.arrowSize.width,
                  height: CustomMeasure.arrowSize.height,
                  margin: EdgeInsets.only(left: Device.width - marginLeft - 2 * CustomMeasure.arrowSize.width - titleWidth),
                  child: Image.asset(CommonAsset.rightArrow),
            ),
          ],
        ),
      ),
      onTap: () => onItemPressed(model),
    );
  }
}
