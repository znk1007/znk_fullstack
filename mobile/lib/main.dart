import 'package:bloc/bloc.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:path_provider/path_provider.dart' as path;
import 'package:english_words/english_words.dart' as wp;
import 'package:flutter/material.dart';
import 'package:znk/modules/launch/bloc_delegate.dart';
import 'package:znk/modules/tabs/tabs.dart';
import 'package:znk/core/user/auth_state.dart';
import 'package:znk/core/user/index.dart';
import 'package:znk/core/user/user_repository.dart';
import 'package:znk/utils/base/device.dart';
import 'package:znk/utils/base/routes.dart';
import 'package:znk/utils/database/user.dart';

import 'modules/launch/launch_screen.dart';



void main() {
  BlocSupervisor.delegate = MainBlocDelegate();
  final UserRepository userRepository = UserRepository();

  runApp(
    
    BlocProvider(
      builder: (context) => AuthBloc(userRepository: userRepository)
      ..dispatch(AppStarted()),
      child: ZnkProject(userRepository: userRepository),
    )
  );

}

class ZnkProject extends StatelessWidget {

  final UserRepository _userRepository;
  ZnkProject({Key key, @required UserRepository userRepository}):
    assert(userRepository != null),
    _userRepository = userRepository,
    super(key: key);

  void loadContactsFromDB(BuildContext ctx) async {
    final docPath = await path.getApplicationDocumentsDirectory();
    print('doc path ${docPath}');
  }

  @override
  Widget build(BuildContext context) {
    loadContactsFromDB(context);
    Device.getPackageInfo();
    return new MaterialApp(
      routes: Routes.generate(_userRepository),
      debugShowCheckedModeBanner: false,
      home: BlocBuilder(
        bloc: BlocProvider.of<AuthBloc>(context),
        builder: (BuildContext context, AuthState state) {
          if (state is Uninitialized) {
            return LaunchScreen();
          } else if (state is UnAuthenticated) {
            return LoginPage(userRepository: this._userRepository);
          } else if (state is Authenticated) {
            return Tabs(userRepository: this._userRepository);
          } else {
            return LaunchScreen();
          }
        },
      ),
    );
  }
}

