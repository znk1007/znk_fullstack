import 'dart:async';
import 'package:bloc/bloc.dart';
import 'package:flutter/widgets.dart';
import 'package:rxdart/rxdart.dart';
import 'package:znk/core/user/login/index.dart';
import 'package:znk/core/user/index.dart';
import 'package:znk/utils/regexps/validator.dart';

class LoginBloc extends Bloc<LoginEvent, LoginState> {
  final UserRepository _userRepository;
  LoginBloc({
    @required UserRepository userRepository,
  }): 
  assert(userRepository != null),
  _userRepository = userRepository;

  @override
  LoginState get initialState => LoginState.empty();

  @override
  Stream<LoginState> transform(Stream<LoginEvent> events, Stream<LoginState> Function(LoginEvent event) next) {
    final observableStream = events as Observable<LoginEvent>;
    final nonDebounceStream = observableStream.where((event) {
      return (event is! AccountChanged && event is! PasswordChanged);
    });
    final debounceStream = observableStream.where((event) {
      return (event is AccountChanged || event is PasswordChanged);
    }).debounceTime(Duration(milliseconds: 300));
    return super.transform(nonDebounceStream.mergeWith([debounceStream]), next);
  }

  @override
  Stream<LoginState> mapEventToState(LoginEvent event) async* {
    if (event is AccountChanged) {
      yield* _mapAccountChangedToState(event.account);
    } else if (event is PasswordChanged) {
      yield* _mapPasswordToState(event.password);
    } else if (event is LoginButtonPressed) {
      yield* _mapLoginButtonPressedToState(event.ctx, event.account, event.password);
    }
  }

  Stream<LoginState> _mapAccountChangedToState(String account) async* {

    yield currentState.update(
      isAccountValid: account.isNotEmpty,
    );
  }

  Stream<LoginState> _mapPasswordToState(String password) async* {
    yield currentState.update(
      isPasswordValid: password.length > 5,
    );
  }

  Stream<LoginState> _mapLoginButtonPressedToState(
    BuildContext ctx, 
    String account, 
    String password
  ) async* {
    UserError userError = await _userRepository.signIn(ctx, account: account, password: password);
    print('is login succ: ${userError.type}, description: ${userError.description}');
    if (userError.type == UserErrorType.none) {
      yield LoginState.success();
    } else {
      yield LoginState.failure();
    }
  }

}
