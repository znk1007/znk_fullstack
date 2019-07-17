import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:znk/core/user/auth_bloc.dart';
import 'package:znk/core/user/auth_event.dart';
import 'package:znk/core/user/login/index.dart';
import 'package:znk/core/user/user_repository.dart';
import 'package:znk/protos/generated/project/login.pb.dart';
import 'package:znk/utils/base/device.dart';
import 'package:znk/utils/base/random.dart';
import 'package:znk/utils/database/settings.dart';
import 'package:znk/utils/database/user.dart';
import 'package:znk/utils/hud/hud.dart';
import 'package:znk/core/user/register/index.dart';

class LoginPage extends StatelessWidget {
  static const String routeName = "/login";

  final UserRepository _userRepository;  

  LoginPage({Key key, @required UserRepository userRepository}):
    assert(userRepository != null),
    _userRepository = userRepository,
    super(key: key);

  @override
  Widget build(BuildContext context) {
    return new Scaffold(
      
      body: BlocProvider<LoginBloc>(
        builder: (context) => LoginBloc(userRepository: _userRepository),
        child: LoginForm(userRepository: _userRepository),
      ),
    );
  }
}

class LoginForm extends StatefulWidget {
  final UserRepository _userRepository;
  LoginForm({Key key, @required UserRepository userRepository}) : 
    assert(userRepository != null),
    _userRepository = userRepository,
    super(key: key);

  _LoginFormState createState() => _LoginFormState();
}

class _LoginFormState extends State<LoginForm> with SingleTickerProviderStateMixin{
  var _isActive = false;

  final TextEditingController _accountController = TextEditingController();
  final TextEditingController _passwordControler = TextEditingController();
  final HUD _hud = HUD();

  LoginBloc _loginBloc;

  bool get isOk => 
    _accountController.text.isNotEmpty && _passwordControler.text.isNotEmpty;

    bool isLoginButtonEnabled(LoginState state) {
      return state.isFormValid && isOk && !state.isSubmitting;
    }

  @override
  void initState() {
    super.initState();
    
    Settings.dao.recordPsw.then((val){
      setState(() {
        _isActive = val;
      });
    });
    _loginBloc = BlocProvider.of<LoginBloc>(context);
    _accountController.addListener(_onAccountChanged);
    _passwordControler.addListener(_onPasswordChanged);
  }
  @override
  Widget build(BuildContext context) {
    return BlocListener(
      bloc: _loginBloc,
      listener: (BuildContext context, LoginState state) {
        if (state.isFailure) {
          Scaffold.of(context)
            ..hideCurrentSnackBar()
            ..showSnackBar(
              SnackBar(
                content: Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: <Widget>[
                    Text("登录失败"),
                    Icon(Icons.error),
                  ],
                ),
              ),
            );
        }
        if (state.isSubmitting) {
          Scaffold.of(context)
            ..hideCurrentSnackBar()
            ..showSnackBar(
              SnackBar(
                content: Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: <Widget>[
                    Text("正在登录..."),
                    CircularProgressIndicator(),
                  ],
                ),
              ),
            );
        }
        if (state.isSuccess) {
          BlocProvider.of<AuthBloc>(context).dispatch(LoggedIn());
        }
      },
      child: BlocBuilder(
        bloc: _loginBloc,
        builder: (BuildContext context, LoginState state) {
          return Stack(
            children: <Widget>[
              _placeholderBackgroundView(),
              _BackgrounView(),
              _logoView(),
              _textField(),
              _loginButton(state),
              _remeberPassword(),
              _forgetButton(),
              _noAccount(context),
              _hud,
            ],
          );
        },
      )
    );
    
  }

  void _onAccountChanged() {
    _loginBloc.dispatch(
      AccountChanged(account: _accountController.text),
    );
  }

  void _onPasswordChanged() {
    _loginBloc.dispatch(
      PasswordChanged(password: _passwordControler.text),
    );
  }

  void _onFormSubmitted() {
    _loginBloc.dispatch(
      LoginButtonPressed(
        ctx: context,
        account: _accountController.text,
        password: _passwordControler.text,
      ),
    );
  }

  // 占位背景视图
  Widget _placeholderBackgroundView() {
    return Container(
      child: Image(
        image: AssetImage('lib/images/iOS/user/background.png'),
        width: Device.width,
      ),
    );
  }
  // logo
  Widget _logoView() {
    return Container(
      margin: EdgeInsets.fromLTRB(Device.iOSRelativeWidth(104), Device.iOSRelativeWidth(105), 0, 0),
      child: Image(
        image: AssetImage('lib/images/iOS/user/daxingzheng.png'),
      ),
    );
  }

  // 输入框
  Widget _textField() {
    final fieldHeight = Device.iOSRelativeWidth(55);
    return Container(
      height: fieldHeight * 2,
      decoration: BoxDecoration(
        color: Colors.white,
        boxShadow: [
          BoxShadow(
            color: Colors.grey[300],
            offset: Offset(-1, 1),
            blurRadius: 1.0
          ),
          BoxShadow(
            color: Colors.grey[300],
            offset: Offset(1, -1),
            blurRadius: 1.0
          ),
        ],
      ),
      margin: EdgeInsets.only(left: Device.iOSRelativeWidth(29), top: Device.iOSRelativeHeight(345), right: Device.iOSRelativeWidth(29)),
      child: Column(
        children: <Widget>[
          TextFormField(
            controller: _accountController,
            decoration: InputDecoration(
              prefixIcon: Image(
                image: AssetImage('lib/images/iOS/user/account.png'),
              ),
              hintText: '请输入用户',
              border: InputBorder.none,
            ),            
          ),
          Divider(
            color: Colors.grey[100],
            height: 1,
          ),
          TextFormField(
            controller: _passwordControler,
            decoration: InputDecoration(
              prefixIcon: Image(
                image: AssetImage('lib/images/iOS/user/password.png'),
              ),
              hintText: '请输入密码',
              border: InputBorder.none,
            ),
            obscureText: true,
            
          ),
        ],
      ),
    );
  }
  // 登录按钮
  Widget _loginButton(LoginState state) {
    // print('current state: $state');
    var startX = Device.iOSRelativeWidth(29);
    return Container(
      color: Colors.blue[600],
      width: Device.width - startX * 2,
      height: Device.iOSRelativeWidth(55),
      margin: EdgeInsets.only(left: Device.iOSRelativeWidth(29), top: Device.iOSRelativeHeight(480)),
      child: FlatButton(
        onPressed: isLoginButtonEnabled(state) == true ? _onFormSubmitted : null,
        textColor: Colors.white,
        child: Text('登录'),
      ),
    );
  }
  // 记住密码
  Widget _remeberPassword() {

    return Container(
      child: Row(
        mainAxisAlignment: MainAxisAlignment.start,
        children: <Widget>[
          Checkbox(
            onChanged: (val) {
              Settings.dao.setRecordPsw(val);
              setState(() {
                _isActive = val;
              });
            },
            value: _isActive,
            materialTapTargetSize: MaterialTapTargetSize.shrinkWrap,
          ),
          Text(
            '记住密码',
            style: TextStyle(
              color: Colors.grey[500],
              fontSize: 12,
            ),
            ),
        ],
      ),
      margin: EdgeInsets.only(left: Device.iOSRelativeWidth(16), top: Device.iOSRelativeHeight(550)),
      width: 120,
      height: 20,
      alignment: Alignment.centerLeft,
    );
  }
  // 忘记密码
  Widget _forgetButton() {
    return Container(
      child: FlatButton(
        onPressed: () {
          UserDB.dao.current.then((val) {
            print('user val: ${val.user.account}');
          });
        },
        textColor: Colors.grey[500],
        child: Text(
          '忘记密码?',
          style: TextStyle(
            fontSize: 12,
            decoration: TextDecoration.underline
          )
        ),
      ),
      width: 100,
      height: 20,
      margin: EdgeInsets.only(top: Device.iOSRelativeHeight(550), left: Device.width - 100 - Device.iOSRelativeWidth(16)),
    );
  }
  // 没有账号
  Widget _noAccount(BuildContext context) {
    double btnWidth = 100;
    return Container(
      child: FlatButton(
        onPressed: () {
          // widget._userRepository.signUp(
          //   context, 
          //   account: RandomManager.randomPhone(),
          //   password: '123456',
          // );
        },
        textColor: Colors.grey[500],
        child: Text(
          '还没有账号?',
          style: TextStyle(
            fontSize: 12,
            decoration: TextDecoration.underline
          )
        ),
      ),
      width: btnWidth,
      height: 20,
      margin: EdgeInsets.only(top: Device.iOSRelativeHeight(650), left: (Device.width - btnWidth) / 2),
    );
  }

  @override
  void dispose() {
    _accountController.dispose();
    _passwordControler.dispose();
    super.dispose();
  }


}

