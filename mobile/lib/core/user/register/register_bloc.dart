import 'dart:async';
import 'package:bloc/bloc.dart';
import 'package:flutter/material.dart';
import 'package:rxdart/rxdart.dart';
import 'package:znk/core/user/index.dart';
import 'package:znk/core/user/register/index.dart';
import 'package:znk/utils/regexps/validator.dart';

class RegisterBloc extends Bloc<RegisterEvent, RegisterState> {
  UserRepository _userRepository;
  @override
  RegisterState get initialState => RegisterState.empty();

  @override
  Stream<RegisterState> transform(Stream<RegisterEvent> events, 
  Stream<RegisterState> Function(RegisterEvent event) next) {
    final observableStream = events as Observable<RegisterEvent>;
    final nonDebounceStream = observableStream.where((event){
      return (event is! RegisterAccountChanged && event is! RegisterPasswordChanged);
    });
    final debounceStream = observableStream.where((event){
      return (event is RegisterAccountChanged || event is RegisterPasswordChanged);
    }).debounceTime(Duration(milliseconds: 300));
    return super.transform(nonDebounceStream.mergeWith([debounceStream]), next);
  }

  @override
  Stream<RegisterState> mapEventToState(RegisterEvent event) async* {
    if (event is RegisterAccountChanged) {
      yield* _mapAccountChangedToState(event.account);
    } else if (event is RegisterPasswordChanged) {
      yield* _mapPasswordChangedToState(event.password);
    } else if (event is RegisterSubmitted) {
      yield* _mapFormSubmittedToState(event.ctx, event.account, event.password);
    }
  }
  
  Stream<RegisterState> _mapAccountChangedToState(String account) async* {
    yield currentState.update(
      isAccountValid: Validators.isValidAccount(account),
    );
  }

  Stream<RegisterState> _mapPasswordChangedToState(String password) async* {
    yield currentState.update(
      isPasswordValid: Validators.isValidPassword(password),
    );
  }

  Stream<RegisterState> _mapFormSubmittedToState(BuildContext ctx, String account, String password) async* {
    yield RegisterState.loading();
    try {
      await _userRepository.signUp(ctx, account: account, password: password);
      yield RegisterState.success();
    } catch (_) {
      yield RegisterState.failed();
    }

  }
  
}
