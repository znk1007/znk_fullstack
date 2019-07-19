import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:znk/core/user/index.dart';
import 'package:znk/modules/tabs/owner/model.dart';
import 'package:znk/modules/tabs/owner/owner_bloc.dart';
import 'package:znk/utils/base/device.dart';

class Owner extends StatelessWidget {
  UserRepository _userRepository;
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
                color: Color.fromARGB(1, 249, 249, 249),
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
    
    // TODO: implement build
    switch (model.type) {
      case OwnerType.person:
        return GestureDetector(
          child: Container(
            height: Device.isIOS ? Device.iOSRelativeHeight(143) : 100,
            color: Colors.red,
            
          ),
          onTap: () => onItemPressed(model),
        );
        break;
      case OwnerType.fileStore:
        return GestureDetector(
          child: Container(
            height: Device.isIOS ? Device.iOSRelativeHeight(53) : 53,
            color: Colors.red,
          ),
          onTap: () => onItemPressed(model),
        );
        break;
      case OwnerType.collection:
        return GestureDetector(
          child: Container(
            height: Device.isIOS ? Device.iOSRelativeHeight(53) : 53,
            color: Colors.red,
          ),
          onTap: () => onItemPressed(model),
        );
        break;
      case OwnerType.setting:
        return GestureDetector(
          child: Container(
            height: Device.isIOS ? Device.iOSRelativeHeight(53) : 53,
            color: Colors.red,
          ),
          onTap: () => onItemPressed(model),
        );
        break;
      default:
        return Container();
    }
    
  }
  
}

