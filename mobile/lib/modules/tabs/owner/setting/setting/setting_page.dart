import 'package:flutter/material.dart';
import 'package:znk/modules/tabs/owner/setting/setting/model.dart';
import 'package:znk/utils/base/custom_theme.dart';

class SettingPage extends StatelessWidget {
  static const String routeName = "/setting";

  List<SettingModel> _models;

  SettingPage() {
    _models = SettingModel.generate();
  }

  @override
  Widget build(BuildContext context) {
    return new Scaffold(
      appBar: new AppBar(
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
            color: Colors.grey[100],
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
    return Container(
      child: Container(),
    );
  }
}
