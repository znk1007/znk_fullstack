import 'package:equatable/equatable.dart';
import 'package:flutter/material.dart';
import 'package:meta/meta.dart';

@immutable 
abstract class RegisterEvent extends Equatable {
  RegisterEvent([List props = const []]): super(props);
}

class RegisterAccountChanged extends RegisterEvent {
  final String account;
  RegisterAccountChanged({@required this.account}):super([account]);
  @override
  String toString() {
    return 'AccountChanged {account: $account}';
  }
}

class RegisterPasswordChanged extends RegisterEvent {
  final String password;
  RegisterPasswordChanged({@required this.password}): super([password]);
  @override
  String toString() {
    return 'PasswordChanged {password: $password}';
  }
}

class RegisterSubmitted extends RegisterEvent {
  final String account;
  final String password;
  final BuildContext ctx;
  RegisterSubmitted({@required this.account, @required this.password, @required this.ctx}): super([account, password, ctx]);
  @override
  String toString() {
    return 'RegisterSubmitted {account: $account, password: $password}';
  }
}

