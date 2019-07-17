import 'dart:async';
import 'package:bloc/bloc.dart';
import 'package:flutter/material.dart';
import 'package:znk/core/user/index.dart';

class AuthBloc extends Bloc<AuthEvent, AuthState> {
  final UserRepository _userRepository;
  AuthBloc({@required UserRepository userRepository}):
    assert(userRepository != null), 
    _userRepository = userRepository; 
  
  @override
  AuthState get initialState => Uninitialized();

  @override
  Stream<AuthState> mapEventToState(AuthEvent event) async* {
    if (event is AppStarted) {
      yield* _mapAppStartedToState();
    } else if (event is LoggedIn){
      yield * _mapLoggedInToState();
    } else if (event is LoggedOut) {
      yield * _mapLoggedOutToState();
    }
  }

  // 启动状态
  Stream<AuthState> _mapAppStartedToState() async* {
    try {
      final isSignedIn = await _userRepository.isSignedIn();
      if (isSignedIn) {
        final userId = await _userRepository.getUserId();
        yield Authenticated(userId);
      } else {
        yield UnAuthenticated();
      }
    } catch (_) {
      yield UnAuthenticated();
    }
  }
  // 已登录状态
  Stream<AuthState> _mapLoggedInToState() async* {
    final userId = await _userRepository.getUserId();
    yield Authenticated(userId);
  }
  // 退出登录状态
  Stream<AuthState> _mapLoggedOutToState() async* {
    yield UnAuthenticated();
    _userRepository.signOut();
  }
  
}
