import 'package:equatable/equatable.dart';
import 'package:flutter/cupertino.dart';
import 'package:meta/meta.dart';

@immutable 
abstract class LoginEvent extends Equatable {
  LoginEvent([List props = const []]): super(props);
}

class AccountChanged extends LoginEvent {
  final String account;
  AccountChanged({
    @required this.account,
  }): super([account]);

  @override
  String toString() {
    return 'AccountChanged {account: $account}';
  }
}

class PasswordChanged extends LoginEvent {
  final String password;
  PasswordChanged({
    @required this.password,
  }): super([password]);

  @override
  String toString() {
    return 'PasswordChanged {password: $password}';
  }
}


class LoginButtonPressed extends LoginEvent {
  final BuildContext ctx;
  final String account;
  final String password;
  LoginButtonPressed({
    @required this.ctx,
    @required this.account,
    @required this.password,
  }):super([ctx, account, password]);

  @override
  String toString() {
    return 'LoginButtonPressed {account: $account, password: $password}';
  }
}
