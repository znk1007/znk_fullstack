import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:znk/core/user/index.dart';
import 'package:znk/images/manager.dart';
import 'package:znk/modules/tabs/owner/model.dart';
import 'package:znk/modules/tabs/owner/owner_bloc.dart';
import 'package:znk/utils/base/device.dart';
import 'package:znk/utils/base/random.dart';

class Owner extends StatelessWidget {
  final UserRepository _userRepository;
  Owner({Key key, @required UserRepository userRepository}) : 
  assert(userRepository != null),
  _userRepository = userRepository,
  super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(
          '我的',
          style: TextStyle(
            color: Colors.black,
          ),
        ),
        centerTitle: true,
        backgroundColor: Colors.white,
      ),
      body: BlocProvider(
        builder: (ctx) =>
          OwnerBloc(ownerModel: OwnerModel())..dispatch(Generate()),
          child: _OwnerLists(),
      ),
    );
  }
}

class _OwnerLists extends StatefulWidget {
  _OwnerLists({Key key}) : super(key: key);

  _OwnerListsState createState() => _OwnerListsState();
}

class _OwnerListsState extends State<_OwnerLists> {
  @override
  Widget build(BuildContext context) {
    final ownerBloc = BlocProvider.of<OwnerBloc>(context);
    return BlocBuilder(
      bloc: ownerBloc,
      builder: (BuildContext ctx, OwnerState state) {
        if (state is Loaded) {
          if (state.models.isEmpty) {
            return Container();
          }
          return Container(
            color: Color.fromARGB(255, 249, 249, 249),
            child: ListView.separated(
              itemCount: state.models.length,
              separatorBuilder: (BuildContext ctx, int section) {
                double sepHeight = 0;
                switch (section) {
                  case 0:
                    sepHeight = 20;
                    break;
                  case 1:
                  sepHeight = 1;
                    break;
                  case 2:
                  sepHeight = 20;
                    break;
                  default:
                }
                return Container(
                  height: sepHeight,
                  color: Colors.grey[100],
                );
              },
              itemBuilder: (BuildContext ctx, int idx) {
                return _OwnerItem(
                  model: state.models[idx],
                  onItemPressed: (OwnerModel model) {
                    print('current model type: ${model.type}');
                  },
                );
              },
            ),
          );
        } else {
          return Container();
        }
      },
    );
  }
}


class _OwnerItem extends StatelessWidget {
  final OwnerModel model;
  final Function(OwnerModel) onItemPressed;
  const _OwnerItem({
    Key key, 
    @required this.model, 
    @required this.onItemPressed
  }): super(key: key);
  
