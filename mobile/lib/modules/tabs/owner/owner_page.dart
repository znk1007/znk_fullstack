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
          child: OwnerList(),
      ),
    );
  }
}

class OwnerList extends StatefulWidget {
  OwnerList({Key key}) : super(key: key);

  _OwnerListState createState() => _OwnerListState();
}

class _OwnerListState extends State<OwnerList> {
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
          return ListView(
            
          );
        } else {
          return Container();
        }
      },
    );
    
  }
}