class _BackgrounView extends StatefulWidget {
  _BackgrounView({Key key}) : super(key: key);

  __BackgrounViewState createState() => __BackgrounViewState();
}

class __BackgrounViewState extends State<_BackgrounView> with TickerProviderStateMixin {
  Animation<dynamic> _movement;
  // 动画控制器
  AnimationController  _controller;

  @override
  void initState() {
    super.initState();
    _initController();
    _initAnimation();
    _startAnimation();
  }

  void _initController() {
    _controller = AnimationController(duration: Duration(seconds: 10), vsync: this);
  }

  void _initAnimation() {
    List<TweenSequenceItem> items = [];
    TweenSequenceItem item = TweenSequenceItem(
      tween: EdgeInsetsTween(
        begin: EdgeInsets.only(left: 2, top: 0, right: 0, bottom: 0),
        end: EdgeInsets.only(left: 0, top: 2, right: 0, bottom: 0)
      ),
      weight: 1,
    );
    items.add(item);
    item = TweenSequenceItem(
      tween: EdgeInsetsTween(
        begin: EdgeInsets.only(left: 0, top: 0, right: 2, bottom: 0),
        end: EdgeInsets.only(left: 0, top: 0, right: 0, bottom: 2)
      ),
      weight: 2,
    );
    items.add(item);
    _movement = TweenSequence(items).animate(
      CurvedAnimation(
        parent: _controller,
        curve: Interval(
          0.1, 
          0.5,
          curve: Curves.linear,
        ),
      ),
    )
    ..addListener((){
      setState(() {
        
      });
    })
    ..addStatusListener((status) {
      // if (status == AnimationStatus.completed) {
      //   _controller.reverse();
      // } else if (status == AnimationStatus.dismissed) {
      //   _controller.forward();
      // }
    });

    // _controller.forward();
    
  }

  Future _startAnimation() async {
    try {
      await _controller.repeat();
    } catch(e) {
      if (e is TickerCanceled) {
        print('ticker canceled');
      } else {
        print('animation failed $e');
      }
    }
    
  }

  @override
  Widget build(BuildContext context) {
    return Container(
       alignment: Alignment.topCenter,
          //  height: 500,
        padding: _movement.value,
        child: Image(
          image: AssetImage('lib/images/iOS/user/background.png'),
        ),
    );
  }

  @override
  void dispose() {
    _controller?.dispose();
    super.dispose();
  }
}