  @override
  Widget build(BuildContext context) {
    double marginLeft = Device.relativeWidth(20);
    Size iconSize = Size(Device.relativeWidth(27), Device.relativeWidth(27));
    Size arrowSize = Size(Device.relativeWidth(15), Device.relativeWidth(15));
    Size headerSize = Size(Device.relativeWidth(80), Device.relativeWidth(79));
    EdgeInsets headerMargin = EdgeInsets.only(left: 15, top: 10);
    Size infoSize = Size(Device.relativeWidth(150), Device.relativeWidth(79));
    double iconTxtSpace = Device.relativeWidth(10);
    switch (model.type) {
      case OwnerType.person:
        final header = (model.icon.startsWith('http') || 
                        model.icon.startsWith('https')) ?
                        Image.network(
                          model.icon,
                          fit: BoxFit.fill,
                        ) : Image.asset(
                          model.icon,
                          fit: BoxFit.fill,
                        );
        return InkWell(
          child: Container(
            color: Colors.white,
            height: Device.relativeHeight(120),
            child: Row(
              children: <Widget>[
                Container(
                  child: header,
                  width: headerSize.width,
                  height: headerSize.height,
                  margin: EdgeInsets.only(left: marginLeft),
                ),
                Container(
                  width: infoSize.width,
                  height: infoSize.height,
                  child: Column(
                    children: <Widget>[
                      Container(
                        margin: headerMargin,
                        alignment: Alignment.centerLeft,
                        child: Text(
                          model.title,
                          style: TextStyle(
                            color: Colors.black,
                            fontSize: 18,
                            fontWeight: FontWeight.w600,
                          ),
                        ),
                      ),
                      Container(
                        margin: headerMargin,
                        alignment: Alignment.centerLeft,
                        child: Text(
                          model.detail,
                          style: TextStyle(
                            color: Colors.grey[400],
                            fontSize: 15,
                            fontWeight: FontWeight.w300,
                          ),
                        ),
                      ),
                    ],
                  ),
                ),
                Container(
                  width: arrowSize.width,
                  height: arrowSize.height,
                  margin: EdgeInsets.only(left: Device.width - marginLeft - 2 * arrowSize.width - headerSize.width - infoSize.width),
                  child: Image.asset(CommonAsset.rightArrow),
                ),
              ],
            ),
          ),
          onTap: () => onItemPressed(model),
        );
        break;
      case OwnerType.fileStore:
        return InkWell(
          child: Container(
            color: Colors.white,
            height: Device.relativeHeight(53),
            child: Row(
              children: <Widget>[
                Container(
                  width: iconSize.width,
                  height: iconSize.height,
                  child: Image.asset(OwnerAsset.fileStore),
                  margin: EdgeInsets.only(left: marginLeft),
                ),
                Container(
                  width: infoSize.width,
                  height: iconSize.height,
                  margin: EdgeInsets.only(left: iconTxtSpace),
                  child: Text(
                    '我的网盘',
                    strutStyle: StrutStyle(
                      leading: 1,
                    ),
                  ),
                ),
                Container(
                  width: arrowSize.width,
                  height: arrowSize.height,
                  margin: EdgeInsets.only(left: Device.width - marginLeft - 2 * arrowSize.width - iconSize.width - infoSize.width - iconTxtSpace),
                  child: Image.asset(CommonAsset.rightArrow),
                ),
              ],
            ),
          ),
          onTap: () => onItemPressed(model),
        );
        break;
      case OwnerType.collection:
        return InkWell(
          child: Container(
            color: Colors.white,
            height: Device.relativeHeight(53),
            child: Row(
              children: <Widget>[
                Container(
                  width: iconSize.width,
                  height: iconSize.height,
                  child: Image.asset(OwnerAsset.collection),
                  margin: EdgeInsets.only(left: marginLeft),
                ),
                Container(
                  width: infoSize.width,
                  height: iconSize.height,
                  margin: EdgeInsets.only(left: iconTxtSpace),
                  child: Text(
                    '我的收藏',
                    strutStyle: StrutStyle(
                      leading: 1,
                    ),
                  ),
                ),
                Container(
                  width: arrowSize.width,
                  height: arrowSize.height,
                  margin: EdgeInsets.only(left: Device.width - marginLeft - 2 * arrowSize.width - iconSize.width - infoSize.width - iconTxtSpace),
                  child: Image.asset(CommonAsset.rightArrow),
                ),
              ],
            ),
          ),
          onTap: () => onItemPressed(model),
        );
        break;
      case OwnerType.setting:
        return InkWell(
          child: Container(
            color: Colors.white,
            height: Device.relativeHeight(53),
            child: Row(
              children: <Widget>[
                Container(
                  width: iconSize.width,
                  height: iconSize.height,
                  child: Image.asset(OwnerAsset.setting),
                  margin: EdgeInsets.only(left: marginLeft),
                ),
                Container(
                  width: infoSize.width,
                  height: iconSize.height,
                  margin: EdgeInsets.only(left: iconTxtSpace),
                  child: Text(
                    '设置',
                    strutStyle: StrutStyle(
                      leading: 1,
                    ),
                  ),
                ),
                Container(
                  width: arrowSize.width,
                  height: arrowSize.height,
                  margin: EdgeInsets.only(left: Device.width - marginLeft - 2 * arrowSize.width - iconSize.width - infoSize.width - iconTxtSpace),
                  child: Image.asset(CommonAsset.rightArrow),
                ),
              ],
            ),
          ),
          onTap: () => onItemPressed(model),
        );
        break;
      default:
        return Container();
    }
    
  }
  
}

