import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:znk/core/user/index.dart';
import 'package:znk/modules/tabs/owner/model.dart';
import 'package:znk/modules/tabs/owner/owner_bloc.dart';
import 'package:znk/utils/base/device.dart';

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
        title: Text('我的'),
        centerTitle: true,
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
          return ListView.separated(
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
            height: Device.isIOS ? Device.iOSRelativeHeight(120) : 100,
            child: Row(
              children: <Widget>[
                Container(
                  child: header,
                  width: Device.isIOS ? Device.iOSRelativeWidth(80) : 80,
                  height: Device.isIOS ? Device.iOSRelativeWidth(79) : 79,
                  margin: EdgeInsets.only(left: 20),
                ),
                Container(
                  width: 150,
                  height: Device.isIOS ? Device.iOSRelativeWidth(79) : 79,
                  child: Column(
                    children: <Widget>[
                      Container(
                        margin: EdgeInsets.only(left: 15, top: 10),
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
                        margin: EdgeInsets.only(left: 15, top: 10),
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
              ],
            ),
          ),
          onTap: () => onItemPressed(model),
            );
        break;
      case OwnerType.fileStore:
        return InkWell(
          child: Container(
            height: Device.isIOS ? Device.iOSRelativeHeight(53) : 53,
          ),
          onTap: () => onItemPressed(model),
        );
        break;
      case OwnerType.collection:
      
        return InkWell(
          child: Container(
            height: Device.isIOS ? Device.iOSRelativeHeight(53) : 53,
          ),
          onTap: () => onItemPressed(model),
        );
        break;
      case OwnerType.setting:
        return InkWell(
          child: Container(
            height: Device.isIOS ? Device.iOSRelativeHeight(53) : 53,
          ),
          onTap: () => onItemPressed(model),
        );
        break;
      default:
        return Container();
    }
    
  }
  
}

