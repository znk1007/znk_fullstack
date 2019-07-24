import 'package:equatable/equatable.dart';
import 'package:flutter/widgets.dart';
import 'package:meta/meta.dart';

@immutable
abstract class AuthEvent extends Equatable {
  AuthEvent([List props = const []]): super(props);
}

class AppStarted extends AuthEvent {
  @override
  String toString() {
    return 'AppStarted';
  }
}

class LoggedIn extends AuthEvent {
  final bool isLogined;
  LoggedIn({this.isLogined}): super([isLogined]);
  @override
  String toString() {
    return 'LoggedIn';
  }
}

class LoggedOut extends AuthEvent {
  final BuildContext ctx;
  LoggedOut({@required this.ctx});
  @override
  String toString() {
    return 'LoggedOut';
  }
}
