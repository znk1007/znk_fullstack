import 'package:flutter/cupertino.dart';
import 'package:meta/meta.dart';

@immutable 
class LoginState {
  final bool isAccountValid;
  final bool isPasswordValid;
  final bool isSubmitting;
  final bool isSuccess;
  final bool isFailure;

  bool get isFormValid => isAccountValid && isPasswordValid;

  LoginState({
    @required this.isAccountValid,
    @required this.isPasswordValid,
    @required this.isSubmitting,
    @required this.isSuccess,
    @required this.isFailure,
  });

  factory LoginState.empty() {
    return LoginState(
      isAccountValid: true,
      isPasswordValid: true,
      isSubmitting: false,
      isSuccess: false,
      isFailure: false,
    );
  }

  factory LoginState.loading() {
    return LoginState(
      isAccountValid: true,
      isPasswordValid: true,
      isSubmitting: true,
      isSuccess: false,
      isFailure: false,
    );
  }

  factory LoginState.failure() {
    return LoginState(
      isAccountValid: true,
      isPasswordValid: true,
      isSubmitting: false,
      isSuccess: false,
      isFailure: true,
    );
  }

  factory LoginState.success() {
    return LoginState(
      isAccountValid: true,
      isPasswordValid: true,
      isSubmitting: false,
      isSuccess: true,
      isFailure: false,
    );
  }

  LoginState update({
    bool isAccountValid,
    bool isPasswordValid,
  }){
    return copyWith(
      isAccountValid: isAccountValid,
      isPasswordValid: isPasswordValid,
      isSubmitting: false,
      isSuccess: false,
      isFailure: false, 
    );
  }
  
  LoginState copyWith({
    bool isAccountValid,
    bool isPasswordValid,
    bool isSubmitting,
    bool isSuccess,
    bool isFailure,
  }) {
    return LoginState(
      isAccountValid: isAccountValid ?? this.isAccountValid,
      isPasswordValid: isPasswordValid ?? this.isPasswordValid,
      isSubmitting: isSubmitting ?? this.isSubmitting,
      isSuccess: isSuccess ?? this.isSuccess,
      isFailure: isFailure ?? this.isFailure,
    );
  }

  @override
  String toString() {
    return '''
      LoginState {
        isAccountValid: $isAccountValid,
        isPasswordValid: $isPasswordValid,
        isSubmitting: $isSubmitting,
        isSuccess: $isSuccess,
        isFailure: $isFailure,
      }
    ''';
  }

}

